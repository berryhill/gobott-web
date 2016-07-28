package models

import (
	"time"
	"log"
	"encoding/json"
	"fmt"

	//"github.com/gobott-web/store"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
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

type ReportJson struct {
	BaseModel
	Machine 		[]byte            `json:"machine"`
}

func (r *Report) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(r, "", "    ")
}

func (r *Report) UnmarshalJson(data []byte) error {
	report_json_struct := struct {
		BaseModel
		Date 				time.Time            `json:"date"`
		Machine 			interface{}          `json:"sensors"`
	}{}

	if err := json.Unmarshal(data, &report_json_struct); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	fmt.Println(report_json_struct.Machine)
	//r.Id = report_json_struct.Id
	//r.Name = report_json_struct.Name
	//mapp := report_json_struct.Machine.(map[string]interface{})
	//r.Machine = MakeMachine(mapp)

	fmt.Println(report_json_struct)

	return nil
}

func (r *Report) Save() error {
	session, err := mgo.Dial("localhost:27017")
	defer session.Close()

	if err != nil {
		log.Fatal(err)
	}

	c := session.DB("test").C("reports")
	err = c.Insert(r)
	if err != nil {
		log.Fatal(err)
	}

	session.Close()

	return nil
}
