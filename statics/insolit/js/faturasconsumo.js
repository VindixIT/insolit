function editFaturaConsumo(e) {
    var editForm = document.getElementById('edit-form');
    editForm.style.display = 'block';
    var contratoId = e.parentNode.parentNode.childNodes[3].innerText;
    // var clienteName = e.parentNode.parentNode.childNodes[5].innerText;
    var contratoConcessionaria = e.parentNode.parentNode.childNodes[7].innerText;
    var consumoMes = e.parentNode.parentNode.childNodes[9].innerText;
    var creditosMes = e.parentNode.parentNode.childNodes[11].innerText;
    var saldoAcumulado = e.parentNode.parentNode.childNodes[11].childNodes[1].value;
    var valor = e.parentNode.parentNode.childNodes[13].innerText;
    var vencimentoEm = e.parentNode.parentNode.childNodes[15].innerText;
    var emissaoEm = e.parentNode.parentNode.childNodes[15].childNodes[1].value;
    var leituraAnterior = e.parentNode.parentNode.childNodes[15].childNodes[2].value;
    var leituraAtual = e.parentNode.parentNode.childNodes[15].childNodes[3].value;
    var pago = e.parentNode.parentNode.childNodes[17].innerText;
	document.getElementById('FaturaConsumoIdForUpdate').value = contratoId;
    document.getElementById('ContratoForUpdate').value = contratoConcessionaria;
    document.getElementById('ContratoConcessionariaForUpdate').value = contratoConcessionaria;
	document.getElementById('ConsumoMesForUpdate').value = consumoMes;
    document.getElementById('CreditosMesForUpdate').value = creditosMes;
    document.getElementById('SaldoAcumuladoForUpdate').value = saldoAcumulado;
	document.getElementById('ValorForUpdate').value = valor;
    document.getElementById('VencimentoEmForUpdate').value = vencimentoEm;
    document.getElementById('EmissaoEmForUpdate').value = emissaoEm;
    document.getElementById('LeituraAnteriorForUpdate').value = leituraAnterior;
    document.getElementById('LeituraAtualForUpdate').value = leituraAtual;
    document.getElementById('PagoForUpdate').value = pago;
}

function deleteFaturaConsumo(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var faturaId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('FaturaConsumoIdForDelete').value = faturaId;
}