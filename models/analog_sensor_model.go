package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gobott-web/store"
	"gopkg.in/mgo.v2/bson"
)

type AnalogSensor struct {
	BaseModel
	Value 		int32                `json:"value"`
	Peak		int32                `json:"peak"`
	Floor 		int32                `json:"floor"`
}

func NewAnalogSensor (name string) *AnalogSensor {
	as := new(AnalogSensor)
	as.Name = name
	as.Id = bson.NewObjectId()

	return as
}

func MakeAnalogSensor(mapp map[string]interface{}) *AnalogSensor {
	as := NewAnalogSensor("test")
	if val, ok := mapp["value"]; ok && val != nil {
		as.Value = int32(val.(float64))
	}
	if val, ok := mapp["peak"]; ok && val != nil {
		as.Peak = int32(val.(float64))
	}
	if val, ok := mapp["floor"]; ok && val != nil {
		as.Peak = int32(val.(float64))
	}

	return as
}

func (as *AnalogSensor) Set(value int32) {
	as.Value = value
}

func (as *AnalogSensor) Listen() int32 {
	return as.Value
}

func (as *AnalogSensor) MarshalJson() ([]byte, error) {
	json, err := json.MarshalIndent(as, "", "    ")

	if err != nil {
		fmt.Println(err)
		return json, err
	}

	return json, err
}

func (as *AnalogSensor) UnmarshalJson(data []byte) error {
	if err := json.Unmarshal(data, &as); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	return nil
}

func (as *AnalogSensor) Save() error {
	json, err := as.MarshalJson()
	store.AddToDb([]byte("Sensors"), json)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

