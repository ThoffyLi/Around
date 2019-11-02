package entity

import (
	"encoding/json"
)

type Location struct {
	Lat float64 `json:"lat"` // -90 ~ 90
	Lon float64 `json:"lon"` // -180 ~ 180
}

type Post struct {
	User     string   `json:"user"`
	Message  string   `json:"message"`
	Location Location `json:"location"`
	Url      string   `json:"url"`
	Type     string   `json:"type"`
}
