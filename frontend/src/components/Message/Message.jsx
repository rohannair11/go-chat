// import React, {Component} from "react";
// import "./Message.scss"

// class Message extends Component{
//     constructor(props){
//         super(props);
//         let temp = JSON.parse(this.props.message);
//         this.state = {
//             message: temp
//         }
//     }
//     render(){
//         return(
//             <div className="message">
//                 {this.state.message.body}
//             </div>
//         );
//     };

// }
// export default Message;

import React, { useState, useEffect } from "react";
import "./Message.scss";

const Message = (props) => {
  const [message, setMessage] = useState(null);

  useEffect(() => {
    const temp = JSON.parse(props.message);
    setMessage(temp);
  }, [props.message]);

  return <div className="message">{message && message.body}</div>;
};

export default Message;