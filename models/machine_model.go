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
	Sensors 			[]*AnalogSensor          `json:"sensors"`
	//SensorIds 		[]bson.ObjectId          `json:"sensor_ids"`
	//Instructions 		[]*Instruction           `json:"instructions"`
}

type MachineJson struct {
	BaseModel
	Name 				string   	         	 `json:"name"`
	Sensors 			map[int][]byte  		 `json:"sensors"`
	//Data 				[]uint8                  `json:"data"`
}

func NewMachine(name string) *Machine {
	m := new(Machine)
	m.Id = bson.NewObjectId()
	m.Name = name

	return m
}

func (m *Machine) MarshalJson() ([]byte, error) {
	//var err error

	machineJson := &MachineJson{}
	machineJson.Id = m.Id
	machineJson.Name = m.Name

	var sensors map[int][]byte

	loop := 0
	for k := 0; k < len(m.Sensors); k++ {
		sensorJson, err := m.Sensors[k].MarshalJson()
		if err != nil {
			return []byte("ERROR"), err
		}

		sensors[loop] = sensorJson
		loop = loop + 1
	}

	machineJson.Sensors = sensors

	//machineJson.Sensors, err = json.Marshal(sensors)
	//if err != nil {
	//	return machineJson, err
	//}

	return json.MarshalIndent(machineJson, "", "    ")
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
		for val, _ := range machine_json_struct.Sensors {
			mapp := val.(map[string]interface{})
			temp_sensors = append(m.Sensors, MakeAnalogSensor(mapp))
		}

		m.Sensors = temp_sensors
	}

	return nil
}

	//err := m.UnmarshalSensors(machineJson.Sensors)
	//machine := &Machine{}
	//if err := json.Unmarshal(data, &machine); err != nil {
	//	return fmt.Errorf("error unmarshaling report: %v", err)
	//}
//}

//func (m *Machine) UnmarshalSensors(json []byte) error {
//	var sensors [][]byte
//	err := json.Unmarshal(json, &sensors)
//	if err != nil {
//		return err
//	}
//
//	return nil
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

