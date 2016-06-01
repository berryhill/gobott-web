package bot

import(
	"fmt"
	"encoding/json"

	"github.com/hybridgroup/gobot/platforms/mqtt"
	"github.com/hybridgroup/gobot"
	"github.com/gobott/models"
	//"time"
)

var mqttAdaptor *mqtt.MqttAdaptor
var owork func()

func init() {
	owork = func() {
		mqttAdaptor.On("bot_to_web", func(data []byte) {
			button := new(models.Button)
			err := json.Unmarshal(data, button)

			if err != nil {
				//fmt.Println(err)
				//return
			}


			fmt.Println(button)
			return
		})
/*
		gobot.Every(750*time.Millisecond, func() {
			button := new(models.Button)
			button.Name = "WebTester"
			json := button.MarshalJson()
			SendMessage(json)
			fmt.Print("Send Message")
			fmt.Println(json)
		})
*/
	}
}

func NewWebOperator() *gobot.Robot {
	mqttAdaptor = mqtt.NewMqttAdaptor("server", "tcp://test.mosquitto.org:1883", "web-operator")
	robot := gobot.NewRobot("WebOperator", []gobot.Connection{mqttAdaptor}, owork, )

	return robot
}

func SendMessage(b []byte) {
	mqttAdaptor.Publish("web_to_bot", b)
}
