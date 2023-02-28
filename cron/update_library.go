package cron

import (
	"GoContractDeployment/handler/http"
	"GoContractDeployment/internal"
	"GoContractDeployment/internal/deploy"
	"GoContractDeployment/models"
	"GoContractDeployment/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robfig/cron"
	"log"
)

// UpdateLibrary Update database scheduled tasks
func UpdateLibrary(jobHandler *handler.CreateTask) {

	loading, err := utils.ConfigurationLoading("web3", []string{"minter", "tokenUri"})
	if err != nil {
		log.Println(err)
	}

	cronJob := cron.New()
	spec := "*/20 * * * * ?"
	err = cronJob.AddFunc(spec, func() {

		jobData, err := jobHandler.Repo.GetOne()
		if err == nil {
			log.Printf("UpdateLibrary:Automatic deployment task starts")

			structure := deploy.Structure{
				Name:           jobData.ContractName,
				Symbol:         jobData.ContractName,
				Minter:         common.HexToAddress(loading[0]),
				TokenURIPrefix: loading[1],
			}

			addressHex, txDataHashHex, gasUsed, currentStatus := deploy.GoContractDeployment(structure)
			if addressHex == "" && txDataHashHex == "" {
				log.Println(structure.Name, "<==== UpdateLibrary:Deployment failed ====>")
			} else {
				log.Println(structure.Name, "<++++ UpdateLibrary:Deployed ++++>")
			}

			gasUse := gasUsed.SetInt64(gasUsed.Int64())

			//gasUsed := deploy.GoTransactionNews(client, txDataHashHex)
			var gasUST float64
			if gasUse.Int64() != 0 {
				gasUST = internal.GetBnbToUsdt(gasUsed)
				log.Println("<++++ UpdateLibrary:Price query completed ++++>")
			}

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
				CurrentStatus: currentStatus,
			}

			jobHandler.Repo.UpdateTask(models.UpdateTaskOne, dataPos)
			log.Printf("<++++ UpdateLibrary:update completed ++++>")
		}
	})
	if err != nil {
		return
	}
	cronJob.Start()
}
