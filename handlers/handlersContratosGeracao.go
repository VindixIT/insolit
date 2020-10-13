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

func CreateContratoGeracaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create ContratosGeracao")
		name := r.FormValue("Name")
		cnpj := r.FormValue("Cnpj")
		snippet1 := ""
		snippet2 := ""
		if cnpj != "" {
			snippet1 = ", cnpj"
			snippet2 = ", '" + cnpj + "'"
		}
		sqlStatement := "INSERT INTO contratos_geracao ( name" + snippet1 + " ) VALUES ( $1" + snippet2 + " ) RETURNING id"
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
		http.Redirect(w, r, route.ContratosGeracaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateContratoGeracaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update ContratoGeracao")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		cnpj := r.FormValue("CNPJ")
		sqlStatement := "UPDATE contratos_geracao SET name=$1, cnpj=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, cnpj, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " |CNPJ: " + cnpj)
		http.Redirect(w, r, route.ContratosGeracaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteContratoGeracaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete ContratoGeracao")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM contratos_geracao WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ContratosGeracaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListContratosGeracaoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List ContratosGeracao")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, contrato_concessionaria FROM contratos_geracao order by id asc")
		sec.CheckInternalServerError(err, w)
		var contratos []mdl.ContratoGeracao
		var contrato mdl.ContratoGeracao
		var i = 1
		for rows.Next() {
			err = rows.Scan(&contrato.Id, &contrato.ContratoConcessionaria)
			sec.CheckInternalServerError(err, w)
			contrato.Order = i
			i++
			contratos = append(contratos, contrato)
		}
		var page mdl.PageContratosGeracao
		page.AppName = mdl.AppName
		page.ContratosGeracao = contratos
		page.Title = "Contratos de Geração"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/contratosgeracao/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-ContratosGeracao", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
