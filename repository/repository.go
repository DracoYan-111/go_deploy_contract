package repsitory

import (
	models "GoContractDeployment/models"
	"context"
)

// PostRepo explain...
type PostRepo interface {
	Fetch(ctx context.Context, id int64) *models.Post

	//
	AddJob(ctx context.Context, p models.Post) string
}
