function editContratoGeracao(e) {
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var contratoId = e.parentNode.parentNode.childNodes[3].innerText;
    var contratoConcessionariaId = e.parentNode.parentNode.childNodes[3].childNodes[1].value;
    // var concessionariaName = e.parentNode.parentNode.childNodes[5].innerText;
    var concessionariaId = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
    // var clienteName = e.parentNode.parentNode.childNodes[7].innerText;
    var clienteId = e.parentNode.parentNode.childNodes[7].childNodes[1].value;
    var unidadeConsumidora = e.parentNode.parentNode.childNodes[9].innerText;
    var vencimentoEm = e.parentNode.parentNode.childNodes[11].innerText;
    var assinaturaEm = e.parentNode.parentNode.childNodes[13].innerText;
	document.getElementById('ContratoGeracaoIdForUpdate').value = contratoId;
    document.getElementById('ClienteForUpdate').value = clienteId;
    document.getElementById('ConcessionariaForUpdate').value = concessionariaId;
	document.getElementById('ContratoConcessionariaForUpdate').value = contratoConcessionariaId;
    document.getElementById('UCForUpdate').value = unidadeConsumidora;
	document.getElementById('VencimentoEmForUpdate').value = vencimentoEm;
    document.getElementById('AssinaturaEmForUpdate').value = assinaturaEm;
}

function deleteContratoGeracao(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var contratoId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ContratoGeracaoIdForDelete').value = contratoId;
}