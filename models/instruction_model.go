package models

type Instruction interface {
	GetName() string
	GetDescription() string
}

type BaseInstruction struct {
	Description 		string                `json:"description"`
	Condition 		bool                  `json:"condition"`
	Alerts 			[]*Alert              `json:"alerts"`
}

type TrueInstruction struct {
	BaseModel
	BaseInstruction
}

func (ti *TrueInstruction) GetName() string {
	return ti.Name
}

func (ti *TrueInstruction) GetDescription() string {
	return ti.Description
}

func (ti *TrueInstruction) Evaluate() bool {
	return ti.Condition
}
