package main

import{
	"fmt"
	"net/http"
}
func serverWS(pool *webSocket.Pool,w http.ResponseWriter,r http.Request){
	fmt.println("websocket server reached")
    
	if err!=nil{
		fmt.fprintf(w,"%V\n",err)
	}
	conn,err:= webSocket.Upgrade(w,r)
	client:= &webSocket.Client{
		Conn:conn,
		Pool:pool,
	}

	pool.Register <- client
	client.Read()

}

func setupRoutes(){
	pool:=webSocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws",func(w http.ResponseWriter,r http.Request){
		serverWS(pool,w,r)
	})
}
fuc main(){
	fmt.println("ASIF web project")
	setupRoutes()
	http.ListenAndServe(":9000",nil)

}

