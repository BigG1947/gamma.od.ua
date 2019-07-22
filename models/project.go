package models

import "time"

type Project struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Photos      []Photo   `json:"photos"`
	Date        time.Time `json:"date"`
}

type Photo struct {
	Id        int64     `json:"id"`
	Src       string    `json:"src"`
	Date      time.Time `json:"date"`
	IdProject int64     `json:"id_project"`
}
