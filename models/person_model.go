package models

import (
 	"encoding/json"
	"log"

	"github.com/gobott-web/store"
)

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

func (p *Person) Save() ([]byte, error) {
	json, err := p.MarshalJson()
	store.AddToDb([]byte("people"), json)

	if err != nil {
		log.Fatal(err)
		return []byte("ERROR"), err
	}

	return json, nil
}

