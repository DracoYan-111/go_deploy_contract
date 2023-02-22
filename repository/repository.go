package repository

import (
	"GoContractDeployment/models"
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
	// @param dataPost Add the information passed in by the task
	// @return return success status
	AddJob(ctx context.Context, dataPost []models.ReceivePost) string

	// GetOne 获取未更新的数据
	// @return 单条数据库信息
	GetOne() (*models.DataPost, error)

	// UpdateTask 更新任务信息
	// @param which Update the task information
	// @param dataPost Update the information passed in by the task
	// @return 成功信息
	UpdateTask(which string, dataPost models.DataPost) string
}
