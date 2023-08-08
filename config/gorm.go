package config

import (
	"fmt"
	"log"
	"os"
	"time"

	mysql "gorm.io/driver/mysql"
	gm "gorm.io/gorm"
	logger "gorm.io/gorm/logger"
	schema "gorm.io/gorm/schema"
)

var (
	db *gm.DB 
)

func GetDb() *gm.DB {
	return db 
}


func ConnectGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?multiStatements=true&parseTime=true",
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASS"),	
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	)

	var err error 
	db, err = gm.Open(mysql.Open(dsn), initConfig())
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to db")
	}

	dbQuery := fmt.Sprintf("create database if not exists %s", os.Getenv("DB_DATABASE"))
	db = db.Exec(dbQuery)

	useDb := fmt.Sprintf("use %s", os.Getenv("DB_DATABASE"))
	db = db.Exec(useDb)

	database, _ := db.DB() 
	database.SetMaxIdleConns(20)
	database.SetMaxOpenConns(20)
} 


func initConfig() *gm.Config {
	return &gm.Config{
		Logger: initLogger(),
		NamingStrategy: initNamingStrategy(),
	}
}


func initLogger() logger.Interface {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		Colorful: true,
		LogLevel: logger.Info,
		SlowThreshold: time.Second,
	})
	return newLogger
}

func initNamingStrategy() *schema.NamingStrategy {
	return &schema.NamingStrategy{
		SingularTable: false,
		TablePrefix: "",
	}
}