package main

// import (
// 	"context"
// 	"crypto/tls"
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/Khan/genqlient/graphql"
// 	"golang.org/x/net/websocket"
// )

// type MyDialer struct {
// 	dialer *net.Dialer
// }

// type MyConn struct {
// 	conn *websocket.Conn
// }

// func (c MyConn) ReadMessage() (messageType int, p []byte, err error) {
// 	if err := websocket.Message.Receive(c.conn, &p); err != nil {
// 		return websocket.UnknownFrame, nil, err
// 	}
// 	return messageType, p, err
// }

// func (c MyConn) WriteMessage(_ int, data []byte) error {
// 	err := websocket.Message.Send(c.conn, data)
// 	return err
// }

// func (c MyConn) Close() error {
// 	c.conn.Close()
// 	return nil
// }

// func (md *MyDialer) DialContext(ctx context.Context, urlStr string, requestHeader http.Header) (graphql.WSConn, error) {
// 	if md.dialer == nil {
// 		return nil, fmt.Errorf("nil dialer")
// 	}
// 	config, err := websocket.NewConfig(urlStr, "https://localhost")
// 	if err != nil {
// 		fmt.Println("Error creating WebSocket config:", err)
// 		return nil, err
// 	}
// 	config.TlsConfig = &tls.Config{}
// 	config.TlsConfig.InsecureSkipVerify = true // Ignore certificate verification (for testing only)
// 	config.Dialer = md.dialer
// 	config.Protocol = append(config.Protocol, "graphql-transport-ws")

// 	// Connect to the WebSocket server
// 	conn, err := websocket.DialConfig(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return graphql.WSConn(MyConn{conn: conn}), err
// }

// func main() {
// 	var err error
// 	defer func() {
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
// 	}()

// 	key := os.Getenv("GITHUB_TOKEN")
// 	if key == "" {
// 		err = fmt.Errorf("must set GITHUB_TOKEN=<github token>")
// 		return
// 	}

// 	headers := http.Header{}
// 	headers.Add("Authorization", "bearer "+key)
// 	headers.Add("auth", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiI3MmJiNjViMS1mYWI4LTRkNDAtODA3Yi0xZmRkNzExNTRiNmMiLCJwcm9kdWN0SUQiOiIiLCJleHAiOjE2OTQ2MTQwMjR9.rDsIlJk4GQRfl5KbvENR2uCqyHgDhVz_7MUVTDlDsyA")

// 	graphqlClient := graphql.NewClientUsingWebSocket(
// 		"wss://localhost:8080/query",
// 		&MyDialer{dialer: &net.Dialer{}},
// 		headers,
// 	)

// 	respChan, errChan, err := count(context.Background(), graphqlClient)
// 	if err != nil {
// 		return
// 	}
// 	defer graphqlClient.CloseWebSocket()
// 	for loop := true; loop; {
// 		select {
// 		case msg, more := <-respChan:
// 			if !more {
// 				loop = false
// 				break
// 			}
// 			if msg.Data != nil {
// 				fmt.Println(msg.Data.Count)
// 			}
// 			if msg.Errors != nil {
// 				fmt.Println("error:", msg.Errors)
// 				loop = false
// 			}
// 		case err = <-errChan:
// 			return
// 		case <-time.After(time.Second * 5):
// 			loop = false
// 		}
// 	}
// }

// //go:generate go run github.com/Khan/genqlient genqlient.yaml
