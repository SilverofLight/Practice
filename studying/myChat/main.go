package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

//来自https://www.thepolyglotdeveloper.com/2016/12/create-real-time-chat-app-golang-angular-2-websockets/ 的项目
//写一个可以和服务端对话的客户端(只抄了golang部分)
/* 引入了go get github.com/gorilla/websocket
go get github.com/satori/go.uuid */

//程序中包含3个自定义结构体

/*
	ClientManager 将跟踪所有已连接的客户端、试图注册的客户端、

已被销毁并等待删除的客户端以及将向所有已连接的
客户端和从所有已连接的客户端广播的消息
*/
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// 每一个Client都有一个专用的id，socket连接和待发送的信息
type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// 定义一个全局变量massager
var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

// 需要创建三个goroutine，一个用来处理客户端，一个处理socket数据，一个编写socket数据
func (manager *ClientManager) start() {
	for { //死循环，一直处理信息
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected"})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected"})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

/*每次 manager. register 通道拥有数据时，
客户机都会被添加到由客户机管理器管理的可用客户机映射中。
添加客户端之后，将向所有其他客户端发送一条 JSON 消息，不包括刚刚连接的客户端*/
/*如果客户机由于任何原因断开连接，manager. unregister 通道将拥有数据。
断开连接的客户端中的通道数据将被关闭，客户端将从客户端管理器中删除。
将向所有剩余的连接发送一条消息，宣布套接字消失。*/
/*如果 manager. Broadcasting 通道拥有数据，那意味着我们正在尝试发送和接收消息。
我们希望循环遍历每个托管客户机，将消息发送给每个托管客户机。
如果由于某种原因，通道被阻塞或者消息无法发送，我们假设客户端已经断开连接，我们就移除它们。*/

// 为了保存重复的代码，创建了一个 manager. send 方法来遍历每个客户机:
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

//现在我们可以探索读取从客户端发送的 websocket 数据的 goroutine。
//这个 goroutine 程序的要点是读取套接字数据并将其添加到 manager. Broadcasting 以进行进一步的编排。

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func main() {
	fmt.Println("Starting application...")
	go manager.start()
	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":12345", nil)
}

func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	client := &Client{
		id:     uuid.NewV4().String(),
		socket: conn,
		send:   make(chan []byte),
	}

	manager.register <- client

	go client.read()
	go client.write()
}
