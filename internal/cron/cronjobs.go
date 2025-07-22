package cron

import (
	"log/slog"

	"github.com/robfig/cron/v3"
	"greenlight.goodlooking.com/internal/data"
)

type CronJobs struct {
	PhotoLocation PhotoLocationCronJobs
	logger        *slog.Logger
}

func NewCronJobs(models data.Models, logger *slog.Logger) *CronJobs {
	if logger != nil {
		logger.Info("It is ok")
	}
	return &CronJobs{
		PhotoLocation: PhotoLocationCronJobs{},
		logger:        logger,
	}
}

func (cj *CronJobs) StartAll() {

	c := cron.New(cron.WithSeconds())

	cj.StartCronJobs(c)
	c.Start()
	cj.logger.Info("all cron jobs started")
}
