package server

import (
	"log"
	"net/http"
    "github.com/gorilla/websocket"
)

type Server struct {
	clients map[*websocket.Conn]bool
	upgrader websocket.Upgrader
}

func (server *Server) Serve(data []string) {
	for client := range server.clients {
		client.WriteJSON(data)
	}
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
	ws, err := server.upgrader.Upgrade(write, read, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	server.clients[ws] = true

	for {
		var msg Message 
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(server.clients, ws)
			break
		}
	}
}

func CreateServer() Server{
	return Server{make(map[*websocket.Conn]bool), websocket.Upgrader{}}
}

type Message struct {}


