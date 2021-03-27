package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("go-publish-simple")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer c.Disconnect(250)

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)

		if token := c.Publish("go-mqtt/sample", 0, false, text); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	}
}
