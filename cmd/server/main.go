package main

import (
	"ChatRoom/internal/adapters"
	"ChatRoom/internal/app"
	"ChatRoom/internal/domain"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	natsClient, err := adapters.NewNATSClient("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}

	chatroom := domain.NewChatroom()
	service := app.NewChatroomService(chatroom)

	_, err = natsClient.Subscribe("chatroom.join", func(msg *nats.Msg) {
		service.HandleJoin(string(msg.Data))
		fmt.Printf("User joined: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to join messages: %v", err)
	}

	_, err = natsClient.Subscribe("chatroom.leave", func(msg *nats.Msg) {
		service.HandleLeave(string(msg.Data))
		fmt.Printf("User left: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to leave messages: %v", err)
	}

	_, err = natsClient.Subscribe("chatroom", func(msg *nats.Msg) {
		fmt.Printf("Broadcasting message: %s\n", string(msg.Data))
		natsClient.Publish("chatroom.messages", msg.Data)
	})
	if err != nil {
		log.Fatalf("Error subscribing to chat messages: %v", err)
	}

	_, err = natsClient.Subscribe("chatroom.users", func(msg *nats.Msg) {
		users := service.HandleListUsers()
		response := fmt.Sprintf("Active users: %v", users)
		natsClient.Publish("chatroom.users.response."+string(msg.Data), []byte(response))
	})
	if err != nil {
		log.Fatalf("Error subscribing to user list requests: %v", err)
	}

	fmt.Println("Server is running...")
	select {}
}
