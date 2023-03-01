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
			log.Println("ReturnStatus:Cron 作业正在运行")

			if len(jobData) != 0 {
				data, idList, err := processData(jobData)

				encrypt, err := utils.Encrypt(data)
				if err != nil {
					log.Panicln("ReturnStatus:", encrypt, err)
				}

				transfer, err := request(encrypt)
				if err != nil {
					log.Println("ReturnStatus:", transfer, err)
				}
				state := jobHandler.Repo.UpdateState(idList)
				log.Println(state)
			}
		}
	})
	if err != nil {
		return
	}
	cronJob.Start()
}

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
			return "ReturnStatus:json转换失败", idList, err
		}
		return string(jsonBytes), idList, nil
	}

	return "", idList, errors.New("ReturnStatus:数据为空")
}

type returnData struct {
	DataList string `json:"crossChainBack"`
}

func request(sign string) (string, error) {
	// 创建一个returnData结构体
	//data := returnData{
	//	DataList: sign,
	//}

	// 将returnData结构体转换为JSON格式的字节数组
	jsonData, err := json.Marshal(returnData{DataList: sign})
	if err != nil {
		return "请求失败", err
	}

	//log.Fatal(bytes.NewBuffer(jsonData))
	// 创建一个HTTP请求对象
	req, err := http.NewRequest("POST", "http://192.168.18.155:8089/dc/contract/crossChainBack", bytes.NewBuffer(jsonData))
	if err != nil {
		return "请求失败", err
	}

	// 设置请求头中的Content-Type字段
	req.Header.Set("Content-Type", "application/json")

	// 发起HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "请求失败", err
	}

	// 输出响应的状态码
	if resp.StatusCode != 200 {
		return "请求失败", errors.New("<==== 状态相应异常:" + strconv.Itoa(resp.StatusCode) + "====>")
	}
	return "<++++ 请求发起成功 ++++>", nil
}
