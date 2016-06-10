package models

import (
	"time"
	"log"
	"encoding/json"
	"fmt"

	"github.com/gobott-web/store"
)

type Report struct {
	Date 		time.Time               `json:"date"`
	Machine 	*Machine           	`json:"machine"`
}

func NewReport(m *Machine) *Report {
	r := new(Report)
	r.Date = time.Now()
	r.Machine = m

	return r
}

func (r *Report) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(r, "", "    ")
}

func (r *Report) UnmarshalJson(data []byte) error {
	if err := json.Unmarshal(data, &r); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	return nil
}

func (r *Report) Save() error {
	json, err := r.MarshalJson()
	store.AddToDb([]byte("reports"), json)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
