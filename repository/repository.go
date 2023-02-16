package repsitory

import (
	"GoContractDeployment/models"
	"context"
)

// PostRepo explain...
type PostRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.Post, error)

	//
	GetByID(ctx context.Context, id int64) (*models.Post, error)

	//创建
	Create(ctx context.Context, p *models.Post) (int64, error)

	//更新
	Update(ctx context.Context, p *models.Post) (*models.Post, error)

	//删除
	Delete(ctx context.Context, id int64) (bool, error)
}
