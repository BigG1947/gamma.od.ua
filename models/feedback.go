package models

import "time"

type FeedBack struct {
	Id    int64     `json:"id"`
	Name  string    `json:"name"`
	Text  string    `json:"text"`
	Date  time.Time `json:"date"`
	Check bool      `json:"check"`
}
