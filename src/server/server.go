package server

import (
	"sync"
	"log"
	"net/http"
    "github.com/gorilla/websocket"
)

type Client struct {
	id int64
	keycode int
	conn *websocket.Conn
}

type Server struct {
	sync.RWMutex
	id_next int64
	clients map[*Client]bool
	upgrader websocket.Upgrader
}

func (server *Server) Serve(data []string) {
	for client := range server.clients {
		client.conn.WriteJSON(data)
	}
}

func (server *Server) AddClient(client *Client) {
	server.Lock()
	server.clients[client] = true
	server.Unlock()
}

func (server *Server) RemoveClient(client *Client) {
	server.Lock()
	delete(server.clients, client)
	server.Unlock()
}

func (server *Server) GetClientsData() map[int64]int {
	server.Lock()
	data := make(map[int64]int)
	for client := range server.clients {
		data[client.id] = client.keycode
	}
	server.Unlock()
	return data
}

func (server *Server) StartServer() {
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", server.handleConnections)

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (server *Server) handleConnections(write http.ResponseWriter, read *http.Request) {
	conn, err := server.upgrader.Upgrade(write, read, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := Client{server.id_next, 0, conn}
	server.id_next += 1
	server.AddClient(&client)
	
	go func (client *Client) {
		for {
			_, msg, err := client.conn.ReadMessage()
			if err != nil {
				log.Printf("error: %v", err)
				server.RemoveClient(client)
				break
			}
			client.keycode = int(msg[0])
		}
		defer conn.Close()
	}(&client)
}

func CreateServer() Server{
	return Server{sync.RWMutex{}, int64(0), make(map[*Client]bool), websocket.Upgrader{}}
}














