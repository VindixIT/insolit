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

func CreateFaturaReparticaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create FaturaReparticaos")
		name := r.FormValue("Name")
		cnpj := r.FormValue("Cnpj")
		snippet1 := ""
		snippet2 := ""
		if cnpj != "" {
			snippet1 = ", cnpj"
			snippet2 = ", '" + cnpj + "'"
		}
		sqlStatement := "INSERT INTO faturas_reparticao ( name" + snippet1 + " ) VALUES ( $1" + snippet2 + " ) RETURNING id"
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
		http.Redirect(w, r, route.FaturasReparticaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateFaturaReparticaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update FaturaReparticao")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		cnpj := r.FormValue("CNPJ")
		sqlStatement := "UPDATE faturas_reparticao SET name=$1, cnpj=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, cnpj, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " |CNPJ: " + cnpj)
		http.Redirect(w, r, route.FaturasReparticaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteFaturaReparticaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete FaturaReparticao")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM faturas_reparticao WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.FaturasReparticaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListFaturasReparticaoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Faturas Reparticao")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id FROM faturas_reparticao order by id asc")
		sec.CheckInternalServerError(err, w)
		var faturas []mdl.FaturaReparticao
		var fatura mdl.FaturaReparticao
		var i = 1
		for rows.Next() {
			err = rows.Scan(&fatura.Id)
			sec.CheckInternalServerError(err, w)
			fatura.Order = i
			i++
			faturas = append(faturas, fatura)
		}
		var page mdl.PageFaturasReparticao
		page.AppName = mdl.AppName
		page.FaturasReparticao = faturas
		page.Title = "Faturas de Repartição"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/faturasreparticao/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-FaturasReparticao", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
