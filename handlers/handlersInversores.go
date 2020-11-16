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
		modelo:= r.FormValue("Modelo")
		fabricante := r.FormValue("Fabricante")
		potenciaNominal := r.FormValue("PotenciaNominal")
		sqlStatement := "INSERT INTO inversores (modelo,fabricante, potencia_nominal ) VALUES ( $1,$2,$3) RETURNING id"
		log.Println(sqlStatement)
		id := 0
		err := Db.QueryRow(sqlStatement, modelo, fabricante, potenciaNominal).Scan(&id)
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

func UpdateInversorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Inversor")
		modelo:= r.FormValue("Modelo")
		fabricante := r.FormValue("Fabricante")
		potenciaNominal := r.FormValue("PotenciaNominal")
		sqlStatement := "UPDATE inversores SET name=$1, cnpj=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(modelo, fabricante, potenciaNominal)
		log.Println("UPDATE: Id Inversor: " + modelo)
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
		rows, err := Db.Query("SELECT id, modelo, fabricante, potencia_nominal FROM inversores order by id asc")
		sec.CheckInternalServerError(err, w)
		var inversores []mdl.Inversor
		var inversor mdl.Inversor
		var i = 1
		for rows.Next() {
			err = rows.Scan(&inversor.Id,
					&inversor.Modelo,
					&inversor.Fabricante,
					&inversor.PotenciaNominal)
			sec.CheckInternalServerError(err, w)
			inversor.Order = i
			i++
			inversores = append(inversores,inversor)
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
