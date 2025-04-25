package player

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Player struct {
	Conn *websocket.Conn
}

var upgrader = websocket.Upgrader{}

func CreateWebSocket(c *gin.Context) (*Player, error) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return nil, err
	}

	player := &Player{
		Conn: conn,
	}

	return player, nil
}
