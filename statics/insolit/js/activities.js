var activity_tobe_deleted;

class Activity {
	constructor(id, actionId, actionName, startAt, endAt, expTime, expActionId, expActionName, roles){
		this.id = id;
		this.actionId = actionId;
		this.actionName = actionName;
		this.startAt = startAt ;
		this.endAt = endAt;
		this.expTime = expTime;
		this.expActionId = expActionId;
		this.expActionName = expActionName;
		this.roles = roles;
	}
}

class Role {
	constructor(id, name){
		this.id = id;
		this.name = name;
	}
}

function limparCamposActivityForm(form){
	var a = document.getElementById('action-'+form);
	a.options[a.selectedIndex].selected = false;
	document.getElementById('start-at-'+form).value = "";
	document.getElementById('end-at-'+form).value = "";
	document.getElementById('exp-time-'+form).value = "";
	document.getElementById('exp-action-'+form).value = "";
	document.getElementById('roles-'+form).value = "";
}

function criarActivity(){
	var a = document.getElementById('action-create');
	var actionId = a.options[a.selectedIndex].value;
	var erros = '';
	if(actionId==''){
		if(actionId==''){
			erros += 'Falta preencher a ação.\n';
		}
		alert(erros);
		return;
	}
	var actionName = a.options[a.selectedIndex].text;
	var startAt = document.getElementById('start-at-create').value;
	var endAt = document.getElementById('end-at-create').value;
	var expTime = document.getElementById('exp-time-create').value;
	var ea = document.getElementById('exp-action-create');
	var expActionId = ea.options[ea.selectedIndex].value;
	var expActionName = ea.options[ea.selectedIndex].text;
	var roles = getSelectedRoles(document.getElementById('roles-create'));
	activity = new Activity(activities.length, actionId, actionName, startAt, endAt, expTime, expActionId, expActionName, roles);
	activities.push(activity);
	addActRow("table-activities-"+contexto);
	limparCamposActivityForm('create');
	document.getElementById('create-activity-form').style.display='none';
}


function updateActivity(e) {
	//limparCamposActivityForm('edit');
	var editActivityForm = document.getElementById('edit-activity-form');
	editActivityForm.style.display = 'block';
	var id = e.parentNode.parentNode.childNodes[0].childNodes[1].value;
	var actionId = e.parentNode.parentNode.childNodes[0].childNodes[2].value;
	var order = e.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var startAt = e.parentNode.parentNode.childNodes[1].innerText;
	var endAt = e.parentNode.parentNode.childNodes[2].innerText;
	var expTime = e.parentNode.parentNode.childNodes[3].innerText;
	var expActionId = e.parentNode.parentNode.childNodes[4].childNodes[0].value;
	var rolesIds = e.parentNode.parentNode.childNodes[5].childNodes[0].value;
	// Atribuindo os valores de edit-item-form
	document.getElementById('id-edit').value=id;
	document.getElementById('action-edit').value = actionId;
	document.getElementById('start-at-edit').value=startAt;
	document.getElementById('end-at-edit').value=endAt;
	document.getElementById('exp-time-edit').value=expTime;
	document.getElementById('exp-action-edit').value=expActionId;
	setSelectedRoles(document.getElementById('roles-edit'), rolesIds);
	document.getElementById('order-edit').value=order;
}

function editarActivity(){
	var id = document.getElementById('id-edit').value;
	var a = document.getElementById('action-edit');
	var actionId = a.options[a.selectedIndex].value;
	var erros = '';
	if(actionId==''){
		if(actionId==''){
			erros += 'Falta preencher a ação.\n';
		}
		alert(erros);
		return;
	}
	var actionName = a.options[a.selectedIndex].text;
	var startAt = document.getElementById('start-at-edit').value;
	var endAt = document.getElementById('end-at-edit').value;
	var expTime = document.getElementById('exp-time-edit').value;
	var ea = document.getElementById('exp-action-edit');
	var expActionId = ea.options[ea.selectedIndex].value;
	var expActionName = ea.options[ea.selectedIndex].text;
	var roles = getSelectedRoles(document.getElementById('roles-edit'));
	var order = document.getElementById('order-edit').value;
	activity = new Activity(id, actionId, actionName, startAt, endAt, expTime, expActionId, expActionName, roles);
	activities[order]=activity;
	updateActRow("table-activities-"+contexto,order);
	//limparCamposActivityForm('edit');
	var editActivityForm = document.getElementById('edit-activity-form');
	editActivityForm.style.display = 'none';
}

function showDeleteActivityForm(e){
	var deleteActivityForm = document.getElementById('delete-activity-form');
	deleteActivityForm.style.display = 'block';
	activity_tobe_deleted = e;
}


function deleteactivity() {
	var order = activity_tobe_deleted.parentNode.parentNode.childNodes[0].childNodes[0].value;
	var newActivities = [];
	for(i=0;i<activities.length;i++){
		if(i != order){
			newActivities.push(activities[i]);
		}
	}
	activities = newActivities;
	activity_tobe_deleted.parentNode.parentNode.innerHTML = '';
	var deleteActivityForm = document.getElementById('delete-activity-form');
	deleteActivityForm.style.display = 'none';
}

