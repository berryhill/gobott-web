package mqtt

import (
	"fmt"
	"os"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"

	"github.com/gobott-web/models"
)

var MqttClient *MQTT.Client

var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

	HandleReport(msg)

}

func StartMqttClient() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://test.mosquitto.org:1883")
	opts.SetClientID("web_operator")
	opts.SetDefaultPublishHandler(f)

	MqttClient = MQTT.NewClient(opts)

	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := MqttClient.Subscribe("bot_to_web", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())

		os.Exit(1)
	}
}

func HandleReport (msg MQTT.Message) error {
	fmt.Println("Handling Report")
	report := new(models.Report)
	if err := report.UnmarshalJson(msg.Payload()); err != nil {
		fmt.Println("Error Unmarshalling")
		return fmt.Errorf("error unmarshaling report: %v", err)
	}
	fmt.Println("Report Id: ", report.Id)
	fmt.Println("Report Name: ", report.Name)
	fmt.Println("Machine Id: ", report.Machine.Id)
	fmt.Println("Machine Name: ", report.Machine.Name)
	fmt.Println("LightSensor Value: ", report.Machine.Sensors[0])
	fmt.Println("Thermistor Value: ", report.Machine.Sensors[1])

	if err := report.Save(); err != nil {
		fmt.Println("Error Saving")
		return fmt.Errorf("error saving report: %v", err)
	}

	return nil
}

func Send(message []byte) error {
	token := MqttClient.Publish("web_to_bot", 0, false, message)
	token.Wait()

	fmt.Println("Sending Message")

	return nil
}

