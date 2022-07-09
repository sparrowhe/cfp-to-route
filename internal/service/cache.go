package service

import (
	"cfptoroute/global"
	"cfptoroute/internal/model"
)

func SaveCache(route string, points []model.Waypoint) {
	global.Cache[route] = points
}

func LoadCache(route string) []model.Waypoint {
	return global.Cache[route]
}
