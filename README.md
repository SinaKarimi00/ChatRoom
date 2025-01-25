# Chatroom

A simple chatroom application built with Go, using NATS as the messaging system. This application allows users to join a
chatroom, send messages, view active users, and leave the chatroom.

---

## Instructions for Running the Application

### Prerequisites:

- Install [Docker](https://www.docker.com/)
- Install [Go](https://golang.org/) (version 1.23)
- Clone the repository to your local machine.

### Steps:

1. **Start the NATS Server**:
   ```bash
   docker-compose up
2. **Run the Server**:
   ```bash
   go run cmd/server/main.go
3. **Run the Client: Open another terminal and execute**:
   ```bash
   go run cmd/client/main.go

### Interact with the Chatroom

Once the client is running, you can interact with the chatroom using the following commands:

1. **Join the Chatroom**:
    - When prompted, enter a unique username to join the chatroom.

2. **Send a Message**:
    - Type a message (e.g., `Hello, everyone!`) and press Enter to broadcast it to all users in the chatroom.

3. **View Active Users**:
    - Type `#users` and press Enter to see the list of currently active users.

4. **Leave the Chatroom**:
    - Type `#leave` and press Enter to exit the chatroom gracefully. Other users will be notified that you have left.


