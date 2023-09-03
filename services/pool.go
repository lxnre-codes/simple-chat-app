package services

import "fmt"

type Pool struct {
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
	broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.register:
			pool.clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.clients))
			for client := range pool.clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
		case client := <-pool.unregister:
			delete(pool.clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.clients))
			for client := range pool.clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
		case message := <-pool.broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client := range pool.clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
