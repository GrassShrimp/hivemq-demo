package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var msgRcvd mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", message.Topic())
	fmt.Printf("MSG: %s\n", message.Payload())
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("go-subscribe-simple").SetDefaultPublishHandler(msgRcvd)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer (func() {
		if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		c.Disconnect(250)
	})()

	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	select {}
}
