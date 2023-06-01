package main

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wos-project/wos-core-transactor/app/config"
	"github.com/wos-project/wos-core-transactor/app/models"
	"github.com/wos-project/wos-core-transactor/app/utils"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func init() {
	config.ConfigPath = flag.String("config", "config.yaml", "path to YAML config file")
	config.InitializeConfiguration()
	flag.Parse()
	utils.InitMediaStorage()
	models.OpenDatabase()
	models.DropAllTables()
	models.CloseDatabase()
	models.InitializeDatabase()
}

func TestAirdrop(t *testing.T) {
	assert.True(t, true)
}
