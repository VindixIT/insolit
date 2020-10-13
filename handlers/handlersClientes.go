package handlers

import (
	"html/template"
	mdl "insolit/models"
	route "insolit/routes"
	"log"
	"net/http"
	//"strconv"
	sec "insolit/security"
)

func CreateClienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Cliente")
		name := r.FormValue("Name")
		endereco := r.FormValue("Endereco")
		cidade := r.FormValue("Cidade")
		estado := r.FormValue("Estado")
		cnpj := r.FormValue("Cnpj")
		sqlStatement := "INSERT INTO clientes( name, endereco, cidade, estado, cnpj) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, endereco, cidade, estado, cnpj).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		http.Redirect(w, r, route.ClientesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateClienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		endereco := r.FormValue("Endereco")
		cidade := r.FormValue("Cidade")
		estado := r.FormValue("Estado")
		cnpj := r.FormValue("Cnpj")
		sqlStatement := "UPDATE clientes SET name=$1, endereco=$2, cidade=$3, estado=$4, cnpj=$5 WHERE id=$6"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, endereco, cidade, estado, cnpj, id)
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
		rows, err := Db.Query("SELECT id, name, endereco, cidade, estado, cnpj FROM clientes order by id asc")
		sec.CheckInternalServerError(err, w)
		var clientes []mdl.Cliente
		var cliente mdl.Cliente
		var i = 1
		for rows.Next() {
			err = rows.Scan(&cliente.Id, &cliente.Name, &cliente.Endereco, &cliente.Cidade, &cliente.Estado, &cliente.Cnpj)
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
