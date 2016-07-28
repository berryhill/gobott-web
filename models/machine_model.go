package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gobott-web/store"
	"gopkg.in/mgo.v2/bson"
	//"golang.org/x/mobile/exp/sensor"
)

type Machine struct {
	BaseModel
	Name				string                   `json:"name"`
	Sensors 			[]*AnalogSensor          `json:"sensors"`
	//SensorIds	 		[]bson.ObjectId          `json:"sensor_ids"`
	//Instructions 		[]*Instruction           `json:"instructions"`
}

func NewMachine(name string) *Machine {
	m := new(Machine)
	m.Id = bson.NewObjectId()
	m.Name = name

	return m
}

func MakeMachine(mapp map[string]interface{}) *Machine {
	m := NewMachine("Test")
	if val, ok := mapp["id"]; ok && val != nil {
	m.Name = val.(string)
	}
	if val, ok := mapp["name"]; ok && val != nil {
		m.Name = val.(string)
	}
	if val, ok := mapp["sensors"]; ok && val != nil {
		var sensors []*AnalogSensor
		sensor_interfaces := val.([]interface{})

		for _, sensor := range sensor_interfaces {
			sensors = append(sensors, MakeAnalogSensor(sensor.(map[string]interface{})))
		}

		m.Sensors = sensors
	}

	return m
}

func (m *Machine) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(m, "", "    ")
}

func (m *Machine) UnmarshalJson(data []byte) error {
	machine_json_struct := struct {
		BaseModel
		Name				string                   `json:"name"`
		Sensors 			[]interface{}            `json:"sensors"`
	}{}

	if err := json.Unmarshal(data, &machine_json_struct); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	m.Name = machine_json_struct.Name
	m.Id = machine_json_struct.Id

	if len(machine_json_struct.Sensors) > 0 {
		var temp_sensors []*AnalogSensor
		for _, val := range machine_json_struct.Sensors {
			mapp := val.(map[string]interface{})
			temp_sensors = append(temp_sensors, MakeAnalogSensor(mapp))
		}

		m.Sensors = temp_sensors
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

