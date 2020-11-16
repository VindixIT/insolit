function editUsina(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get usina id to update
    var usinaId = e.parentNode.parentNode.childNodes[3].innerText;
	alert("Aqui UsinaId", usinaId)
    var usinaName = e.parentNode.parentNode.childNodes[3].childNodes[1].value;
    //var parqueName = e.parentNode.parentNode.childNodes[5].innerText;
    var parqueId = e.parentNode.parentNode.childNodes[7].childNodes[1].value;
    var potencia = e.parentNode.parentNode.childNodes[9].innerText;
    var energiaMedia = e.parentNode.parentNode.childNodes[11].childNodes[1].value;
    var potenciaNominal = e.parentNode.parentNode.childNodes[13].childNodes[2].value;
    //var moduloName = e.parentNode.parentNode.childNodes[7].innerText;
    var moduloId = e.parentNode.parentNode.childNodes[9].childNodes[1].value;
    //var inversorName = e.parentNode.parentNode.childNodes[7].innerText;
    var inversorId = e.parentNode.parentNode.childNodes[9].childNodes[1].value;
	document.getElementById('UsinaIdForUpdate').value = usinaId;
    document.getElementById('ParqueUpdate').value = parqueId;
	document.getElementById('ModuloUpdate').value = moduloId;
    document.getElementById('InversorUpdate').value = inversorId;
    document.getElementById('NameUpdate').value = usinaName;
    document.getElementById('PotenciaUpdate').value = potencia;
    document.getElementById('PotenciaNominalUpdate').value = potenciaNominal;
    document.getElementById('EnergiaMediaUpdate').value = energiaMedia;
}

function deleteUsina(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var usinaId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('UsinaIdForDelete').value = usinaId;
}