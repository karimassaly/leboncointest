package storage

import (
	"fmt"
	"os"

	"leboncointest/storage/endpointstore/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repositories struct {
	EndpointStatistics sql.EndpointRepository
}

func NewPostgresql() (*Repositories, error) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Paris",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return &Repositories{}, err
	}

	return &Repositories{
		EndpointStatistics: sql.NewEndpointStore(db),
	}, nil
}
