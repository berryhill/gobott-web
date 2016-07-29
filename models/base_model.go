package models

import "gopkg.in/mgo.v2/bson"

type BaseModel struct {
	Id 			bson.ObjectId         `json:"_id"`
	Name 		string                `json:"name"`
}

