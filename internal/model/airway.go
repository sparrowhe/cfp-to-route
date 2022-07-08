package model

type Airway struct {
	Id        uint    `json:"id"`
	LegId     uint    `json:"leg_id"`
	Name      string  `json:"name"`
	Point     string  `json:"point"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
