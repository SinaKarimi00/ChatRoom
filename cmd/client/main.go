package main

import (
	"ChatRoom/internal/adapters"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	natsClient, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer natsClient.Close()

	cli := adapters.NewCLI()

	username := cli.ReadInput("Enter your username: ")

	err = natsClient.Publish("chatroom.join", []byte(username))
	if err != nil {
		log.Fatalf("Error publishing join message: %v", err)
	}

	_, err = natsClient.Subscribe("chatroom.messages", func(msg *nats.Msg) {
		cli.PrintOutput(string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to chat messages: %v", err)
	}

	_, err = natsClient.Subscribe("chatroom.users.response."+username, func(msg *nats.Msg) {
		cli.PrintOutput(string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to user list response: %v", err)
	}

	for {
		message := cli.ReadInput("Enter message (#users to list users, #leave to exit): ")

		if message == "#users" {
			err := natsClient.Publish("chatroom.users", []byte(username))
			if err != nil {
				cli.PrintOutput("Failed to request user list.")
			}
		} else if message == "#leave" {
			err := natsClient.Publish("chatroom.leave", []byte(username))
			if err != nil {
				cli.PrintOutput("Failed to notify server about leaving.")
			}
			cli.PrintOutput("You have left the chatroom.")
			break
		} else {
			err := natsClient.Publish("chatroom", []byte(username+": "+message))
			if err != nil {
				cli.PrintOutput("Failed to send message.")
			}
		}
	}
}
