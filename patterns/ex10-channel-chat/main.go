package main

import (
	"fmt"
	"time"
)

type Hub struct {
	clients              map[*Client]bool
	register, unregister chan *Client
	messages             chan string
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		messages:   make(chan string, 10),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			if _, ok := h.clients[client]; ok {
				break
			}

			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; !ok {
				break
			}

			close(client.recv)
			delete(h.clients, client)
		case message := <-h.messages:
			for client := range h.clients {
				select {
				case client.recv <- message:
				default:
					close(client.recv)
					delete(h.clients, client)
				}
			}
		}
	}
}

type Client struct {
	recv chan string
}

func NewClient() *Client {
	return &Client{
		recv: make(chan string),
	}
}

func (c *Client) Run() {
	go func() {
		for message := range c.recv {
			fmt.Println(message)
		}
	}()
}

func main() {
	hub := NewHub()
	go hub.Run()

	client1 := NewClient()
	client1.Run()
	hub.register <- client1

	client2 := NewClient()
	client2.Run()
	hub.register <- client2

	time.Sleep(3 * time.Second)
	hub.messages <- "hello"
	time.Sleep(3 * time.Second)
}
