function editConcessionaria(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    var concessionariaId = e.parentNode.parentNode.childNodes[3].innerText;
	//alert("conc. Id:"+ concessionariaId);
	var concessionariaName = e.parentNode.parentNode.childNodes[5].innerText;
	var concessionariaCNPJ = e.parentNode.parentNode.childNodes[7].innerText;
	//alert("concessionaria CNPJ:"+ concessionariaCNPJ);
	document.getElementById('ConcessionariaIdForUpdate').value = concessionariaId;
    document.getElementById('ConcessionariaNameForUpdate').value = concessionariaName;
    document.getElementById('ConcessionariaCNPJForUpdate').value = concessionariaCNPJ;
}

function deleteConcessionaria(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var concessionariaId = e.parentNode.parentNode.childNodes[3].innerText;
    document.getElementById('ConcessionariaIdForDelete').value = concessionariaId;
}