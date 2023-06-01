package services

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

// NOTE: can create struct that we marshal config into using viper.UnmarshalKey

type airdropJob struct { count int }
func (a *airdropJob) Run() {
	ServiceAirdrop()
}

// SetupCronJobs reads cron jobs schedules and sets up jobs
func SetupCronJobs() {

	c := cron.New(cron.WithSeconds())
	ja := airdropJob{}

	if viper.GetString("services.tx.cron") != "" {
		c.AddJob(viper.GetString("services.tx.cron"), cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&ja))
	}

	c.Start()
}
