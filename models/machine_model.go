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
	//SensorIds 		[]bson.ObjectId          `json:"sensor_ids"`
	//Instructions 		[]*Instruction           `json:"instructions"`
}

type MachineJson struct {
	BaseModel
	Name 				string   	         	 `json:"name"`
	Sensors 			[][]byte                   `json:"sensors"`
	Data 				[]uint8                  `json:"data"`
}

func NewMachine(name string) *Machine {
	m := new(Machine)
	m.Id = bson.NewObjectId()
	m.Name = name

	return m
}

func (m *Machine) MarshalJson() ([]byte, error) {
	machineJson := &MachineJson{}
	machineJson.Id = m.Id
	machineJson.Name = m.Name

	for k := 0; k < len(m.Sensors); k++ {
		sensorJson, err := m.Sensors[k].MarshalJson()
		if err != nil {
			return []byte("ERROR"), err
		}

		machineJson.Sensors = append(machineJson.Sensors, sensorJson)
	}

	fmt.Println(machineJson)
	return json.MarshalIndent(m, "", "    ")
}

func (m *Machine) UnmarshalJson(data []byte) error {
	machineJson := new(MachineJson)
	if err := json.Unmarshal(data, &machineJson); err!= nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	m.Name = machineJson.Name
	m.Id = machineJson.Id

	machine := &Machine{}
	if err := json.Unmarshal(data, &machine); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	return nil
}

//func UnmarshalSensor(json []byte) {
//	var sensors []byte
//	err := json.Unmarshal(json, &sensors)
//}
//
//func UnmarshalSensors(json []byte) ([]*AnalogSensor, error) {
//	keysBody := []byte(`[{"id": 1,"key": "-"},{"id": 2,"key": "-"},{"id": 3,"key": "-"}]`)
//	keys := make([]PublicKey,0)
//	json.Unmarshal(keysBody, &keys)
//	fmt.Printf("%#v", keys)
//
//
//
//	sensors := make([]AnalogSensor, 0)
//	err := json.Unmarshal(json, &data)
//
//	sensors = append(sensors, )
//}

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

