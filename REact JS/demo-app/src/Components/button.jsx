// Component for rendering the "Send" button
function SendMessage({ handleSend }) {
  
  return (
    <div className="send-button">
      <button type="button" onClick={handleSend}>
        Send
      </button>
    </div>
  );
}

export default SendMessage;