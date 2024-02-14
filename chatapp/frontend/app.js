import React,{Component} from "react";
import Chathistory  from "./components/Chathistory/Chathistory" ;
import Header  from "./components/Header/Header" ;
import Message  from "./components/Message/Measage" ;
import "./app.css";
import {connect,sendMsg} from "./api";

class App from Component{
    constructor(props){
    super(props);
    this.state={
        Chathistory:[]
    };
    }

    componentDidMount(){
        connect((msg)=>{
            console.log("new messaage")
            this.setState(prevState=>({
                Chathistory:[...prevState.Chathistory,msg]
            }))
            console.log(this.state);
        });

        
    }
    render(){
        return{
             <div className="App" >
            <Header/>
            <Chathistory chathistory={this.state.Chathistory}/>
            <Chatinput send={this.send}
            </div>
        }
    }
} 