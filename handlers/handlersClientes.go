package handlers

import (
	"html/template"
	mdl "insolit/models"
	route "insolit/routes"
	sec "insolit/security"
	"log"
	"net/http"
	//"strconv"
)

func CreateClienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Cliente")
		name := r.FormValue("Name")
		endereco := r.FormValue("Endereco")
		capacidade := r.FormValue("Capacidade")
		cnpj := r.FormValue("Cnpj")
		sqlStatement := "INSERT INTO clientes( name, endereco, capacidade, cnpj) VALUES ($1, $2, $3,$4) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, endereco, capacidade, cnpj).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		//log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Code: " + code)
		http.Redirect(w, r, route.ClientesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateClienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		//log.Println("Update Cliente")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		endereco := r.FormValue("Endereco")
		capacidade := r.FormValue("Capacidade")
		cnpj := r.FormValue("Cnpj")
		sqlStatement := "UPDATE clientes SET name=$1, endereco=$2, capacidade=$3, cnpj=$4 WHERE id=$5"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, endereco, capacidade, cnpj, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " | cnpj: " + cnpj)
		http.Redirect(w, r, route.ClientesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteClienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete Cliente")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM clientes WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ClientesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}


func ListClientesHandler(w http.ResponseWriter, r *http.Request) {
	if sec.IsAuthenticated(w, r) {
		log.Println("List Clientes")
		rows, err := Db.Query("SELECT id, name, endereco, capacidade, cnpj FROM clientes order by id asc")
		sec.CheckInternalServerError(err, w)
		var clientes []mdl.Cliente
		var cliente mdl.Cliente
		var i = 1
		for rows.Next() {
			err = rows.Scan(&cliente.Id, &cliente.Name, &cliente.Endereco, &cliente.Capacidade, &cliente.Cnpj)
			sec.CheckInternalServerError(err, w)
			cliente.Order = i
			i++
			clientes = append(clientes, cliente)
		}
		var page mdl.PageClientes
		page.AppName = mdl.AppName
		page.Clientes = clientes
		page.Title = "Clientes"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/clientes/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Clientes", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
