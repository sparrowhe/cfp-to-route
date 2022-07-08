package dao

import (
	"cfptoroute/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

type Airport struct {
	DB *gorm.DB
}

func GetAirport(db *gorm.DB) Airport {
	return Airport{db}
}

func (ap Airport) GetAirportByIcao(icao string) (model.Airport, error) {
	var ap2 model.Airport
	err := errors.WithStack(ap.DB.Where("icao = ?", icao).First(&ap2).Error)
	if err != nil {
		return ap2, err
	}
	return ap2, nil
}
