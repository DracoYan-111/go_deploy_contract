package cron

import (
	handler "GoContractDeployment/handler/http"
	"GoContractDeployment/models"
	"GoContractDeployment/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"strconv"
)

func ReturnStatus(jobHandler *handler.CreateTask) {
	cronJob := cron.New()
	spec := "*/40 * * * * ?"
	err := cronJob.AddFunc(spec, func() {

		jobData, err := jobHandler.Repo.Operate()
		if err == nil {
			log.Println("ReturnStatus:Cron job is running")

			if len(jobData) != 0 {
				data, idList, err := processData(jobData)

				encrypt, err := utils.Encrypt(data)
				if err != nil {
					log.Panicln("ReturnStatus:", encrypt, err)
				}

				transfer, err := request(encrypt)
				if err != nil {
					log.Println("ReturnStatus:", transfer, err)
				} else {
					state := jobHandler.Repo.UpdateState(idList)
					log.Println(state)
				}
			}
		}
	})
	if err != nil {
		return
	}
	cronJob.Start()
}

// processData Process the data to be sent to the server
func processData(jobData []*models.DataPost) (string, []int64, error) {
	// Define a Return Post array
	var returnPosts []*models.ReturnPost
	var idList []int64

	loading, err := utils.ConfigurationLoading("web3", []string{"publicKey", "minter"})
	if err != nil {
		log.Panicln("ReturnStatus:", err)
	}

	if len(jobData) != 0 {
		for _, jobData := range jobData {
			numInt64, err := strconv.ParseInt(jobData.Opcode, 10, 64)
			if err != nil {
				fmt.Println(err)
				return "", idList, err
			}
			id := jobData.ID
			returnPost := &models.ReturnPost{
				Opcode:         numInt64,
				ChainId:        jobData.ChainId,
				GasUST:         jobData.GasUST,
				ContractAddr:   jobData.ContractAddr,
				ContractHash:   jobData.ContractHash,
				ContractOwner:  loading[0],
				ContractMinter: loading[1],
			}
			returnPosts = append(returnPosts, returnPost)
			idList = append(idList, id)
		}
		// Convert structure array to JSON format
		jsonBytes, err := json.Marshal(returnPosts)
		if err != nil {
			fmt.Println(err)
			return "ReturnStatus:conversion failed", idList, err
		}
		return string(jsonBytes), idList, nil
	}

	return "", idList, errors.New("ReturnStatus:data is empty")
}

type returnData struct {
	DataList string `json:"crossChainBack"`
}

// request:Initiate an http request to return the modified information
func request(sign string) (string, error) {

	jsonData, err := json.Marshal(returnData{DataList: sign})
	if err != nil {
		return "Request failed", err
	}

	loading, err := utils.ConfigurationLoading("server", []string{"ask"})
	if err != nil {
		log.Panicln("ReturnStatus:", err)
	}

	// Create an HTTP request object
	req, err := http.NewRequest("POST", loading[0], bytes.NewBuffer(jsonData))
	if err != nil {
		return "ReturnStatus:", err
	}

	// Set the Content-Type field in the request header
	req.Header.Set("Content-Type", "application/json")

	// Initiate an HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "ReturnStatus:", err
	}

	// Output the status code of the response
	if resp.StatusCode != 200 {
		return "ReturnStatus:", errors.New("<==== abnormal state:" + strconv.Itoa(resp.StatusCode) + "====>")
	}
	return "<++++ ReturnStatus:The request was initiated successfully ++++>", nil
}
