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

func CreateInversorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Inversores")
		name := r.FormValue("Name")
		cnpj := r.FormValue("Cnpj")
		snippet1 := ""
		snippet2 := ""
		if cnpj != "" {
			snippet1 = ", cnpj"
			snippet2 = ", '" + cnpj + "'"
		}
		sqlStatement := "INSERT INTO inversores ( name" + snippet1 + " ) VALUES ( $1" + snippet2 + " ) RETURNING id"
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
		http.Redirect(w, r, route.InversoresRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateInversorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Inversor")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		cnpj := r.FormValue("CNPJ")
		sqlStatement := "UPDATE inversores SET name=$1, cnpj=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, cnpj, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " |CNPJ: " + cnpj)
		http.Redirect(w, r, route.InversoresRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteInversorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete Inversor")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM inversores WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.InversoresRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListInversoresHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Inversores")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id FROM inversores order by id asc")
		sec.CheckInternalServerError(err, w)
		var inversores []mdl.Inversor
		var concessionaria mdl.Inversor
		var i = 1
		for rows.Next() {
			err = rows.Scan(&concessionaria.Id)
			sec.CheckInternalServerError(err, w)
			concessionaria.Order = i
			i++
			inversores = append(inversores, concessionaria)
		}
		var page mdl.PageInversores
		page.AppName = mdl.AppName
		page.Inversores = inversores
		page.Title = "Inversores"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/inversores/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Inversores", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
