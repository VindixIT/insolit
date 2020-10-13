function editCredito(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get parque id to update
    var creditoId = e.parentNode.parentNode.childNodes[3].innerText;
    var contratoId = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
    var energiaGerada = e.parentNode.parentNode.childNodes[7].innerText;
    var iniciaEm = e.parentNode.parentNode.childNodes[9].innerText;
    var terminaEm = e.parentNode.parentNode.childNodes[11].innerText;
	document.getElementById('CreditoIdForUpdate').value = creditoId;
    document.getElementById('ContratoForUpdate').value = contratoId;
	document.getElementById('EnergiaGeradaForUpdate').value = energiaGerada;
    document.getElementById('IniciaEmForUpdate').value = iniciaEm;
    document.getElementById('TerminaEmForUpdate').value = terminaEm;
}

function deleteCredito(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var parqueId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('CreditoIdForDelete').value = parqueId;
}