function editProduto(e) {
    var editForm = document.getElementById('edit-form');
    // display update form
    editForm.style.display = 'block';
    // get produto id to update
    var produtoId = e.parentNode.parentNode.childNodes[3].innerText;
	//alert("produto Id:"+ produtoId);
    var produtoName = e.parentNode.parentNode.childNodes[5].innerText;
	//alert("produto Name:"+ produtoName);
	document.getElementById('ProdutoIdForUpdate').value = produtoId;
    document.getElementById('ProdutoNameForUpdate').value = produtoName;
}

function deleteProduto(e) {
    var deleteForm = document.getElementById('delete-form');
    deleteForm.style.display = 'block';
    var produtoId = e.parentNode.parentNode.childNodes[3].innerText;
	//alert("produto ID:"+ produtoId);
    document.getElementById('ProdutoIdForDelete').value = produtoId;
}