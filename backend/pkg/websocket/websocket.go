//wesocket related code goes here

package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//create a two way comm buffer

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// upgrade regular http connection to a websocket one.
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}
