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

func CreateContratoReparticaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create ContratosReparticao")
		name := r.FormValue("Name")
		cnpj := r.FormValue("Cnpj")
		snippet1 := ""
		snippet2 := ""
		if cnpj != "" {
			snippet1 = ", cnpj"
			snippet2 = ", '" + cnpj + "'"
		}
		sqlStatement := "INSERT INTO contratos_reparticao ( name" + snippet1 + " ) VALUES ( $1" + snippet2 + " ) RETURNING id"
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
		http.Redirect(w, r, route.ContratosReparticaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateContratoReparticaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update ContratoReparticao")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		cnpj := r.FormValue("CNPJ")
		sqlStatement := "UPDATE contratos_reparticao SET name=$1, cnpj=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, cnpj, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " |CNPJ: " + cnpj)
		http.Redirect(w, r, route.ContratosReparticaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteContratoReparticaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete ContratoReparticao")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM contratos_reparticao WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ContratosReparticaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListContratosReparticaoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List ContratosReparticao")
	if sec.IsAuthenticated(w, r) {
		rows, _ := Db.Query("SELECT id, contrato_investidor FROM contratos_reparticao order by id asc")
		var contratos []mdl.ContratoReparticao
		var contrato mdl.ContratoReparticao
		var i = 1
		for rows.Next() {
			rows.Scan(&contrato.Id, &contrato.ContratoInvestidor)
			contrato.Order = i
			i++
			contratos = append(contratos, contrato)
		}
		rows, _ = Db.Query("SELECT id, name FROM produtos order by id asc")
		var produtos []mdl.Produto
		var produto mdl.Produto
		i = 1
		for rows.Next() {
			rows.Scan(&produto.Id, &produto.Name)
			produto.Order = i
			i++
			produtos = append(produtos, produto)
		}
		rows, _ = Db.Query("SELECT id, name FROM clientes order by id asc")
		var clientes []mdl.Cliente
		var cliente mdl.Cliente
		i = 1
		for rows.Next() {
			rows.Scan(&cliente.Id, &cliente.Name)
			cliente.Order = i
			i++
			clientes = append(clientes, cliente)
		}
		rows, _ = Db.Query("SELECT id, name FROM usinas order by id asc")
		var usinas []mdl.Usina
		var usina mdl.Usina
		i = 1
		for rows.Next() {
			rows.Scan(&usina.Id, &usina.Name)
			usina.Order = i
			i++
			usinas = append(usinas, usina)
		}
		var page mdl.PageContratosReparticao
		page.AppName = mdl.AppName
		page.ContratosReparticao = contratos
		page.Clientes = clientes
		page.Produtos = produtos
		page.Usinas = usinas
		page.Title = "Contratos de Repartição"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/contratosreparticao/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-ContratosReparticao", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
