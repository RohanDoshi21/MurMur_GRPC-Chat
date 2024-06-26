package service

import (
	"encoding/json"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var client MQTT.Client
var token MQTT.Token

func InitMQTT() {
	broker := "tcp://localhost:1883"

	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetClientID("go-mqtt-client")

	client = MQTT.NewClient(opts)

	if token = client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

type MQTTMessage struct {
	Message string
	Sender  string
}

func SendNotification(payload *NotifBody) {
	marshalPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	token = client.Publish(fmt.Sprintf("message/%s", payload.Receiver), 0, false, marshalPayload)
	token.Wait()
}
