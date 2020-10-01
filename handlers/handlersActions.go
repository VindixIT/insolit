package handlers

import (
	"encoding/json"
	"html/template"
	mdl "insolit/models"
	route "insolit/routes"
	sec "insolit/security"
	"log"
	"net/http"
	"strconv"
)

func CreateActionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Action")
		name := r.FormValue("Name")
		except := r.FormValue("Except")
		otherThan := false
		if except != "" {
			otherThan = true
		}
		log.Println(except)
		originStatus := r.Form["OriginStatusForInsert"]
		log.Println(originStatus)
		destinationStatus := r.Form["DestinationStatusForInsert"]
		log.Println(destinationStatus)
		roles := r.Form["RolesForInsert"]
		sqlStatement := "INSERT INTO actions(name, origin_status_id, destination_status_id, other_than) VALUES ($1, $2, $3, $4) RETURNING id"
		actionId := 0
		err := Db.QueryRow(sqlStatement, name, originStatus[0], destinationStatus[0], otherThan).Scan(&actionId)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		for _, roleId := range roles {
			sqlStatement := "INSERT INTO actions_roles(action_id,role_id) VALUES ($1,$2)"
			Db.QueryRow(sqlStatement, actionId, roleId)
			sec.CheckInternalServerError(err, w)
			if err != nil {
				panic(err.Error())
			}
			sec.CheckInternalServerError(err, w)
		}
		sec.CheckInternalServerError(err, w)
		sqlStatement = "INSERT INTO actions_status(action_id,origin_status_id,destination_status_id) VALUES ($1,$2,$3)"
		Db.QueryRow(sqlStatement, actionId, originStatus[0], destinationStatus[0])
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(actionId) + " | Name: " + name)
		http.Redirect(w, r, route.ActionsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}

}

func UpdateActionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Action")
		actionId := r.FormValue("Id")
		name := r.FormValue("Name")
		sqlStatement := "UPDATE actions SET name=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, actionId)
		log.Println("UPDATE: Id: " + actionId + " | Name: " + name)

		var rolesDB = ListRolesByActionIdHandler(actionId)
		var rolesPage []mdl.Role
		var rolePage mdl.Role
		for _, roleId := range r.Form["RolesForUpdate"] {
			rolePage.Id, _ = strconv.ParseInt(roleId, 10, 64)
			rolesPage = append(rolesPage, rolePage)
		}
		if len(rolesPage) < len(rolesDB) {
			log.Println("Quantidade de Roles da Página: " + strconv.Itoa(len(rolesPage)))
			if len(rolesPage) == 0 {
				DeleteRolesByActionHandler(actionId) //DONE
			} else {
				var diffDB []mdl.Role = rolesDB
				for n := range rolesPage {
					if containsRole(diffDB, rolesPage[n]) {
						diffDB = removeRole(diffDB, rolesPage[n])
					}
				}
				DeleteRolesHandler(diffDB) //DONE
			}
		} else {
			var diffPage []mdl.Role = rolesPage
			for n := range rolesDB {
				if containsRole(diffPage, rolesDB[n]) {
					diffPage = removeRole(diffPage, rolesDB[n])
				}
			}
			var role mdl.Role
			for i := range diffPage {
				role = diffPage[i]
				log.Println("Action Id: " + actionId)
				sqlStatement := "INSERT INTO actions_roles(action_id, role_id) VALUES ($1,$2)"
				log.Println(sqlStatement)
				Db.QueryRow(sqlStatement, actionId, role.Id)
			}
		}
		http.Redirect(w, r, route.ActionsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func containsRole(roles []mdl.Role, roleCompared mdl.Role) bool {
	for n := range roles {
		if roles[n].Id == roleCompared.Id {
			return true
		}
	}
	return false
}

func removeRole(roles []mdl.Role, roleToBeRemoved mdl.Role) []mdl.Role {
	var newRoles []mdl.Role
	for i := range roles {
		if roles[i].Id != roleToBeRemoved.Id {
			newRoles = append(newRoles, roles[i])
		}
	}
	return newRoles
}

func DeleteActionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete Action")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM actions_status WHERE action_id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM actions_roles WHERE action_id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sqlStatement = "DELETE FROM actions WHERE id=$1"
		deleteForm, err = Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ActionsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListActionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Actions")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT a.id, a.name, a.origin_status_id, b.name as origin_name, a.destination_status_id, c.name as destination_name, a.other_than " +
			"FROM actions a, status b, status c where a.origin_status_id = b.id and a.destination_status_id = c.id order by id asc"
		log.Println("List Action -> Query: " + query)
		rows, err := Db.Query(query)
		sec.CheckInternalServerError(err, w)
		var actions []mdl.Action
		var action mdl.Action
		var i = 1
		for rows.Next() {
			err = rows.Scan(&action.Id, &action.Name, &action.OriginId, &action.Origin, &action.DestinationId, &action.Destination, &action.OtherThan)
			sec.CheckInternalServerError(err, w)
			action.Order = i
			i++
			actions = append(actions, action)
		}
		query = "SELECT id, name, stereotype FROM status order by name asc"
		log.Println("List Action -> Query: " + query)
		rows, err = Db.Query(query)
		sec.CheckInternalServerError(err, w)
		var statuss []mdl.Status
		var status mdl.Status
		i = 1
		for rows.Next() {
			err = rows.Scan(&status.Id, &status.Name, &status.Stereotype)
			sec.CheckInternalServerError(err, w)
			status.Order = i
			i++
			statuss = append(statuss, status)
		}
		var page mdl.PageActions
		page.AppName = mdl.AppName
		page.Statuss = statuss
		page.Actions = actions
		page.Title = "Ação"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/actions/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Actions", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func LoadRolesByActionId(w http.ResponseWriter, r *http.Request) {
	log.Println("Load Roles By Action Id")
	r.ParseForm()
	var actionId = r.FormValue("actionId")
	log.Println("actionId: " + actionId)
	roles := ListRolesByActionIdHandler(actionId)
	jsonRoles, _ := json.Marshal(roles)
	w.Write([]byte(jsonRoles))
	log.Println("JSON")
}
