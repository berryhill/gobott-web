package models

import (
	"time"
	"log"

	"github.com/gobott-web/store"
	"encoding/json"
	"fmt"
)

type Report struct {
	Date 		time.Time                `json:"date"`
}

func NewReport() *Report {
	r := new(Report)
	r.Date = time.Now()

	return r
}

func (r *Report) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(r, "", "    ")
}

func (r *Report) UnmarshalJson(data []byte) error {
	report := &Report{}

	if err := json.Unmarshal(data, &report); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	return nil
}

func (r *Report) Save() error {
	json, err := r.MarshalJson()
	store.AddToDb([]byte("reports"), []byte("report"), json)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Report Saved")

	return nil
}

