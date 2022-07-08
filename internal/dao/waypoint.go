package dao

import (
	"cfptoroute/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Waypoint struct {
	DB *gorm.DB
}

func GetWaypoint(db *gorm.DB) Waypoint {
	return Waypoint{db}
}

func (w Waypoint) GetWaypointByName(name string) (model.Waypoint, error) {
	var wp model.Waypoint
	err := errors.WithStack(w.DB.Where("name = ?", name).First(&wp).Error)
	if err != nil {
		return wp, err
	}
	return wp, nil
}

func (w Waypoint) GetWaypointByLatLon(lat, lon float64) (model.Waypoint, error) {
	var wp model.Waypoint
	err := errors.WithStack(w.DB.Where("latitude = ? AND longitude = ?", lat, lon).First(&wp).Error)
	if err != nil {
		return wp, err
	}
	return wp, nil
}
