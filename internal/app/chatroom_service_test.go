package app

import (
	"ChatRoom/internal/domain"
	"testing"
)

func TestChatroomService(t *testing.T) {
	chatroom := domain.NewChatroom()
	service := NewChatroomService(chatroom)

	service.HandleJoin("Sina")
	if len(chatroom.Users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(chatroom.Users))
	}

	service.HandleMessage("Sina", "Hi!")
	if len(chatroom.Messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(chatroom.Messages))
	}

	users := service.HandleListUsers()
	if len(users) != 1 || users[0] != "Sina" {
		t.Errorf("Expected user list to contain 'Sina', got %v", users)
	}
}
