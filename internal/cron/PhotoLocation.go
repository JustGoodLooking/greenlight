package cron

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

type PhotoLocationCronJobs struct {
}

func (cj *CronJobs) StartCronJobs(c *cron.Cron) {
	fmt.Println("start")

	if cj.logger == nil {
		fmt.Println("no loggerrrrrr")
	}

	_, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("123123123123")
		cj.logger.Info("hello", "time", 123)
	})

	if err != nil {
		log.Fatal(err)
	}

}
