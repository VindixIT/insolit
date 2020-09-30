package handlers

import (
	"html/template"
	mdl "insolit/models"
	route "insolit/routes"
	sec "insolit/security"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CreateWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Workflow")
		name := r.FormValue("Name")
		stereotype := r.FormValue("Stereotype")
		sqlStatement := "INSERT INTO workflows(name) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, stereotype).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Entitity: " + name)
		for key, value := range r.Form {
			if strings.HasPrefix(key, "activity") {
				array := strings.Split(value[0], "#")
				beerId := strings.Split(array[1], ":")[1]
				log.Println("beerId: " + beerId)
				//				sqlStatement := "INSERT INTO items(order_id, beer_id, quantity, price, item_value) VALUES ($1,$2,$3,$4,$5) RETURNING id"
				//				err := Db.QueryRow(sqlStatement, orderId, beerId, qtd, price, itemValue).Scan(&itemId)
				sec.CheckInternalServerError(err, w)
				if err != nil {
					panic(err.Error())
				}
				sec.CheckInternalServerError(err, w)
				//				log.Println(l)
			}
		}
		http.Redirect(w, r, route.WorkflowsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Workflow")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		sqlStatement := "UPDATE workflow SET name=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name)
		http.Redirect(w, r, route.WorkflowsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete Workflow")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM workflow WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.WorkflowsRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListWorkflowsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Workflows")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT id, name FROM workflows order by id asc"
		log.Println("List WF -> Query: " + query)
		rows, err := Db.Query(query)
		sec.CheckInternalServerError(err, w)
		var workflows []mdl.Workflow
		var workflow mdl.Workflow
		var i = 1
		for rows.Next() {
			err = rows.Scan(&workflow.Id, &workflow.Name)
			sec.CheckInternalServerError(err, w)
			workflow.Order = i
			i++
			workflows = append(workflows, workflow)
		}
		query = "SELECT a.id, a.name, a.origin_status_id, b.name as origin_status, " +
			"a.destination_status_id, c.name as destination_status, a.other_than " +
			"FROM actions a, status b, status c " +
			"WHERE a.origin_status_id = b.id " +
			"AND a.destination_status_id = c.id " +
			"order by a.id asc"
		log.Println("List WF -> Query: " + query)
		rows, err = Db.Query(query)
		sec.CheckInternalServerError(err, w)
		var actions []mdl.Action
		var action mdl.Action
		i = 1
		for rows.Next() {
			err = rows.Scan(&action.Id, &action.Name, &action.OriginId, &action.Origin, &action.DestinationId, &action.Destination, &action.OtherThan)
			sec.CheckInternalServerError(err, w)
			action.Order = i
			i++
			actions = append(actions, action)
		}
		query = "SELECT id, name FROM roles order by name asc"
		log.Println("List WF -> Query: " + query)
		rows, err = Db.Query(query)
		sec.CheckInternalServerError(err, w)
		var roles []mdl.Role
		var role mdl.Role
		i = 1
		for rows.Next() {
			err = rows.Scan(&role.Id, &role.Name)
			sec.CheckInternalServerError(err, w)
			role.Order = i
			i++
			roles = append(roles, role)
		}
		var page mdl.PageWorkflows
		page.AppName = mdl.AppName
		page.Actions = actions
		page.Roles = roles
		page.Workflows = workflows
		page.Title = "Workflows"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/workflows/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Workflows", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}

}
