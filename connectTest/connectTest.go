package main

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"net/url"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"golang.org/x/net/websocket"
// )

// func main() {
// 	// WebSocket server URL with HTTPS
// 	serverURL := "wss://localhost:8080/query"

// 	// Parse the server URL
// 	_, err := url.Parse(serverURL)
// 	if err != nil {
// 		fmt.Println("Error parsing WebSocket URL:", err)
// 		return
// 	}

// 	// Configure TLS settings (use your actual certificate and key files)
// 	config, err := websocket.NewConfig(serverURL, "https://localhost")
// 	if err != nil {
// 		fmt.Println("Error creating WebSocket config:", err)
// 		return
// 	}
// 	config.TlsConfig = &tls.Config{}
// 	config.TlsConfig.InsecureSkipVerify = true // Ignore certificate verification (for testing only)
// 	config.Header.Add("auth", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiI3MmJiNjViMS1mYWI4LTRkNDAtODA3Yi0xZmRkNzExNTRiNmMiLCJwcm9kdWN0SUQiOiIiLCJleHAiOjE2OTQ2MTQwMjR9.rDsIlJk4GQRfl5KbvENR2uCqyHgDhVz_7MUVTDlDsyA")
// 	// config.Header.Add("Sec-WebSocket-Protocol", "graphql-transport-ws")
// 	// config.Header.Add("Sec-WebSocket-Protocol", "graphql-ws")
// 	config.Protocol = append(config.Protocol, "graphql-transport-ws")

// 	// Connect to the WebSocket server
// 	ws, err := websocket.DialConfig(config)
// 	if err != nil {
// 		fmt.Println("Error connecting to WebSocket server:", err)
// 		return
// 	}
// 	defer ws.Close()

// 	// Handle incoming messages in a separate goroutine
// 	fmt.Println(websocket.Message.Send(ws, `{"type":"connection_init"}`))

// 	go func() {
// 		for {
// 			fmt.Println("waiting ...")
// 			var message string
// 			if err := websocket.Message.Receive(ws, &message); err != nil {
// 				fmt.Println("Error receiving message:", err, ", message:", []byte(message))
// 				break
// 			}
// 			fmt.Println("Received:", message)
// 		}
// 	}()

// 	time.Sleep(time.Second)
// 	subscription := `{"payload":{"query":"subscription count {count}","operationName":"count"},"type":"subscribe"}`
// 	go func() {
// 		fmt.Println(websocket.Message.Send(ws, subscription))
// 	}()

// 	// Handle signals for graceful shutdown (Ctrl+C)
// 	sigCh := make(chan os.Signal, 1)
// 	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
// 	fmt.Println("waiting 2 ...")
// 	// Wait for a signal to gracefully exit
// 	<-sigCh
// }
