package models

import "encoding/json"

type Person struct {
	Name 		string                `json:"name"`
}

func NewPerson(name string) *Person {
	p := new(Person)
	p.Name = name
	return p
}

func (p *Person) MarshalJson() ([]byte, error) {
	json, err := json.Marshal(p)
	return json, err
}

func (p *Person) Save() {
	//TODO implement
}

