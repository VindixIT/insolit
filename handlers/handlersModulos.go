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

func CreateModuloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Modulos")
		modelo:= r.FormValue("Modelo")
		fabricante := r.FormValue("Fabricante")
		potenciaPico := r.FormValue("PotenciaPico")
		sqlStatement := "INSERT INTO modulos (modelo,fabricante, potencia_pico ) VALUES ( $1,$2,$3) RETURNING id"
		log.Println(sqlStatement)
		id := 0
		err := Db.QueryRow(sqlStatement, modelo, fabricante, potenciaPico).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		//		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Qtd: " + qtd + " | Price: " + price)
		log.Println("INSERT: Id: " + strconv.Itoa(id))
		http.Redirect(w, r, route.ModulosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateModuloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Modulo")
		modelo:= r.FormValue("Modelo")
		fabricante := r.FormValue("Fabricante")
		potenciaPico := r.FormValue("PotenciaPico")
		sqlStatement := "UPDATE modulos SET modelo=$1, fabricante=$2 , potencia_pico=$3 WHERE id=$4"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(modelo, fabricante, potenciaPico)
		log.Println("UPDATE: modelo: " + modelo )
		http.Redirect(w, r, route.ModulosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteModuloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete Modulo")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM modulos WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ModulosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListModulosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Modulos")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, modelo, fabricante, potencia_pico FROM modulos order by id asc")
		sec.CheckInternalServerError(err, w)
		var modulos []mdl.Modulo
		var modulo mdl.Modulo
		var i = 1
		for rows.Next() {
			err = rows.Scan(&modulo.Id,
					&modulo.Modelo,
					&modulo.Fabricante,
					&modulo.PotenciaPico)
			sec.CheckInternalServerError(err, w)
			modulo.Order = i
			i++
			modulos = append(modulos, modulo)
		}
		var page mdl.PageModulos
		page.AppName = mdl.AppName
		page.Modulos = modulos
		page.Title = "Módulos"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/modulos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Modulos", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
