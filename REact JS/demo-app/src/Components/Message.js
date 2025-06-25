import { useState } from "react";

// Component for rendering the input field
function InputMessage({ inputText, handleText }) {
  return (
    <>
      <input
        type="text"
        placeholder="Enter your text..."
        onChange={handleText} 
        value={inputText} 
      />
    </>
  );
}

// Component for rendering the "Send" button
function SendMessage({ handleSend }) {
  return (
    <>
      <button type="button" onClick={handleSend}>
        Send
      </button>
    </>
  );
}

// Main chat component managing messages and input state
function Message() {
  // State for tracking all messages in the conversation
  const [messages, setMessages] = useState([
    {
      message: "sed",
      sender: "user",
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
    // Add the current inputText as a new message
    setMessages([
      ...messages,
      {
        message: inputText,
        sender: "user",
        id: crypto.randomUUID(),
      },
    ]);

    // Clear the input field after sending
    setInputText("");

 
    console.log("message sent");
  };

  return (
    <>
     
      <InputMessage inputText={inputText} handleText={handleText} />
      <SendMessage handleSend={handleSend} />

     
      <div>{inputText}</div>
    </>
  );
}

export default Message;
