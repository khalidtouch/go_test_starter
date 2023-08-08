package migrations

import (
	"github.com/williaminfante/go_test_starter/config"
	"github.com/williaminfante/go_test_starter/entity"
)

func getModels() []interface{} {
	return []interface{}{&entity.User{}}
}


func MigrateTable() {
	db := config.GetDb() 
	db.AutoMigrate(getModels()...)
}