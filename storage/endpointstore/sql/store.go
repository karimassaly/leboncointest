package sql

import (
	"errors"

	"gorm.io/gorm"
)

var ErrStatisticsNotFound = errors.New("statistics not found")

type EndpointStatistics struct {
	ID     int `gorm:"primary_key"`
	Route  string
	Method string
	Count  int64 `gorm:"default:1"`
}

type EndpointRepository interface {
	GetAll() ([]EndpointStatistics, error)
	CountUp(es EndpointStatistics) error
}

type Endpointstore struct {
	db *gorm.DB
}

var _ EndpointRepository = &Endpointstore{}

func NewEndpointStore(db *gorm.DB) *Endpointstore {
	db.AutoMigrate(EndpointStatistics{})
	return &Endpointstore{db}
}

func (f *Endpointstore) GetAll() ([]EndpointStatistics, error) {
	var stats []EndpointStatistics

	if err := f.db.Debug().Model(EndpointStatistics{}).Find(&stats).Error; err != nil {
		return []EndpointStatistics{}, ErrStatisticsNotFound
	}

	return stats, nil
}

func (f *Endpointstore) CountUp(es EndpointStatistics) error {
	var existing EndpointStatistics

	if err := f.db.Where("route = ? AND method = ?", es.Route, es.Method).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return f.db.Create(&es).Error
		}
		return err
	}

	existing.Count += 1
	return f.db.Model(&existing).Update("count", existing.Count).Error
}
