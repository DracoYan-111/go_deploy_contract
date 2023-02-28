package repository

import (
	"GoContractDeployment/models"
	"context"
)

// PostRepo explain...
type PostRepo interface {
	// Operate Find tasks with incoming status
	// @param ctx context.Context
	// @param status The status of the task in the database
	// @return related information array
	Operate() ([]*models.DataPost, error)

	// AddJob Insert data
	// @param ctx context.Context
	// @param dataPost Add the information passed in by the task
	// @return return success status
	AddJob(ctx context.Context, dataPost []models.ReceivePost) string

	// GetOne Get updated data
	// @return Single database information
	GetOne() (*models.DataPost, error)

	// UpdateTask Update task information
	// @param which Update the task information
	// @param dataPost Update the information passed in by the task
	// @return Success message
	UpdateTask(which string, dataPost models.DataPost) string

	// UpdateState Update task status to complete
	// @param idArray Id list
	UpdateState(idArray []int64) string
}
