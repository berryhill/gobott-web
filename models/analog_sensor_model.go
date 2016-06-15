package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gobott-web/store"
)

type AnalogSensor struct {
	BaseModel
	Value 		int32                `json:"value"`
	Peak		int32                `json:"peak"`
	floor 		int32                `json:"floor"`
}

func NewAnalogSensor (name string) *AnalogSensor {
	as := new(AnalogSensor)
	as.Name = name

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
