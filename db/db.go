package db

import (
	"fmt"
	"go-graphql-2/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *config.Config) *gorm.DB {
	// dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=5432 dbname=%s sslmode=disable", config.DBHost, config.DBUsername, config.DBPassword, config.DBName)
	// dsn := fmt.Sprintf("host=localhost user=postgres password=1377 port=5432 dbname=sms2 sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("🚀 Failed to the Database, err message: ", err.Error())
		return nil
	}

	fmt.Println("🚀 Connected Successfully to the Database")
	return db
}
