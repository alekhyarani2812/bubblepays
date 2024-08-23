package jobs

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

func GetJobs(ctx context.Context) ([]models.Jobs, error) {
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errort(ctx, "error in db connection")
		return nil, err
	}
	var jobs []models.Jobs
	if err := dbConnection.Find(&jobs).Error; err != nil {
		log.Errorft(ctx, "failed to retrieve jobs data: %v", err)
		return nil, err
	}
	log.Infof("jobs : %v ", jobs)
	return jobs, nil
}
