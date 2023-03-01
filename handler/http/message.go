package handler

import (
	"GoContractDeployment/models"
	"GoContractDeployment/navigation"
	"GoContractDeployment/repository"
	"GoContractDeployment/repository/create"
	"GoContractDeployment/utils"
	"encoding/json"
	"log"
	"net/http"
)

// NewJobHandler new task handler
// @param db Database connection information
func NewJobHandler(db *navigation.DB) *CreateTask {
	return &CreateTask{
		// instance loaded into the interface
		Repo: create.NewSQLPostRepo(db.SQL),
	}
}

// CreateTask return all interfaces
type CreateTask struct {
	Repo repository.PostRepo
}

// Receive receive incoming data
type Receive struct {
	DataList string `json:"sign"`
}

// ReturnData return end data
type ReturnData struct {
	State   int         `json:"state"`
	Payload interface{} `json:"data"`
}

// CreateJob add task
// @param writer Build http response
// @param request Received by the server and returned to the client
func (task *CreateTask) CreateJob(writer http.ResponseWriter, request *http.Request) {

	var requestBody Receive
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Println("<==== message:获取传入数据异常 ====>", err)
	}

	// Judgment string cannot be empty
	if len(requestBody.DataList) != 0 {
		var data []models.ReceivePost
		requestBody.DataList, err = utils.AesDecrypt(requestBody.DataList)
		if err != nil {
			log.Println("<==== message:解密字符串异常 ====>", err)
		}
		err = json.Unmarshal([]byte(requestBody.DataList), &data)
		if err != nil {
			log.Println("<==== message:收到的字符串为空 ====>", err)
		}
		// insert database
		okData := task.Repo.AddJob(request.Context(), data)

		respondWithData(writer, http.StatusOK, okData)

	} else {
		respondWithData(writer, http.StatusBadRequest, "message:数据为空")
	}
}

// respondWithData Returned messages
func respondWithData(writer http.ResponseWriter, code int, payload interface{}) {
	var returnData ReturnData
	returnData.State = code
	returnData.Payload = payload

	respondentJSON(writer, code, returnData)
}

// respondentJSON Handle returned information
func respondentJSON(writer http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	_, err := writer.Write(response)
	if err != nil {
		log.Println("message:响应者 JSON 异常:", err)
	}
}
