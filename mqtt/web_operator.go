package mqtt

import (
	"fmt"
	"os"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"

	"github.com/gobott-web/models"
)

var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

	if msg.Topic() == "bot_to_web_report" {
		HandleReport(msg)
	}
}

func StartMqttClient() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://test.mosquitto.org:1883")
	opts.SetClientID("web_operator")
	opts.SetDefaultPublishHandler(f)

	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("bot_to_web", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())

		os.Exit(1)
	}

/*
	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}

	time.Sleep(3 * time.Second)

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
*/
}

func HandleReport (msg MQTT.Message) error {
	report := new(models.Report)

	if err := report.UnmarshalJson(msg.Payload()); err != nil {
		return fmt.Errorf("error unmarshaling report: %v", err)
	}

	if err := report.Save(); err != nil {
		return fmt.Errorf("error saving report: %v", err)
	}

	return nil
}


