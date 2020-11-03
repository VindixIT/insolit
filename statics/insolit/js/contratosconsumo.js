function editContratoConsumo(e) {
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var contratoId = e.parentNode.parentNode.childNodes[3].innerText;
    var contratoConcessionariaId = e.parentNode.parentNode.childNodes[3].childNodes[1].value;
    // var concessionariaName = e.parentNode.parentNode.childNodes[5].innerText;
    var clienteId = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
    var concessionariaId = e.parentNode.parentNode.childNodes[7].childNodes[1].value;
    // var clienteName = e.parentNode.parentNode.childNodes[7].innerText;
    var unidadeConsumidora = e.parentNode.parentNode.childNodes[9].innerText;
	var enderecoUnidadeConsumidora = e.parentNode.parentNode.childNodes[9].childNodes[1].value;
    var assinaturaEm = e.parentNode.parentNode.childNodes[11].innerText;
    var vencimentoEm = e.parentNode.parentNode.childNodes[13].innerText;
	document.getElementById('ContratoConsumoIdForUpdate').value = contratoId;
    document.getElementById('ClienteIdForUpdate').value = clienteId;
    document.getElementById('ConcessionariaIdForUpdate').value = concessionariaId;
	document.getElementById('ContratoConcessionariaForUpdate').value = contratoConcessionariaId;
    document.getElementById('UCForUpdate').value = unidadeConsumidora;
    document.getElementById('EnderecoUCForUpdate').value = enderecoUnidadeConsumidora;
    document.getElementById('AssinaturaEmForUpdate').value = formatarData(assinaturaEm);
	document.getElementById('VencimentoEmForUpdate').value = vencimentoEm;
}

function deleteContratoConsumo(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var contratoId = e.parentNode.parentNode.childNodes[3].innerText;
	alert(contratoId);
    document.getElementById('ContratoIdForDelete').value = contratoId;
}