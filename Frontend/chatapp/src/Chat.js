import React, { Component } from 'react';
import './Chat.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '', // Username of the sender
      message: '',  // Message input
      messages: []  // Array to store messages
    };
    this.socket = new WebSocket('ws://localhost:8000/ws');
  }

  componentDidMount() {
    const username = prompt("Please enter your username:");
    this.setState({ username });

    this.socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      this.setState((prevState) => ({
        messages: [...prevState.messages, message]
      }));
    };
  }

  handleSubmit = (event) => {
    event.preventDefault();
    const { username, message } = this.state;
    this.socket.send(JSON.stringify({ username, message }));
    this.setState({ message: '' });
  }




  // ... previous React code ...

render() {
    const { username, message, messages } = this.state;
    return (
      <div className="App">
        <h1>Chat Application</h1>
        <div className="chat-container">
          <div className="chat-messages">
            {messages.map((msg, index) => (
              <div
                key={index}
                className={`message ${msg.username === username ? 'sender' : 'receiver'}`}
              >
                {msg.username !== username && (
                  <span className={`name ${msg.username === username ? 'sender-name' : 'receiver-name'}`}>
                    {msg.username}
                  </span>
                )}
                {msg.message}
              </div>
            ))}
          </div>
          <form onSubmit={this.handleSubmit}>
            <input
              type="text"
              placeholder="Type a message..."
              value={message}
              onChange={(e) => this.setState({ message: e.target.value })}
            />
            <button type="submit">Send</button>
          </form>
        </div>
      </div>
    );
  }
}

export default App;
