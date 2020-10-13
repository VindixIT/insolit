function editInversor(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get modulo id to update
    var moduloId = e.parentNode.parentNode.childNodes[3].innerText;
    var modelo = e.parentNode.parentNode.childNodes[5].innerText;
    var fabricante = e.parentNode.parentNode.childNodes[7].innerText;
    var potenciaNominal = e.parentNode.parentNode.childNodes[9].innerText;
	document.getElementById('InversorIdForUpdate').value = moduloId;
    document.getElementById('ModeloForUpdate').value = modelo;
	document.getElementById('FabricanteForUpdate').value = fabricante;
    document.getElementById('PotenciaNominalForUpdate').value = potenciaNominal;
}

function deleteInversor(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var moduloId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('InversorIdForDelete').value = moduloId;
}