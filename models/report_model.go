package models

import (
	"time"
	"log"
	"encoding/json"
	"fmt"

	"github.com/gobott-web/store"
	"gopkg.in/mgo.v2/bson"
)

type Report struct {
	BaseModel
	Date 		time.Time               `json:"date"`
	Machine 	*Machine           	`json:"machine"`
}

func NewReport(m *Machine) *Report {
	r := new(Report)
	r.Date = time.Now()
	r.Id = bson.NewObjectIdWithTime(r.Date)
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
