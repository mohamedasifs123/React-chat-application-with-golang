package websocket

import{
	"log"
	"net/http"
	"github.com/gorilla/websocket"
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

func Upgrade(w http.ResponseWriter, http.Request)(*websocket.Conn,error){
	Upgrader.CheckOrigin = func (r *http.Request) bool {return true}
		conn,err:=upgrader.Upgrade(w,r,nil)
		if err!=nil{
			log.println(err)
			return nil,err
		}
	return conn,nil
	}
