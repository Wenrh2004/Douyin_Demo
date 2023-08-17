package controller

import (
	"Douyin_Demo/common"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

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
	fromID string
	toID   string
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
	Receiver   chan *User
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
	Receiver:   make(chan *User),
	Register:   make(chan *User),
	Unregister: make(chan *User),
}

func CreatUserID(fromUserID, toUserID string) string {
	return fromUserID + "->" + toUserID
}

func GetMessage(ctx *gin.Context) {
	fromUserID := ctx.Query("from_user_id")
	toUserID := ctx.Query("to_user_id")

	// Resolve cross-domain issues with the ws protocol
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(router *http.Request) bool {
			return true
		}}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	// Create user instances
	user := &User{
		fromID: CreatUserID(fromUserID, toUserID),
		toID:   CreatUserID(toUserID, fromUserID),
		Socket: conn,
		Send:   make(chan []byte),
	}

	// Register user
	Manager.Register <- user

	// Start read and write concurrent program to receive messages.
	go user.Read()
	go user.Write()
}

// Read() ==> read message from chat
func (ctx *User) Read() {
	defer func() {
		Manager.Unregister <- ctx
		_ = ctx.Socket.Close()
	}()

	for {
		ctx.Socket.PongHandler()
		sendMessages := new(SendMessages)
		err := ctx.Socket.ReadJSON(&sendMessages)
		if err != nil {
			fmt.Println("数据格式不正确", err.Error())
			Manager.Unregister <- ctx
			_ = ctx.Socket.Close()
			break
		}
		if sendMessages != nil {
			common.RedisClient.Incr(ctx.fromID)

			// Expiration date (of document)
			_, _ = common.RedisClient.Expire(ctx.fromID, time.Hour*60*24).Result()

			// Send message to broadcast channel
			Manager.Broadcast <- &Broadcast{
				Client:  ctx,
				Message: []byte(sendMessages.Content), // Send message
			}
		}
	}
}

// Write() ==> write message to chat
func (ctx *User) Write() {
	defer func() {
		_ = ctx.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-ctx.Send:
			if !ok {
				_ = ctx.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			receiveMessages := ReceiveMessages{
				Code:    200,
				Content: fmt.Sprint("%s", string(message)),
			}
			msg, _ := json.Marshal(receiveMessages)
			_ = ctx.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
