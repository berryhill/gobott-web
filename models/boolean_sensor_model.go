package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gobott-web/store"
)

type BooleanSensor struct {
	BaseModel
	State 				int32             `json:"state"`
	TrueValue			int32             `json:"true_value"`
	FalseValue 			int32             `json:"false_value"`
}

func NewBooleanSensor(name string) *BooleanSensor {
	bs := new(BooleanSensor)
	bs.Name = name
	bs.TrueValue = 1
	bs.FalseValue = 0

	return bs
}

func (bs *BooleanSensor) Set(state int32) {
	bs.State = state
}

func (bs *BooleanSensor) Listen() int32 {
	return bs.State
}

func (bs *BooleanSensor) MarshalJson() ([]byte, error) {
	json, err := json.MarshalIndent(bs, "", "    ")

	if err != nil {
		fmt.Println(err)
		return json, err
	}

	return json, err
}

func (bs *BooleanSensor) UnmarshalJson(data []byte) error {
	if err := json.Unmarshal(data, &bs); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	return nil
}

func (bs *BooleanSensor) Save() error {
	json, err := bs.MarshalJson()
	store.AddToDb([]byte("Sensors"), json)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
