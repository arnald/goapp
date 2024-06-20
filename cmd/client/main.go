package main

import (
	"flag"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func main() {
	var numConnections int
	flag.IntVar(&numConnections, "n", 1, "number of parallel connections")
	flag.Parse()

	var wg sync.WaitGroup

	for i := 0; i < numConnections; i++ {
		wg.Add(1)
		go func(connID int) {
			defer wg.Done()
			connectToWebSocket(connID)
		}(i)
	}

	wg.Wait()
}

func connectToWebSocket(connID int) {
	url := "ws://localhost:8080/goapp/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Printf("[conn #%d] failed to connect: %v\n", connID, err)
		return
	}
	defer conn.Close()

	log.Printf("[conn #%d] connected", connID)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("[conn #%d] read error: %v\n", connID, err)
			return
		}
		log.Printf("[conn #%d] %s\n", connID, message)
	}
}
