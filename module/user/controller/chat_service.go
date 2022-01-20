package user_controller

import (
	"flashare/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Client struct {
	Connection *websocket.Conn
	Sender     string
	Receiver   string
}

var clients []Client

func HandleChatConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	clients = append(clients, Client{Connection: ws, Sender: c.Query("sender"), Receiver: c.Query("receiver")})
}

type Msg struct {
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	Content  string    `json:"content"`
	Time     time.Time `json:"time"`
}

func HandleMessage(c *gin.Context) {
	var msg Msg
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
	}

	//Code trong day ne ck iu

	//

	for i := 0; i < len(clients); {
		client := clients[i]
		if msg.Receiver == client.Sender && msg.Sender == client.Receiver {
			err := client.Connection.WriteJSON(msg)
			if err != nil {
				client.Connection.Close()
				clients = append(clients[:i], clients[i+1:]...)
				continue
			}
		}
		i++
	}
}
