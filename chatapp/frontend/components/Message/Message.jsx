import React,{Component} from "react";
import "./Message.scss" ;

class Message extends Component{
    constructor(props){
        super(props)
        let tmp = JSON.parse(this.props.message)
        this.state = {
            message:tmp
        }
    }
     render(){
        return{
            <div>
             {this.state.message.body}
            </div>
        };
     };
}

export default Message;