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

func CreateUsinaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Usinas")
		name := r.FormValue("Name")
		endereco := r.FormValue("Endereco")
		cidade := r.FormValue("Cidade")
		estado := r.FormValue("Estado")
		sqlStatement := "INSERT INTO usinas(name, endereco, cidade, estado) VALUES ($1,$2,$3,$4) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, endereco, cidade, estado).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		//		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Qtd: " + qtd + " | Price: " + price)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name)
		http.Redirect(w, r, route.UsinasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateUsinaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Usina")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		endereco := r.FormValue("Endereco")
		cidade := r.FormValue("Cidade")
		estado := r.FormValue("Estado")
		sqlStatement := "UPDATE usinas SET name=$1, endereco=$2, cidade=$3, estado=$4 WHERE id=$5"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, endereco, cidade, estado, id)
		//log.Println("UPDATE: Id: " + id + " | Name: " + name + " | Endereco: " + endereco + )"
		//" | Cidade: " + cidade+ " | Estado: " + estado)
		http.Redirect(w, r, route.UsinasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteUsinaHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Usina")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM usinas WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.UsinasRoute, 301)
}

func ListUsinasHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Usinas")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id FROM usinas order by id asc")
		sec.CheckInternalServerError(err, w)
		var usinas []mdl.Usina
		var usina mdl.Usina
		var i = 1
		for rows.Next() {
			err = rows.Scan(&usina.Id)
			sec.CheckInternalServerError(err, w)
			usina.Order = i
			i++
			usinas = append(usinas, usina)
		}
		var page mdl.PageUsinas
		page.AppName = mdl.AppName
		page.Usinas = usinas
		page.Title = "Usinas"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/usinas/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Usinas", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
