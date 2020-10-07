package socket

import (
	"github.com/Mindgamesnl/wnal/queue"
	"github.com/gorilla/websocket"
	"log"
)

var NextID int

type client struct {
	id   int
	conn *websocket.Conn
	ch   chan interface{}
	quit chan struct{}
}

func NewClient(conn *websocket.Conn) *client {
	NextID++
	return &client{
		id:   NextID,
		conn: conn,
		quit: make(chan struct{}),
		ch:   make(chan interface{}),
	}
}

func (c *client) Close() {
	close(c.quit)
}

func (c *client) handle() {

	for i := range queue.LogLines.Imports {
		c.conn.WriteJSON(MakeOutNormal(queue.LogLines.Imports[i].Text))
	}

	for {
		select {
		case <-c.quit:
			if err := c.conn.Close(); err != nil {
				log.Printf("client %d connection close error %v\n", c.id, err)
			}
			return
		case n := <-c.ch:
			jmsg := n
			if err := c.conn.WriteJSON(jmsg); err != nil {
				log.Println("ws write error:", err)
				return
			}
		}
	}

}
