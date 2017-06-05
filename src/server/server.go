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
	conn, err := server.upgrader.Upgrade(write, read, nil)
	if err != nil {
		log.Fatal(err)
	}

	server.clients[conn] = true

	/*go func(conn *websocket.Conn) {
		for {
			m_type, msg, _ := conn.ReadMessage()
			log.Println("mtype: ", m_type, " msg: ", msg)
		}
	}(conn)*/
	
	go func (conn *websocket.Conn) {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Printf("error: %v", err)
				delete(server.clients, conn)
				break
			}
			log.Println(msg)
		}
		defer conn.Close()
	}(conn)
}

func CreateServer() Server{
	return Server{make(map[*websocket.Conn]bool), websocket.Upgrader{}}
}

