package models

import "gopkg.in/mgo.v2/bson"

type Alert interface {
	Run() error
	GetDescription() string
	GetMessage() string
}

type EmailAlert struct {
	BaseModel
	Name 		string                `json:"name"`
	Description 	string                `json:"description"`
	Message 	string                `json:"message"`
}

func NewEmailAlert(name string, description string, message string) *EmailAlert {
	ea := new(EmailAlert)
	ea.Id = bson.NewObjectId()
	ea.Name = name
	ea.Description = description
	ea.Message = message

	return ea
}

func (ea *EmailAlert) Run() error {
	//TODO implement
	return nil
}

func (ea *EmailAlert) GetDescription() string {
	return ea.Description
}

func (ea *EmailAlert) GetMessage() string {
	return ea.Message
}

