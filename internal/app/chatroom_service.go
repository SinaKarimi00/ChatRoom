package app

import (
	"ChatRoom/internal/domain"
)

type ChatroomService struct {
	Chatroom *domain.Chatroom
}

func NewChatroomService(chatroom *domain.Chatroom) *ChatroomService {
	return &ChatroomService{Chatroom: chatroom}
}

func (cs *ChatroomService) HandleJoin(username string) {
	user := domain.NewUser(username)
	cs.Chatroom.Join(user)
}

func (cs *ChatroomService) HandleLeave(username string) {
	cs.Chatroom.Leave(username)
}

func (cs *ChatroomService) HandleMessage(username, message string) {
	formattedMessage := username + ": " + message
	cs.Chatroom.AddMessage(formattedMessage)
}

func (cs *ChatroomService) HandleListUsers() []string {
	return cs.Chatroom.ListUsers()
}
