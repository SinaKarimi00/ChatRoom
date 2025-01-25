package domain

type Chatroom struct {
	Users    map[string]bool
	Messages []string
}

func NewChatroom() *Chatroom {
	return &Chatroom{
		Users:    make(map[string]bool),
		Messages: []string{},
	}
}

func (c *Chatroom) Join(user *User) {
	c.Users[user.Username] = true
}

func (c *Chatroom) Leave(username string) {
	delete(c.Users, username)
}

func (c *Chatroom) AddMessage(message string) {
	c.Messages = append(c.Messages, message)
}

func (c *Chatroom) ListUsers() []string {
	users := []string{}
	for user := range c.Users {
		users = append(users, user)
	}
	return users
}
