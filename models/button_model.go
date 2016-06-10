package models

import (
	"encoding/json"
	"fmt"
)

type PushButton struct {
	BaseModel
	State 		bool                `json:"state"`
}

func NewPushButton(name string) *PushButton {
	pb := new(PushButton)
	pb.Name = name

	return pb
}

func (pb *PushButton) Set(state bool) {
	pb.State = state
}

func (pb *PushButton) Listen() bool {
	return pb.State
}

func (pb *PushButton) MarshalJson() ([]byte, error) {
	json, err := json.MarshalIndent(pb, "", "    ")

	if err != nil {
		fmt.Println(err)
		return json, err
	}

	return json, err
}

func (bp *PushButton) UnmarshalJson() {
	//TODO implement
}

