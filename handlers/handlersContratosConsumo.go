package handlers

import (
	"html/template"
	mdl "insolit/models"
	route "insolit/routes"
	sec "insolit/security"
	"log"
	"net/http"
	//"strconv"
)

func CreateContratoConsumoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create ContratosConsumo")
		concessionariaId := r.Form["ConcessionariaForInsert"]
		log.Println(concessionariaId)
		log.Println(concessionariaId[0])
		clienteId := r.Form["ClienteForInsert"]
		log.Println("clienteId" + clienteId[0])
		contratoConcessionaria := r.FormValue("contratoConcessionaria")
		unidadeConsumidora := r.FormValue("unidadeConsumidora")
		VencimentoEm := r.FormValue("vencimentoEm")
		AssinaturaEm := r.FormValue("assinaturaEm")
		sqlStatement := "INSERT INTO contratos_consumo(" +
						" concessionaria_id, cliente_id, "+ 
						" contrato_concessionaria, unidade_consumidora, " +
						" vencimento, assinatura_em) " +
						" VALUES ($1, $2, $3, $4)"
		log.Println(sqlStatement)
		id := 0
		err := Db.QueryRow(sqlStatement,concessionariaId[0], clienteId[0], 
			contratoConcessionaria, unidadeConsumidora,
			VencimentoEm, AssinaturaEm).Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		http.Redirect(w, r, route.ContratosConsumoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateContratoConsumoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update ContratoConsumo")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		cnpj := r.FormValue("Cnpj")
		log.Println("name: " + name)
		sqlStatement := "UPDATE contratos_consumo SET name=$1, " +
		 "cnpj=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, cnpj, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " |CNPJ: " + cnpj)
		http.Redirect(w, r, route.ContratosConsumoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteContratoConsumoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete ContratoConsumo")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM contratos_consumo WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ContratosConsumoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListContratosConsumoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List ContratosConsumo")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT "+
				"a.id,"+ 
				"a.concessionaria_id," + 
				"c.name as concessionaria_nome, "+
				"a.cliente_id, "+
				"b.name as cliente_nome," + 
				"a.contrato_concessionaria," + 
				"a.unidade_consumidora, "+
				"a.endereco_uc, "+
				"a.vencimento, "+
				"a.assinatura_em "+
				"FROM contratos_consumo a "+ 
				"LEFT JOIN clientes b ON b.id = a.cliente_id "+
				"LEFT JOIN concessionarias c ON c.id = a.concessionaria_id"
		log.Println("Query: " + query)
		rows, _:= Db.Query(query)
		var contratos []mdl.ContratoConsumo
		var contrato mdl.ContratoConsumo
		var i = 1
		for rows.Next() {
			rows.Scan(&contrato.Id,
				&contrato.ContratoConcessionaria)
			contrato.Order = i
			i++
			contratos = append(contratos, contrato)
		}
		query = "SELECT id, name FROM public.clientes"
		log.Println("Query cliente: " + query)
		rows, _ = Db.Query(query)
		var clientes []mdl.Cliente
		var cliente mdl.Cliente
		i = 1
		for rows.Next() {
			rows.Scan(&cliente.Id,
				 &cliente.Name)
			cliente.Order = i
			i++
			clientes = append(clientes, cliente)
		}
		query = "SELECT id, name FROM public.concessionarias"
		log.Println("Query concessionaria: " + query)
		rows, _ = Db.Query(query)
		var concessionarias []mdl.Concessionaria
		var concessionaria mdl.Concessionaria
		i = 1
		for rows.Next() {
			rows.Scan(&concessionaria.Id,
				 &concessionaria.Name)
			concessionaria.Order = i
			i++
			concessionarias = append(concessionarias, concessionaria)
		}
		var page mdl.PageContratosConsumo
		page.AppName = mdl.AppName
		page.Clientes = clientes
		page.Concessionarias = concessionarias
		page.Title = "Contratos de Consumo"
		page.ContratosConsumo = contratos
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/contratosconsumo/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-ContratosConsumo", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