function addActRow(tableID) {
	let tableRef = document.getElementById(tableID);
	let newRow = tableRef.insertRow(-1);
	order = activities.length-1;
	activity = activities[order];
	// actvt
	let newCell = newRow.insertCell(0);
	let newText = document.createTextNode(activity.actionName);
	var jsonActvt = JSON.stringify(activity);
	jsonActvt = jsonActvt.split(',').join('#');
	jsonActvt = jsonActvt.split('"').join('');
	jsonActvt = jsonActvt.split('{').join('');
	jsonActvt = jsonActvt.split('}').join('');
	newCell.appendChild(newText);
	newCell.innerHTML =	'<input type="hidden" name="activity'+activity.actionId+'" value="'+jsonActvt+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="actionId" value="'+activity.actionId+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="id" value="'+activity.id+'"/>'+newCell.innerHTML;
	newCell.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+newCell.innerHTML;
	// Inicia Em
	newCell = newRow.insertCell(1);
	newText = document.createTextNode(activity.startAt);
	newCell.appendChild(newText);
	// Termina Em
	newCell = newRow.insertCell(2);
	newText = document.createTextNode(activity.endAt);
	newCell.appendChild(newText);
	// Expiration Time
	newCell = newRow.insertCell(3);
	newText = document.createTextNode(activity.expTime);
	newCell.appendChild(newText);
	// Action Time Expired
	newCell = newRow.insertCell(4);
	newText = document.createTextNode(activity.expActionName);
	newCell.appendChild(newText);
	newCell.innerHTML = '<input type="hidden" name="expActionId" value="'+activity.expActionId+'"/>'+newCell.innerHTML;
	// Roles Allowed
	newCell = newRow.insertCell(5);
	newText = document.createTextNode('');
	newCell.appendChild(newText);
	let res = getRolesCommaSep(activity.roles);
	newCell.innerHTML = res[0];
	newCell.innerHTML = '<input type="hidden" name="roles" value="'+res[1]+'"/>'+newCell.innerHTML;
	newCell = newRow.insertCell(6);
	// Botão Editar
	var btnEditar = document.createElement('input');
	btnEditar.type = "button";
	btnEditar.className = "w3-btn w3-teal";
	btnEditar.value = "Editar";
	btnEditar.onclick = function() {updateActivity(btnEditar)};
	newCell.appendChild(btnEditar);
	// Botão Apagar
	var btnApagar = document.createElement('input');
	btnApagar.type = "button";
	btnApagar.className = "w3-btn w3-red";
	btnApagar.value = "Apagar";
	btnApagar.onclick = function() {showDeleteActivityForm(btnApagar)};
	newCell.appendChild(btnApagar);
}

function updateActRow(tableID, order){
	let tableRef = document.getElementById(tableID);
	let rowNumber = 3+parseInt(order);
	let row = tableRef.childNodes[1].childNodes[rowNumber];
	let celula = row.childNodes[0];
	celula.innerText = activities[order].actionName;
	var jsonActivity = JSON.stringify(activities[order]);
	jsonActivity = jsonActivity.split(',').join('#');
	jsonActivity = jsonActivity.split('"').join('');
	jsonActivity = jsonActivity.split('{').join('');
	jsonActivity = jsonActivity.split('}').join('');
	celula.innerHTML =	'<input type="hidden" name="activity'+activity.actionId+'" value="'+jsonActivity+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="actionId" value="'+activity.actionId+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="id" value="'+activity.id+'"/>'+celula.innerHTML;
	celula.innerHTML = '<input type="hidden" name="order" value="'+order+'"/>'+celula.innerHTML;
	row.childNodes[1].innerText = activities[order].startAt;
	row.childNodes[2].innerText = activities[order].endAt;
	row.childNodes[3].innerText = activities[order].expTime;
	row.childNodes[4].innerText = activities[order].expActionName;
	row.childNodes[4].innerHTML = '<input type="hidden" name="expActionId" value="'+activities[order].expActionId+'"/>'+row.childNodes[4].innerHTML;
	let res = getRolesCommaSep(activities[order].roles);
	row.childNodes[5].innerHTML = res[0];
	row.childNodes[5].innerHTML = '<input type="hidden" name="roles" value="'+res[1]+'"/>'+row.childNodes[5].innerHTML;
}

function getSelectedRoles(select) {
  var result = [];
  var options = select && select.options;
  var opt;
  for (var i=0, iLen=options.length; i<iLen; i++) {
    opt = options[i];
    if (opt.selected) {
	  var role = new Role(opt.value, opt.text);
      result.push(role);
    }
  }
  return result;
}


function setSelectedRoles(select, ids) {
  var sels = ids.split(',');
  for (var n=0; n<sels.length; n++) {
  	var options = select.options;
  	var opt;
	for (var i=0; i<options.length; i++) {
	  opt = options[i];
	  let sel = sels[n].trim();
	  if (opt.value == sel) {
		opt.selected = 'selected';
	  }
	}
  }
}

function getRolesCommaSep(roles){
	let txts = '';
	let vls = '';
	let results = [];
	let n = 0;
	for(;n<roles.length;n++){
		txts += roles[n].name + ", ";
		vls += roles[n].id + ", ";
	}
	if(n>0){
		txts = txts.substring(0, txts.length-2);
		vls = vls.substring(0, vls.length-2);
	}
	results.push(txts);
	results.push(vls);
	return results;
}

