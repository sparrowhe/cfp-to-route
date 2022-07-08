package model

type Airport struct {
	Id        uint    `json:"id"`
	Icao      string  `json:"icao"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
