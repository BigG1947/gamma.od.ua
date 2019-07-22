package models

import "time"

type Video struct {
	Id   int64     `json:"id"`
	Src  string    `json:"src"`
	Date time.Time `json:"date"`
}
