package domain

import "testing"

func TestChatroom(t *testing.T) {
	chatroom := NewChatroom()

	user := NewUser("Sina")
	chatroom.Join(user)

	if len(chatroom.Users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(chatroom.Users))
	}

	chatroom.Leave("Sina")
	if len(chatroom.Users) != 0 {
		t.Errorf("Expected 0 users, got %d", len(chatroom.Users))
	}

	chatroom.AddMessage("Hello, world!")
	if len(chatroom.Messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(chatroom.Messages))
	}
}
