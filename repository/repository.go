package repsitory

import (
	models "GoContractDeployment/models"
	"context"
)

// PostRepo explain...
type PostRepo interface {
	Operate(ctx context.Context, status int64) []*models.Post

	// AddJob 插入数据
	// @param ctx context.Context
	// @param post *models.Post
	AddJob(ctx context.Context, p []models.ReceivePost) string
}
