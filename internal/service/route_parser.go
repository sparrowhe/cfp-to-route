package service

import (
	"cfptoroute/global"
	"cfptoroute/internal/dao"
	"cfptoroute/internal/model"
	"regexp"
	"strings"
)

type Segment struct {
	From string `json:"from"`
	Via  string `json:"via"`
	To   string `json:"to"`
	Type uint   `json:"type"`
}

type EnumType uint

const (
	Airport uint = iota
	Waypoint
)

func notIncludeNumbers(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return false
		}
	}
	return true
}

func ParseCFPRoute(route string) []Segment {
	segments := []Segment{}
	SlashRegex := regexp.MustCompile(`\/.*$`)
	SlashRegex.ReplaceAllString(route, "")
	tmp := strings.Split(route, " ")
	for i := 0; i < len(tmp); i++ {
		if i == 0 {
			segments = append(segments, Segment{From: tmp[i], Type: Airport})
		} else if i == len(tmp)-1 {
			segments = append(segments, Segment{From: tmp[i], Type: Airport})
		} else if (i+2)%2 == 0 {
			if tmp[i] == "DCT" || (len(tmp[i]) == 5 && notIncludeNumbers(tmp[i])) || (len(tmp[i]) == 3 && notIncludeNumbers(tmp[i])) {
				segments = append(segments, Segment{From: tmp[i-1], Via: "DCT", To: tmp[i], Type: Waypoint})
			} else {
				segments = append(segments, Segment{From: tmp[i-1], To: tmp[i+1], Via: tmp[i], Type: Waypoint})
			}
		}
	}
	return segments
}

// func RemoveDuplicate(pl []model.Waypoint) []model.Waypoint {
// 	// 除去重复的点
// 	points := make([]model.Waypoint, 0)
// 	for _, p := range pl {
// 		if !Contains(points, p) {
// 			points = append(points, p)
// 		}
// 	}
// 	return points
// }

// func Contains(points []model.Waypoint, point model.Waypoint) bool {
// 	for _, p := range points {
// 		if p.Name == point.Name {
// 			return true
// 		}
// 	}
// 	return false
// }

func SegmentToPointsList(s []Segment) ([]model.Waypoint, error) {
	points := make([]model.Waypoint, 0)
	AirwayDao := dao.GetAirway(global.DB)
	AirportDao := dao.GetAirport(global.DB)
	WaypointDao := dao.GetWaypoint(global.DB)
	var id uint = 0
	for _, seg := range s {
		if seg.Via == "DCT" {
			wp, err := WaypointDao.GetWaypointByName(seg.From)
			if dao.NotFound(err) {
				continue
			} else if err != nil {
				return []model.Waypoint{}, err
			}
			wp.Id = id
			id += 1
			points = append(points, wp)
			wp, err = WaypointDao.GetWaypointByName(seg.To)
			if dao.NotFound(err) {
				continue
			} else if err != nil {
				return []model.Waypoint{}, err
			}
			wp.Id = id
			id += 1
			points = append(points, wp)
		} else if seg.Type == Airport {
			wp, err := AirportDao.GetAirportByIcao(seg.From)
			if dao.NotFound(err) {
				continue
			} else if err != nil {
				return []model.Waypoint{}, err
			}
			points = append(points, model.Waypoint{
				Id:        id,
				Name:      wp.Icao,
				Latitude:  wp.Latitude,
				Longitude: wp.Longitude,
			})
			id += 1
		} else {
			wp, err := AirwayDao.GetAirwayByWhereToWhere(seg.Via, seg.From, seg.To)
			if len(wp) > 1 && wp[0].Point != s[1].From {
				wp = wp[1:]
			}
			if dao.NotFound(err) {
				continue
			} else if err != nil {
				return []model.Waypoint{}, err
			}
			for _, wp := range wp {
				points = append(points, model.Waypoint{
					Id:        id,
					Name:      wp.Point,
					Latitude:  wp.Latitude,
					Longitude: wp.Longitude,
				})
				id += 1
			}
		}
	}
	// 去重
	//points = RemoveDuplicate(points)
	return points, nil
}
