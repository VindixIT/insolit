function editParque(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get parque id to update
    var parqueId = e.parentNode.parentNode.childNodes[3].innerText;
    var parqueName = e.parentNode.parentNode.childNodes[5].innerText;
	document.getElementById('ParqueIdForUpdate').value = parqueId;
    document.getElementById('ParqueNameForUpdate').value = parqueName;
}

function deleteParque(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var parqueId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ParqueIdForDelete').value = parqueId;
}