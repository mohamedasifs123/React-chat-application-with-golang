import React,{Component} from "react";
import "./Chathistory.scss" ;
import Message from  "../Message/Message";

class Chathistory extends Component{
    render(){
        console.log(this.props.Chathistory);
        this.props.Chathistory.map(msg=><Message key ={msg.timestamp} message ={msg.data}/>);

        return{
            <div className="Chathistory">
            <H2>Chat History</H2>
            {message}
            </div>
        }
    }
}

export default Chathistory;
