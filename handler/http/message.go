package handler

import (
	"GoContractDeployment/models"
	"GoContractDeployment/navigation"
	repository "GoContractDeployment/repository"
	"GoContractDeployment/repository/create"
	"encoding/json"
	"log"
	"net/http"
)

// NewJobHandler 新任务处理程序
func NewJobHandler(db *navigation.DB) *CreateTask {
	return &CreateTask{
		// 加载到接口的实例
		Repo: create.NewSQLPostRepo(db.SQL),
	}
}

// CreateTask 返回所有的接口
type CreateTask struct {
	Repo repository.PostRepo
}

// Receive 接收传入数据
type Receive struct {
	DataList string `json:"sign"`
}

// ReturnData 返回结束数据
type ReturnData struct {
	State   int         `json:"state"`
	Payload interface{} `json:"data"`
}

// CreateJob 添加任务
func (task *CreateTask) CreateJob(writer http.ResponseWriter, request *http.Request) {
	// 获取传入数据
	var requestBody Receive
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Println("<==== 获取传入数据异常 ====>", err)
	}

	// 判断字符串不能为空
	if len(requestBody.DataList) != 0 {
		var data []models.ReceivePost

		err = json.Unmarshal([]byte(requestBody.DataList), &data)
		if err != nil {
			log.Println("<==== 接收字符串为空 ====>", err)
		}
		// 插入数据库
		//okData := task.repo.AddJob(request.Context(), data)

		a := task.Repo.Operate(request.Context(), 0)

		log.Println(a)
		respondWithData(writer, http.StatusOK, "okData")
	} else {
		respondWithData(writer, http.StatusBadRequest, "data is empty")
	}
}

// respondWithData 返回信息
func respondWithData(writer http.ResponseWriter, code int, payload interface{}) {
	var returnData ReturnData
	returnData.State = code
	returnData.Payload = payload

	respondentJSON(writer, code, returnData)
}

// respondentJSON 处理返回信息
func respondentJSON(writer http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	_, err := writer.Write(response)
	if err != nil {
		log.Println("Respondent JSON Exception:", err)
	}
}
