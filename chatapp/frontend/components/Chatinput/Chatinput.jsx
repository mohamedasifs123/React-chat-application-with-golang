import React,{Component} from "react";
import "./Chatinput.scss" ;

class Chatinput extends Component{
    renter(){
        return{
            <div className = "Chatinput">
            <input onKeyDown = {this.props.send} placeholder ="enter the message" >
            </div>
        }
    }
}

export default Chatinput;