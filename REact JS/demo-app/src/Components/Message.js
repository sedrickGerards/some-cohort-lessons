import { useState } from "react";
import "./message.css";
import {getOfflineBotReply} from "./Response.js";

// Component for rendering the input field
function InputMessage({ inputText, handleText, handleSend }) {
  return (
    <div className="input-text">
      <input
        type="text"
        placeholder="Enter your text..."
        onChange={handleText}
        value={inputText}
        onKeyDown={(e)=> {
          if (e.key === "Enter") {
            handleSend();
          }
        }
        }
      />
    </div>
  );
}

// Component for rendering the "Send" button
function SendMessage({ handleSend }) {
 
  return (
    <div className="send-button">
      <button type="button" 
      onClick={handleSend}>
        Send
      </button>
    </div>
  );
}

// function EnterKeySendMessage (e){
//   if(e.key === "Enter"){
//     handleSend()
//   }
// }
// Main chat component managing messages and input state
function Message() {
  // State for tracking all messages in the conversation
  const [messages, setMessages] = useState([
    {
      message: "hey there Gentle-Sed Welcome!",
      sender: "bot",
      id: crypto.randomUUID(),
    },
  ]);

  // State for tracking the user input
  const [inputText, setInputText] = useState("");

  // Called when user types in the input box
  const handleText = (event) => {
    setInputText(event.target.value);
  };
  
  
  
  // Called when the "Send" button is clicked
  const handleSend = () => {

    // calling the bot response function
    const botResponse = getOfflineBotReply(inputText);
    // Add the current inputText as a new message
    setMessages([
      ...messages,
      {
        message: inputText,
        sender: "user",
        id: crypto.randomUUID(),
      },
      {
        message: botResponse,
        sender: "bot",
        id: crypto.randomUUID(),
      },
    ]);

    // Clear the input field after sending
    setInputText("");

    console.log("message sent");
  };

  // mapping through message to get the inputed message

  const newMessage = messages.map((message) => {
    return (
      <div className="message-and-logo-container" key={message.id}>
        {message.sender === "user" ? (
          <div className="user-chat" >
            <p>{message.message}</p> <img src="/img/user.png" alt="logo" />
          </div>
        ) : (
          <div className="bot-chat" >
            <img src="/img/robot.png" alt="logo" /> <p>{message.message}</p>
          </div>
        )}
      </div>
    );
  });

  return (
    <div className="main-chat-container">
      <div className="new-message-container">{newMessage}</div>

      <div className="user-input">
        <InputMessage inputText={inputText} handleText={handleText} handleSend={handleSend}/>
        <SendMessage handleSend={handleSend} />
      </div>
    </div>
  );
}

export default Message;
