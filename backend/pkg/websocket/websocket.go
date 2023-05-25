// contains all the code for upgrading a normal http connection to a websocket simul connection

package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//from the docs https://pkg.go.dev/github.com/gorilla/websocket

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}
