package socket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

func getWsUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
}

var BroadcasterCh = make(chan interface{})

func StartSocket(port string) {
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()

	globalQuit := make(chan struct{})
	hub := NewHub(BroadcasterCh, globalQuit)

	defer close(globalQuit)

	go hub.start()

	r.GET("/status", wsHandler(getWsUpgrader(), globalQuit, hub))
	r.Run("localhost:" + port)
}

func wsHandler(wsupgrader *websocket.Upgrader, globalQuit chan struct{}, hub *hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		client := NewClient(conn)
		go client.handle()

		hub.Register(client)
	}
}

func Broadcast(data interface{}, BroadcasterCh chan interface{}) {
	BroadcasterCh <- data
}

