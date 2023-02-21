package cron

import (
	"GoContractDeployment/handler/http"
	"fmt"
	"github.com/robfig/cron"
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

func UpdateLibrary(jobHandler *handler.CreateTask) {
	cronJob := cron.New()
	spec := "*/1 * * * * ?"
	cronJob.AddFunc(spec, func() {
		a := jobHandler.Repo.GetOne()
		fmt.Println(a)
	})

	cronJob.Start()
}
