package bot

import(
	"fmt"
	"encoding/json"

	"github.com/hybridgroup/gobot/platforms/mqtt"
	"github.com/hybridgroup/gobot"
	"github.com/gobott/models"
)

var mqttAdaptor *mqtt.MqttAdaptor
var owork func()

func init() {
	owork = func() {
		mqttAdaptor.On("data", func(data []byte) {
			d := json.Unmarshal(data, models.Button{})
			fmt.Println(d)
		})

		//gobot.Every(1*time.Second, func() {
		//	json := Buttons[0].MarshalJson()
		//	SendMessage(json)
		//})
	}
}

func NewOperator() *gobot.Robot {
	mqttAdaptor = mqtt.NewMqttAdaptor("server", "tcp://test.mosquitto.org:1883", "pinger")
	robot := gobot.NewRobot("mqttBot", []gobot.Connection{mqttAdaptor}, owork, )

	return robot
}

func SendMessage(b []byte) {
	mqttAdaptor.Publish("data", b)
}