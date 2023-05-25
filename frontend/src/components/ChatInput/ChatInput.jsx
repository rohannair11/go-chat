import React, {Component} from 'react'
import "./ChatInput.scss"

class ChatInput extends Component{
  render(){
    return(
      <div className='chatInput'>
        <input onKeyDown={this.props.send} placeholder='type a message' />
      </div>
    )
  }
}

export default ChatInput