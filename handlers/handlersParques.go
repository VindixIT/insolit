package handlers

import (
	"html/template"
	mdl "insolit/models"
	route "insolit/routes"
	sec "insolit/security"
	"log"
	"net/http"
	"strconv"
)

func CreateParqueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Parques")
		name := r.FormValue("Name")
		endereco := r.FormValue("Endereco")
		cidade := r.FormValue("Cidade")
		estado := r.FormValue("Estado")
		sqlStatement := "INSERT INTO parques(name, endereco, cidade, estado) VALUES ($1,$2,$3,$4) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, endereco, cidade, estado).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		//		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Qtd: " + qtd + " | Price: " + price)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name)
		http.Redirect(w, r, route.ParquesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateParqueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Parque")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		endereco := r.FormValue("Endereco")
		cidade := r.FormValue("Cidade")
		estado := r.FormValue("Estado")
		sqlStatement := "UPDATE parques SET name=$1, endereco=$2, cidade=$3, estado=$4 WHERE id=$5"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, endereco, cidade, estado, id)
		//log.Println("UPDATE: Id: " + id + " | Name: " + name + " | Endereco: " + endereco + )"
		//" | Cidade: " + cidade+ " | Estado: " + estado)
		http.Redirect(w, r, route.ProdutosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteParqueHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Parque")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM parques WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.ParquesRoute, 301)
}

func ListParquesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Parques")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, name, endereco, cidade, estado FROM parques order by id asc")
		sec.CheckInternalServerError(err, w)
		var parques []mdl.Parque
		var parque mdl.Parque
		var i = 1
		for rows.Next() {
			err = rows.Scan(&parque.Id, &parque.Name, &parque.Endereco, &parque.Cidade, &parque.Estado)
			sec.CheckInternalServerError(err, w)
			parque.Order = i
			i++
			parques = append(parques, parque)
		}
		var page mdl.PageParques
		page.AppName = mdl.AppName
		page.Parques = parques
		page.Title = "Parques"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/parques/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Parques", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
