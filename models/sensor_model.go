package models

type Sensor interface {
	Set()
	Listen()
	MarshalJson()
	UnmarshhalJson()
}


type AnalogSensor struct {
	BaseModel
	Value 		int32                `json:"value"`
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

func (as *AnalogSensor) MarshalJson() {
	//TODO implement
}

func (as *AnalogSensor) UnmarshalJson() {
	//TODO implement
}


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

func (pb *PushButton) MarshalJson() {
	//TODO implement
}

func (bp *PushButton) UnmarshalJson() {
	//TODO implement
}

