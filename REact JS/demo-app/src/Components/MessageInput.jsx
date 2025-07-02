
// Component for rendering the input field
function InputMessage({ inputText, handleText, handleSend }) {
  return (
    <div className="input-text">
      <input
        type="text"
        placeholder="Enter your text..."
        onChange={handleText}
        value={inputText}
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            handleSend();
          }
        }}
      />
    </div>
  );
}

export default InputMessage;