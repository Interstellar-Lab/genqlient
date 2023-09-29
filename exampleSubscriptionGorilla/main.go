package main

// import (
// 	"context"
// 	"crypto/tls"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/Interstellar-Lab/genqlient/graphql"
// 	"github.com/gorilla/websocket"
// )

// type MyDialer struct {
// 	*websocket.Dialer
// }

// func (md *MyDialer) DialContext(ctx context.Context, urlStr string, requestHeader http.Header) (graphql.WSConn, error) {
// 	conn, _, err := md.Dialer.DialContext(ctx, urlStr, requestHeader)
// 	return graphql.WSConn(conn), err
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

// 	dialer := websocket.DefaultDialer
// 	dialer.TLSClientConfig = &tls.Config{}
// 	dialer.TLSClientConfig.InsecureSkipVerify = true
// 	headers := http.Header{}
// 	headers.Add("Authorization", "bearer "+key)
// 	headers.Add("auth", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiI3MmJiNjViMS1mYWI4LTRkNDAtODA3Yi0xZmRkNzExNTRiNmMiLCJwcm9kdWN0SUQiOiIiLCJleHAiOjE2OTQ2MTQwMjR9.rDsIlJk4GQRfl5KbvENR2uCqyHgDhVz_7MUVTDlDsyA")

// 	graphqlClient := graphql.NewClientUsingWebSocket(
// 		"wss://localhost:8080/query",
// 		&MyDialer{Dialer: dialer},
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

// //go:generate go run github.com/Interstellar-Lab/genqlient genqlient.yaml
