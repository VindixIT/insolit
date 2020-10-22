package handlers

import (
	"html/template"
	mdl "insolit/models"
	route "insolit/routes"
	sec "insolit/security"
	"log"
	"net/http"
	//	"strconv"
)

func CreateContratoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Contrato")
	ContratoConcessionaria := r.FormValue("ContratoConcessionaria")
	log.Println(ContratoConcessionaria)
	UC := r.FormValue("UC")
	log.Println(UC)
	EnderecoUC := r.FormValue("EnderecoUC")
	log.Println(EnderecoUC)
	VencimentoEm := r.FormValue("VencimentoEm")
	log.Println(VencimentoEm)
	AssinaturaEm := r.FormValue("AssinaturaEm")
	log.Println(AssinaturaEm)
	clienteId := r.Form["clientes"]
	log.Println(clienteId[0])
	concessionariaId := r.Form["concessionarias"]
	log.Println(concessionariaId[0])
	sqlStatement := "INSERT INTO contratos_consumo(" +
		" concessionaria_id, cliente_id, " +
		" contrato_concessionaria, unidade_consumidora, endereco_uc, " +
		" vencimento, assinatura_em) " +
		" VALUES ($1, $2, $3, $4, $5, $6, $7)"
	log.Println(sqlStatement)
	Db.QueryRow(sqlStatement, concessionariaId[0], clienteId[0],
		ContratoConcessionaria, UC, EnderecoUC,
		VencimentoEm, AssinaturaEm)
	http.Redirect(w, r, route.ContratosRoute, 301)
}

func ListContratosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List ContratosConsumo")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT " +
			"a.id," +
			"a.concessionaria_id," +
			"c.name as concessionaria_nome, " +
			"a.cliente_id, " +
			"b.name as cliente_nome," +
			"a.contrato_concessionaria," +
			"a.unidade_consumidora, " +
			"a.endereco_uc, " +
			"a.vencimento, " +
			"a.assinatura_em " +
			"FROM contratos_consumo a " +
			"LEFT JOIN clientes b ON b.id = a.cliente_id " +
			"LEFT JOIN concessionarias c ON c.id = a.concessionaria_id"
		log.Println("Query: " + query)
		rows, _ := Db.Query(query)
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
		query = "SELECT id, name FROM clientes"
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
		page.Title = "NOVO Contratos de Consumo"
		page.ContratosConsumo = contratos
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/contratos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Contratos", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
