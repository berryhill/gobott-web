package models

type Instruction interface {
	GetName() string
	GetDescription() string
}

type TrueInstruction struct {
	BaseModel
	Description 		string                `json:"description"`
	Condition 		bool                `json:"condition"`
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
