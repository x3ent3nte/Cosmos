package server

import (
	"log"
	"net/http"
    "github.com/gorilla/websocket"
)

type Client struct {
	code int
	conn *websocket.Conn
}

type Server struct {
	clients map[*Client]bool
	upgrader websocket.Upgrader
}

func (server *Server) Serve(data []string) {
	for client := range server.clients {
		client.conn.WriteJSON(data)
	}
}

func (server *Server) GetClientsData() []int {
	data := make([]int, len(server.clients))
	i := 0
	for client := range server.clients {
		data[i] = client.code 
		i++
	}
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

	client := Client{0, conn}
	server.clients[&client] = true
	
	go func (client *Client) {
		for {
			_, msg, err := client.conn.ReadMessage()
			if err != nil {
				log.Printf("error: %v", err)
				delete(server.clients, client)
				break
			}
			log.Println(msg)
		}
		defer conn.Close()
	}(&client)
}

func CreateServer() Server{
	return Server{make(map[*Client]bool), websocket.Upgrader{}}
}

