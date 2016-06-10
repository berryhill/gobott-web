package models

type Sensor interface {
	Set()
	Listen()
	MarshalJson()
	UnmarshhalJson()
	Save()
}
