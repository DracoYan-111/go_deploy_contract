package handler

import (
	getPrice "GoContractDeployment/internal"
	deploy "GoContractDeployment/internal/deploy"
	"GoContractDeployment/models"
	"GoContractDeployment/navigation"
	repository "GoContractDeployment/repository"
	create "GoContractDeployment/repository/create"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-ini/ini"
	"log"
	"net/http"
)

// NewJobHandler 新作业处理程序
func NewJobHandler(db *navigation.DB) *CreateTask {
	return &CreateTask{
		// 加载到接口的实例
		repo: create.NewSQLPostRepo(db.SQL),
	}
}

// CreateTask 返回所有的接口
type CreateTask struct {
	repo repository.PostRepo
}

// operate all create data
//func (p *CreateTask) operate(w http.ResponseWriter, r *http.Request) {
//
//	num, _ := strconv.Atoi(chi.URLParam(r, "id"))
//	log.Print(int64(num))
//
//	payload := p.repo.Operate(r.Context(), int64(num))
//
//	for i := 0; i < len(payload); i++ {
//		log.Println(payload[i])
//	}
//
//	//respondentJSON(w, http.StatusOK, payload)
//}

type Receive struct {
	DataList string `json:"sign"`
}

// CreateJob 添加任务
func (p *CreateTask) CreateJob(w http.ResponseWriter, request *http.Request) {

	// 获取传入数据
	var requestBody Receive
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// 判断字符串不能为空
	if len(requestBody.DataList) != 0 {
		var data []models.ReceivePost

		err = json.Unmarshal([]byte(requestBody.DataList), &data)
		if err != nil {
			log.Fatal("====接收字符串为空====", err)
		}

		// 插入数据库
		okData := p.repo.AddJob(request.Context(), data)

		respondWithData(w, http.StatusOK, okData)
		// 读取配置文件
		cfg, err := ini.Load("config.ini")
		if err != nil {
			log.Fatal("====配置文件读取异常====", err)
		}
		minter := cfg.Section("web3").Key("minter").String()
		tokenUri := cfg.Section("web3").Key("tokenUri").String()
		// 返回完成后进行任务
		payload := p.repo.Operate(request.Context(), int64(0))

		// todo 完善数据库中的信息
		for i := 0; i < len(payload); i++ {
			log.Println(payload[i])

			// 开始部署合约
			structure := deploy.Structure{
				Name:           payload[i].ContractName,
				Symbol:         payload[i].ContractName,
				Minter:         common.HexToAddress(minter),
				TokenURIPrefix: tokenUri,
			}

			// 获取合约部署环境
			client, _, hashHex := deploy.GoContractDeployment(structure)
			// 获取交易的gas使用
			gasUsd := deploy.GoTransactionNews(client, hashHex)
			// gas使用转换为usdt
			_ = getPrice.GetBnbToUsdt(gasUsd)

		}
	} else {
		respondWithData(w, http.StatusBadRequest, "data is empty")
	}
}

type ReturnData struct {
	State   int         `json:"state"`
	Payload interface{} `json:"data"`
}

// respondWithData 返回信息
func respondWithData(w http.ResponseWriter, code int, payload interface{}) {
	var returnData ReturnData
	returnData.State = code
	returnData.Payload = payload

	respondentJSON(w, code, returnData)
}

// respondentJSON 处理返回信息
func respondentJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Println("Respondent JSON Exception:", err)
	}
}
