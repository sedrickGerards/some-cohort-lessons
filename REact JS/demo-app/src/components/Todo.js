

function Todo(props) {
    return (
        <div className="todo-box">
        <h2>{props.text}</h2>
        <div className="delete-button">
          <button>Delete </button>
        </div>
      </div>

    )

}

export default Todo;