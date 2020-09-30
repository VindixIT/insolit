function editCliente(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get produto id to update
    var clienteId = e.parentNode.parentNode.childNodes[3].innerText;
	alert("cliente Id:"+ clienteId);
    var clienteName = e.parentNode.parentNode.childNodes[5].innerText;
    var clienteEndereco = e.parentNode.parentNode.childNodes[7].innerText;
    var clienteCapacidade = e.parentNode.parentNode.childNodes[9].innerText;
    var clienteCnpj = e.parentNode.parentNode.childNodes[11].innerText;
	
	document.getElementById('ClienteIdForUpdate').value = clienteId;
    document.getElementById('ClienteNameForUpdate').value = clienteName;
    document.getElementById('ClienteEnderecoForUpdate').value = clienteName;
    document.getElementById('ClienteCapacidadeForUpdate').value = clienteName;
    document.getElementById('ClienteCnpjForUpdate').value = clienteName;
}

function deleteCliente(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var clienteId = e.parentNode.parentNode.childNodes[3].innerText;
	alert("cliente ID:"+ clienteId);
    document.getElementById('clienteIdForDelete').value = clienteId;
}