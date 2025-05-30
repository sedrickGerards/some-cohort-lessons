function Modal(props) {

    return <div className="modal">
        <p>Are you sure?</p>
        
            <button className="btn btn--cancel" onClick={props.onCancel}>Cancel</button>
        <button className="btn" onClick={props.onCancel}>Delete</button>
        
    </div>
}


export default Modal;