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

func CreateConcessionariaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Concessionarias")
		name := r.FormValue("Name")
		cnpj := r.FormValue("Cnpj")
		snippet1 := ""
		snippet2 := ""
		if cnpj != "" {
			snippet1 = ", cnpj"
			snippet2 = ", '" + cnpj + "'"
		}
		sqlStatement := "INSERT INTO concessionarias ( name" + snippet1 + " ) VALUES ( $1" + snippet2 + " ) RETURNING id"
		log.Println(sqlStatement)
		id := 0
		err := Db.QueryRow(sqlStatement, name).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		//		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Qtd: " + qtd + " | Price: " + price)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name)
		http.Redirect(w, r, route.ConcessionariasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateConcessionariaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Concessionaria")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		cnpj := r.FormValue("CNPJ")
		sqlStatement := "UPDATE concessionarias SET name=$1, cnpj=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, cnpj, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " |CNPJ: " + cnpj)
		http.Redirect(w, r, route.ConcessionariasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteConcessionariaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete Concessionaria")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM concessionarias WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ConcessionariasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListConcessionariasHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Concessionarias")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, name, coalesce(cnpj,'') as cnpj FROM concessionarias order by id asc")
		sec.CheckInternalServerError(err, w)
		var concessionarias []mdl.Concessionaria
		var concessionaria mdl.Concessionaria
		var i = 1
		for rows.Next() {
			err = rows.Scan(&concessionaria.Id, &concessionaria.Name, &concessionaria.Cnpj)
			sec.CheckInternalServerError(err, w)
			concessionaria.Order = i
			i++
			concessionarias = append(concessionarias, concessionaria)
		}
		var page mdl.PageConcessionarias
		page.AppName = mdl.AppName
		page.Concessionarias = concessionarias
		page.Title = "Concessionarias"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/concessionarias/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Concessionarias", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
