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
		log.Println("Create Contratos Consumo")
		r.ParseForm()
		clienteId := r.Form["Cliente"]
		concessionariaId := r.Form["Concessionaria"]
		log.Println("clienteId" + clienteId[0])
		contratoConcessionaria := r.FormValue("ContratoConcessionaria")
		unidadeConsumidora := r.FormValue("UC")
		EnderecoUC := r.FormValue("EnderecoUC")
		VencimentoEm := r.FormValue("VencimentoEm")
		AssinaturaEm := r.FormValue("AssinaturaEm")
		sqlStatement := "INSERT INTO contratos_consumo(" +
			" concessionaria_id, cliente_id, " +
			" contrato_concessionaria, unidade_consumidora, endereco_uc, " +
			" vencimento, assinatura_em) " +
			" VALUES ($1, $2, $3, $4, $5, $6, $7)"
		log.Println(sqlStatement)
		// id = 0
		Db.QueryRow(sqlStatement, concessionariaId[0], clienteId[0],
			contratoConcessionaria, unidadeConsumidora, EnderecoUC,
			VencimentoEm, AssinaturaEm)
		//if err != nil {
		//panic(err.Error())
		http.Redirect(w, r, route.ContratosConsumoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateContratoConsumoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		clienteId := r.Form["ClienteForUpdate"]
		concessionariaId := r.Form["ConcessionariaForUpdate"]
		contratoConcessionaria := r.FormValue("ContratoConcessionariaForUpdate")
		unidadeConsumidora := r.FormValue("UnidadeConsumidoraForUpdate")
		EnderecoUC := r.FormValue("EnderecoUCForUpdate")
		VencimentoEm := r.FormValue("VencimentoEmForUpdate")
		AssinaturaEm := r.FormValue("AssinaturaEmForUpdate")
		sqlStatement := "UPDATE contratos_consumo(" +
			" concessionaria_id, cliente_id, " +
			" contrato_concessionaria, unidade_consumidora, endereco_uc, " +
			" vencimento, assinatura_em) " +
			" VALUES ($1, $2, $3, $4, $5, $6, $7)"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		//sec.CheckInternalServerError(err, w)
		updtForm.Exec(concessionariaId[0], clienteId[0], contratoConcessionaria, unidadeConsumidora, EnderecoUC,
			VencimentoEm, AssinaturaEm)
		log.Println("UPDATE: clienteId: " + clienteId[0])
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
		deleteForm, _ := Db.Prepare(sqlStatement)
		deleteForm.Exec(id)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ContratosConsumoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListContratosConsumoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List ContratosConsumo")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT " +
			"a.id," +
			"a.concessionaria_id," +
			"coalesce(c.name,'') as concessionaria_nome, " +
			"a.cliente_id, " +
			"coalesce(b.name,'') as cliente_nome, " +
			"coalesce(a.contrato_concessionaria,''), " +
			"coalesce(a.unidade_consumidora,''), " +
			"a.endereco_uc, " +
			"a.vencimento, " +
			"coalesce(to_char(a.assinatura_em,'DD/MM/YYYY'),'') " +
			"FROM contratos_consumo a " +
			"LEFT JOIN clientes b ON b.id = a.cliente_id " +
			"LEFT JOIN concessionarias c ON c.id = a.concessionaria_id"
		log.Println("sql: " + query)
		rows, _ := Db.Query(query)
		var contratosConsumo []mdl.ContratoConsumo
		var contratoConsumo mdl.ContratoConsumo
		var i = 1
		for rows.Next() {
			rows.Scan(
				&contratoConsumo.Id,
				&contratoConsumo.ConcessionariaId,
				&contratoConsumo.ConcessionariaName,
				&contratoConsumo.ClienteId,
				&contratoConsumo.ClienteName,
				&contratoConsumo.ContratoConcessionaria,
				&contratoConsumo.UnidadeConsumidora,
				&contratoConsumo.EnderecoUC,
				&contratoConsumo.VencimentoEm,
				&contratoConsumo.AssinaturaEm)
			contratoConsumo.Order = i
			i++
			log.Println(contratoConsumo)
			contratosConsumo = append(contratosConsumo, contratoConsumo)
		}
		query = "SELECT id, name FROM public.clientes"
		// log.Println("Query cliente: " + query)
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
		// log.Println("Query concessionaria: " + query)
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
		log.Println(len(clientes))
		page.Concessionarias = concessionarias
		page.Title = "Contratos de Consumo"
		page.ContratosConsumo = contratosConsumo
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/contratosconsumo/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-ContratosConsumo", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
