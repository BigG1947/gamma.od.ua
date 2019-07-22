package models

import "time"

type Video struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Src         string    `json:"src"`
	Date        time.Time `json:"date"`
}
