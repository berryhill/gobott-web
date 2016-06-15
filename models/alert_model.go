package models

type Alert interface {
	Run() error
	GetDescription() (string, error)
}

type EmailAlert struct {
	BaseModel
	Name 		string                `json:"name"`
}

func (ea *EmailAlert) Run() error {
	//TODO implement
	return nil
}

func (ea *EmailAlert) GetDescription() (string, error) {
	//TODO implement
	return "", nil
}
