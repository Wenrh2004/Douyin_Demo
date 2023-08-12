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

for {
	case conn := <-Manager.Register:
		log.Println("A new socket has connected.",conn.ID)
		Manager.Clients[conn.ID] = conn
		replyMessage := &ReceiveMessages {
			Code:    e.WebSocketSuccess,
			Content: "connect success",
		}
		message, _ := json.Marshal(replyMessage)
		_ = conn.Socket.WriteMessage(websocket.TextMessage, message)
		
	case conn := <-Manager.Unregister:
		log.Println("A socket hasn't disconnected.",conn.ID)
		if _, ok := Manager.Clients[conn.ID]; ok {
			replyMessage := &ReplyMessages {
				Code:    e.WebSocketEnd,
				Context: "connect broken",
			}
			message, _ := json.Marshal(replyMessage)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, message)
			close(conn.Send)
			delete(Manager.Clients, conn.ID)
		}
}
