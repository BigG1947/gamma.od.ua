package models

import "time"

type News struct {
	Id          int64     `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	Text        string    `json:"text"`
	Images      string    `json:"images"`
	Date        time.Time `json:"date"`
	CountSee    int       `json:"count_see"`
}
