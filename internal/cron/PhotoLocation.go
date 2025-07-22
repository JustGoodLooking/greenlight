package cron

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/robfig/cron/v3"
	"greenlight.goodlooking.com/internal/data"
)

type PhotoLocationCronJobs struct {
	models data.Models
	logger *slog.Logger
}

func (cj *CronJobs) StartCronJobs(c *cron.Cron) {
	fmt.Println("start")
	if cj.PhotoLocation.logger == nil {
		fmt.Println("no loggerrrrrr")
	}

	_, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("123123123123")
		cj.PhotoLocation.logger.Info("hello", "time", 123)
	})

	if err != nil {
		log.Fatal(err)
	}

}
