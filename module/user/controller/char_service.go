package user_controller

import (
	"flashare/entity"
	"flashare/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Client struct {
	Connection *websocket.Conn
	UserID     string
}

var clients []Client

func HandleChatConnection(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	clients = append(clients, Client{Connection: ws, UserID: c.Param("userid")})

}

func HandleMessage(c *gin.Context) {
	var msg entity.Message
	if err := c.ShouldBind(&msg); err != nil {
		c.JSON(http.StatusBadRequest, utils.DataResponse{
			Success: false,
			Data:    err.Error(),
		})
	}
	for _, client := range clients {
		if client.UserID == msg.Receiver {

		}
	}
}
