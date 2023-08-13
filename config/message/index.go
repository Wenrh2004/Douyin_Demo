package message

import "golang.org/x/net/websocket"

// SendMessages ==> send message to client
type SendMessages struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// ReceiveMessages ==> receive message from client
type ReceiveMessages struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

// User ==> user information
type User struct {
	ID     string
	sendID string
	Socket *websocket.Conn
	Send   chan []byte
}

// Broadcast ==> broadcast message
type Broadcast struct {
	Client  *User
	Message []byte
	Type    int
}

// UserManager ==> user manager
type UserManager struct {
	Clients    map[string]*User
	Broadcast  chan *Broadcast
	Reply      chan *User
	Register   chan *User
	Unregister chan *User
}

// Message ==> message to JSON
type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

// Manager ==> user manager
var Manager = UserManager{
	Clients:    make(map[string]*User),
	Broadcast:  make(chan *Broadcast),
	Reply:      make(chan *User),
	Register:   make(chan *User),
	Unregister: make(chan *User),
}
