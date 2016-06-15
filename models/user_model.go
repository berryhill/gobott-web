package models

import (
 	"encoding/json"
	"log"

	"github.com/gobott-web/store"
)

type User struct {
	Name 		string                `json:"name"`
}

func NewUser(name string) *User {
	u := new(User)
	u.Name = name
	return u
}

func (u *User) MarshalJson() ([]byte, error) {
	json, err := json.Marshal(u)
	return json, err
}

func (u *User) Save() ([]byte, error) {
	json, err := u.MarshalJson()
	store.AddToDb([]byte("users"), json)

	if err != nil {
		log.Fatal(err)
		return []byte("ERROR"), err
	}

	return json, nil
}

