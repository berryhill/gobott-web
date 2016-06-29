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

func NewReport(name string, m *Machine) *Report {
	r := new(Report)
	r.Date = time.Now()
	r.Id = bson.NewObjectIdWithTime(r.Date)
	r.Name = name
	r.Machine = m

	return r
}

type ReportJson struct {
	BaseModel
	Name 			string            `json:"name"`
	Machine 		[]byte            `json:"machine"`
}

func (r *Report) MarshalJson() ([]byte, error) {
	var err error

	machineJson := new(MachineJson)
	machineJson.Id = r.Machine.Id
	machineJson.Name = r.Machine.Name

	reportJson := new(ReportJson)
	reportJson.Id = r.Id

	reportJson.Name = r.Name
	if reportJson.Machine, err = json.Marshal(machineJson); err != nil {
		return nil, fmt.Errorf("error unmarshaling machine: %v", err)
	}


	return json.MarshalIndent(reportJson, "", "    ")
}

func (r *Report) UnmarshalJson(data []byte) error {
	reportJson := new(ReportJson)
	if err := json.Unmarshal(data, &reportJson); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	r.Id = reportJson.Id
	r.Name = reportJson.Name
	if err := json.Unmarshal(reportJson.Machine, &r.Machine); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	//if err := json.Unmarshal(data, &r); err != nil {
	//	return fmt.Errorf("error unmarshaling report: %v", err)
	//}

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
