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
	Date 			time.Time            `json:"date"`
	Machine 		*Machine	         `json:"machine"`
}

func NewReport(m *Machine) *Report {
	r := new(Report)
	r.Date = time.Now()
	r.Id = bson.NewObjectIdWithTime(r.Date)
	r.Machine = m

	return r
}

func (r *Report) MarshalJson() ([]byte, error) {
	machineJson := new(MachineJson)
	machineJson.Id = r.Machine.Id
	machineJson.Name = r.Machine.Name

	return json.MarshalIndent(r, "", "    ")
}

func (r *Report) UnmarshalJson(data []byte) error {
	machineJson := new(MachineJson)
	if err := json.Unmarshal(data, &machineJson); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	r.Id = machineJson.Id
	r.Name = machineJson.Name

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
