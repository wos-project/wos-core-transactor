package models

import (
	"testing"

	"github.com/golang/glog"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	viper.SetConfigFile("../config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		glog.Fatalf("Error initializing config, %s", err)
	}
}

func TestDatabase(t *testing.T) {
	OpenDatabase()
	DropAllTables()
	Db.Exec("DROP TABLE IF EXISTS migrations")
	CloseDatabase()

	InitializeDatabase()

	// create sample data
	tx := Transaction{Uid: "xyz"}	
	res := Db.Create(&tx)
	assert.Nil(t, res.Error)
}
