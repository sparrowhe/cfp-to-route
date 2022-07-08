package service

import (
	"cfptoroute/global"
	"cfptoroute/internal/dao"
	"cfptoroute/internal/model"
	"strings"
)

type Segment struct {
	From string `json:"from"`
	Via  string `json:"via"`
	To   string `json:"to"`
}

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
	tmp := strings.Split(route, " ")
	for i := 0; i < len(tmp); i++ {
		if (i+1)%2 == 0 {
			if tmp[i] == "DCT" || (len(tmp[i]) == 5 && notIncludeNumbers(tmp[i])) || (len(tmp[i]) == 3 && notIncludeNumbers(tmp[i])) {
				segments = append(segments, Segment{From: tmp[i-1], Via: "DCT", To: tmp[i]})
			} else {
				segments = append(segments, Segment{From: tmp[i-1], To: tmp[i+1], Via: tmp[i]})
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
	WaypointDao := dao.GetWaypoint(global.DB)
	for _, seg := range s {
		if seg.Via == "DCT" {
			wp, err := WaypointDao.GetWaypointByName(seg.From)
			if dao.NotFound(err) {
				continue
			} else if err != nil {
				return []model.Waypoint{}, err
			}
			points = append(points, wp)
			wp, err = WaypointDao.GetWaypointByName(seg.To)
			if dao.NotFound(err) {
				continue
			} else if err != nil {
				return []model.Waypoint{}, err
			}
			points = append(points, wp)
		} else {
			wp, err := AirwayDao.GetAirwayByWhereToWhere(seg.Via, seg.From, seg.To)
			if len(wp) > 1 && wp[0].Point != s[0].From {
				wp = wp[1:]
			}
			if dao.NotFound(err) {
				continue
			} else if err != nil {
				return []model.Waypoint{}, err
			}
			for _, wp := range wp {
				points = append(points, model.Waypoint{
					Name:      wp.Point,
					Latitude:  wp.Latitude,
					Longitude: wp.Longitude,
				})
			}
		}
	}
	// 去重
	//points = RemoveDuplicate(points)
	return points, nil
}
