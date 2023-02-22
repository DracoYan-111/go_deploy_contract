package cron

import (
	"GoContractDeployment/handler/http"
	"GoContractDeployment/internal/deploy"
	"GoContractDeployment/models"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-ini/ini"
	"github.com/robfig/cron"
	"log"
)

//// NewJobHandler 新任务处理程序
//func NewJobHandler(db *navigation.DB) *CreateTask {
//	return &CreateTask{
//		// 加载到接口的实例
//		repo: create.NewSQLPostRepo(db.SQL),
//	}
//}
//
//// CreateTask 返回所有的接口
//type CreateTask struct {
//	repo repository.PostRepo
//}

func UpdateLibrary(cfg *ini.File, jobHandler *handler.CreateTask) {
	cronJob := cron.New()
	spec := "*/12 * * * * ?"
	err := cronJob.AddFunc(spec, func() {
		jobData, err := jobHandler.Repo.GetOne()
		if err == nil {
			log.Printf("自动部署任务开始")

			structure := deploy.Structure{
				Name:           jobData.ContractName,
				Symbol:         jobData.ContractName,
				Minter:         common.HexToAddress(cfg.Section("web3").Key("minter").String()),
				TokenURIPrefix: cfg.Section("web3").Key("tokenUri").String(),
			}

			addressHex, txDataHashHex, gasUsed := deploy.GoContractDeployment(structure)

			fmt.Println(structure.Name, "部署完毕")

			gasUsed.SetInt64(gasUsed.Int64())
			//gasUsed := deploy.GoTransactionNews(client, txDataHashHex)

			gasUST := 10.0000000000 //:= internal.GetBnbToUsdt(gasUsed)

			dataPos := models.DataPost{
				ID:            jobData.ID,
				Opcode:        jobData.Opcode,
				ContractName:  jobData.ContractName,
				ContractAddr:  addressHex,
				ContractHash:  txDataHashHex,
				GasUsed:       gasUsed.Int64(),
				GasUST:        gasUST,
				ChainId:       jobData.ChainId,
				CreatedAt:     jobData.CreatedAt,
				CurrentStatus: int64(1),
			}

			jobHandler.Repo.UpdateTask(models.UpdateTaskOne, dataPos)
		}

	})
	if err != nil {
		return
	}
	cronJob.Start()
}
