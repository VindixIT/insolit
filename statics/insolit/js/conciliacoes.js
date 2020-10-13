function editConciliacao(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get parque id to update
    var creditoId = e.parentNode.parentNode.childNodes[3].innerText;
    var credito = e.parentNode.parentNode.childNodes[5].childNodes[1].value;
    var fatura = e.parentNode.parentNode.childNodes[7].innerText;
	document.getElementById('ConciliacaoIdForUpdate').value = creditoId;
	document.getElementById('CreditoForUpdate').value = credito;
    document.getElementById('FaturaConsumoForUpdate').value = fatura;
}

function deleteConciliacao(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var parqueId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ConciliacaoIdForDelete').value = parqueId;
}