package bot

import (
	"github.com/hybridgroup/gobot"
)

//type Bot struct {
//
//}

func NewBot() {
	//bot := new(Bot)
	gbot := gobot.NewGobot()

	gbot.AddRobot(NewWebOperator())
	gbot.Start()
}

