package models

type Sensor interface {
	Set(int32)
	Listen() int32
	MarshalJson() ([]byte, error)
	UnmarshhalJson(data []byte) error
	Save() error
}
