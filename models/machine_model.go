package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gobott-web/store"
	"gopkg.in/mgo.v2/bson"
)

type Machine struct {
	BaseModel
	Name				string                   `json:"name"`
	Sensors 			[]*AnalogSensor                `json:"sensors"`
	SensorIds 			[]bson.ObjectId          `json:"sensor_ids"`
	Instructions 		[]*Instruction           `json:"instructions"`
}

type MachineJson struct {
	BaseModel
	Name 				string   	         	 `json:"name"`
	Data 				[]uint8                  `json:"data"`
}

func NewMachine(name string) *Machine {
	m := new(Machine)
	m.Id = bson.NewObjectId()
	m.Name = name

	return m
}

func (m *Machine) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(m, "", "    ")
}

func (m *Machine) UnmarshalJson(data []byte) error {
	machine := &Machine{}
	if err := json.Unmarshal(data, &machine); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	return nil
}

func (m *Machine) Save() error {
	json, err := m.MarshalJson()
	store.AddToDb([]byte("machines"), json)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (m *Machine) AddSensor(s *AnalogSensor) error {
	m.Sensors = append(m.Sensors, s)
	//m.SensorIds = append(m.SensorIds, s.Id)

	return nil
}

