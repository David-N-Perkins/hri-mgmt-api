// +build !tests

/**
 * (C) Copyright IBM Corp. 2020
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/Alvearie/hri-mgmt-api/batches"
	"github.com/Alvearie/hri-mgmt-api/common/actionloopmin"
	"github.com/Alvearie/hri-mgmt-api/common/elastic"
	"github.com/Alvearie/hri-mgmt-api/common/kafka"
	"github.com/Alvearie/hri-mgmt-api/common/param"
	"github.com/Alvearie/hri-mgmt-api/common/response"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	actionloopmin.Main(createMain)
}

func createMain(params map[string]interface{}) map[string]interface{} {
	logger := log.New(os.Stdout, "batches/create: ", log.Llongfile)
	start := time.Now()
	logger.Printf("start createMain, %s \n", start)

	esClient, err := elastic.ClientFromParams(params)
	if err != nil {
		return response.Error(http.StatusInternalServerError, err.Error())
	}

	kafkaWriter, err := kafka.NewWriterFromParams(params)
	if err != nil {
		return response.Error(http.StatusInternalServerError, err.Error())
	}
	resp := batches.Create(params, param.ParamValidator{}, esClient, kafkaWriter)
	logger.Printf("processing time createMain, %d milliseconds \n", time.Since(start).Milliseconds())
	return resp
}
