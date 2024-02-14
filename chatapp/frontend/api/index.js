var socket = new WebSocket('ws://localhost:9000/ws');

let connect =(cb) => {
    console.log("connecting")
    
    
    socket.onopen = ()=>{
        console.log("successfully connnected")
    }
    
    socket.onmessage =(msg) =>{
        console.log("message received",msg)
    }

    socket.onclose =(event) =>{
        console.log("connection closed",event)

    }
    socket.onerror =(error) =>{
        console.log("error occured",error)
    }
    

};

let sendMsg =(msg)=>{
    console.log("sending msg",msg)
    socket.send(msg)
};

export {connect,sendMsg} ;