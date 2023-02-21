package repsitory

import (
	models "GoContractDeployment/models"
	"context"
)

// PostRepo explain...
type PostRepo interface {
	// Operate 查找传入状态的任务
	// @param ctx context.Context
	// @param status The status of the task in the database
	// @return related information array
	Operate(ctx context.Context, status int64) []*models.DataPost

	// AddJob 插入数据
	// @param ctx context.Context
	// @param post Add the information passed in by the task
	// @return return success status
	AddJob(ctx context.Context, post []models.ReceivePost) string

	GetOne() *models.DataPost
}
