package models

import "time"

type Report struct {
	Date 		time.Time                `json:"date"`
}

func NewReport() *Report {
	r := new(Report)
	r.Date = time.Now()

	return r
}



