package socket

import (
	"github.com/Mindgamesnl/wnal/process"
	"io"
)

type hub struct {
	clients map[int]*client
	dispatch chan interface{}
	quit chan struct{}
	register chan *client
	deregister chan *client
}

func NewHub(ch chan interface{}, quit chan struct{}) *hub {
	return &hub{
		clients:    make(map[int]*client),
		register:   make(chan *client),
		deregister: make(chan *client),
		quit:       quit,
		dispatch:   ch,
	}
}

func (h *hub) start() {
	for {
		select {
		case <-h.quit:
			return
		case message := <-h.dispatch:
			for _, client := range h.clients {
				client.ch <- message
			}
		case client := <-h.register:
			h.push(client)
			go h.watchDisconnect(client)
		case client := <-h.deregister:
			h.delete(client)
		}
	}
}

func (h *hub) watchDisconnect(client *client) {
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			client.Close()
			h.Deregister(client)
			return
		}
		io.WriteString(process.CommandWriter, string(message))
	}
}

func (h *hub) Register(c *client) {
	h.register <- c
}

func (h *hub) Deregister(c *client) {
	h.deregister <- c
}

func (h *hub) push(client *client) {
	h.clients[client.id] = client
}

func (h *hub) delete(client *client) {
	delete(h.clients, client.id)
}

func (h *hub) Iter(f func(*client)) {
	for _, client := range h.clients {
		f(client)
	}
}
