typr Pool struct {
	Register      chan *Client
	Unregistered  chan *Client
	Clients       map[*Client]bool
	Broadcast     chan Message
}

func NewPool() *Pool{
	return &Pool{
		Register:make(chan *Client),
		Unregister:make(chan *Client),
	    Clients:make(map[*CLient]bool),
		Broadcast:make(chan Message)
	}
	
}

func (pool *Pool) Start(){
	for{
		select{
		case client = <-pool.Register:
			pool.Clients[Client]=true
			fmt.println("size of connected clients:",len(pool.clients))
			for client,_ := range pool.Clients{
				fmt.println(client)
				client.conn.WriteJSON(Message(Type:!,Body:"New user joined"))
			}
			break
		case client = <-pool.Unregister:
			delete(pool.Clients,Client)
			fmt.println("size of connected clients:",len(pool.clients))
			for client,_ := range pool.Clients{
				
				client.conn.WriteJSON(Message(Type:1,Body:"User Disconnected"))
			}
			break
		case message:<-pool.Broadcast:
			fmt.println("Sending mesge to all clients in the room")
			for client,_ := range pool.Clients{
			if err	:= client.conn.WriteJSON(Message) err!=nil{
				fmt.println(err)
				return
			}
		}
	}
}