package models

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

