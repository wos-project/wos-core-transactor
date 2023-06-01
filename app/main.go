package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/spf13/viper"

	"github.com/wos-project/wos-core-transactor/app/config"
	"github.com/wos-project/wos-core-transactor/app/models"
	"github.com/wos-project/wos-core-transactor/app/services"
)

func main() {

	config.ConfigPath = flag.String("config", "config.yaml", "path to YAML config file")
	flag.Parse()
	config.InitializeConfiguration()

	glog.Infof("starting %s mode", viper.GetString("mode"))

	models.InitializeDatabase()
	defer models.CloseDatabase()

	services.SetupCronJobs()

	select{}
}
