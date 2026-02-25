package network

import (
	"log"
	"net/http"
	"recreational-actitude/game"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for the prototype
	},
}

type Server struct {
	clients   map[*websocket.Conn]bool
	broadcast chan game.Event
	mu        sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan game.Event),
	}
}

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	s.mu.Lock()
	s.clients[ws] = true
	s.mu.Unlock()

	log.Println("New client connected")

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			s.mu.Lock()
			delete(s.clients, ws)
			s.mu.Unlock()
			log.Println("Client disconnected")
			break
		}
	}
}

func (s *Server) HandleMessages() {
	for {
		msg := <-s.broadcast
		s.mu.Lock()
		for client := range s.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(s.clients, client)
			}
		}
		s.mu.Unlock()
	}
}

func (s *Server) BroadcastEvent(event game.Event) {
	s.broadcast <- event
}

func (s *Server) Start(addr string) {
	http.HandleFunc("/ws", s.HandleConnections)
	go s.HandleMessages()
	log.Printf("WebSocket server started on %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
