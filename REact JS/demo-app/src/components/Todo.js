  import { useState } from "react";

import Modal from "./Modal";
import Backdrop from "./Backdrop";


  
  function Todo(props) {

    const [modalIsOpen, setModalIsOpen] = useState(false)

    // const [modalIsOpen, setModalIsOpen] = useState(false);
    
    // creating func that hadles delete '
    function handleDelete() { 
      setModalIsOpen(true)
    }
function handleCloseModal(){
  setModalIsOpen(false)
}


    return (
      <div className="main-modal-container">
        <div className="todo-box">
          <h2>{props.text}</h2>
          <div className="delete-button">
            <button className="delete-action" onClick={handleDelete}>
              Delete
            </button>
          </div>
          {modalIsOpen && <Modal onCancel={handleCloseModal}/>}
          {modalIsOpen && <Backdrop onCancel={handleCloseModal} />}
        </div>
      </div>
    );

}

export default Todo;