package builder

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	clients   = make(map[chan string]bool)
	clientsMu sync.Mutex
)

// HandleSSE manages Server-Sent Events for live reload
func HandleSSE(w http.ResponseWriter, r *http.Request) {
	setupSSEHeaders(w)

	messageChan := make(chan string)
	registerClient(messageChan)
	defer unregisterClient(messageChan)

	streamMessages(w, r, messageChan)
}

// setupSSEHeaders configures response headers for Server-Sent Events
func setupSSEHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

// registerClient adds a new SSE client
func registerClient(messageChan chan string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[messageChan] = true
}

// unregisterClient removes an SSE client
func unregisterClient(messageChan chan string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	delete(clients, messageChan)
	close(messageChan)
}

// streamMessages sends SSE messages to a client
func streamMessages(w http.ResponseWriter, r *http.Request, messageChan chan string) {
	for {
		select {
		case msg := <-messageChan:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
		case <-r.Context().Done():
			return
		}
	}
}

// NotifyClients broadcasts a message to all connected SSE clients
func NotifyClients(message string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for client := range clients {
		select {
		case client <- message:
		default:
			close(client)
			delete(clients, client)
		}
	}
}
