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

func CreateContratoGeracaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Contratos Geracao")
		r.ParseForm()
		clienteId := r.Form["ClienteGForInsert"]
		usinaId := r.Form["UsinaGForInsert"]
		concessionariaId := r.Form["ConcessionariaGForInsert"]
		contratoConcessionaria := r.FormValue("ContratoConcessionaria")
		unidadeConsumidora := r.FormValue("UC")
		VencimentoEm := r.FormValue("VencimentoEm")
		AssinaturaEm := r.FormValue("AssinaturaEm")
		sqlStatement := "INSERT INTO contratos_geracao(" +
			" concessionaria_id, cliente_id, usina_id" +
			" contrato_concessionaria, unidade_consumidora," +
			" vencimento, assinatura_em) " +
			" VALUES ($1, $2, $3, $4, $5, $6, $7)"
		log.Println(sqlStatement)
		// id = 0
		Db.QueryRow(sqlStatement, concessionariaId[0], clienteId[0], 
			usinaId[0], contratoConcessionaria, unidadeConsumidora,
			VencimentoEm, AssinaturaEm)
		//if err != nil {
		//panic(err.Error())
		http.Redirect(w, r, route.ContratosGeracaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateContratoGeracaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		r.ParseForm()
		id := r.FormValue("Id")
		clienteId := r.Form["ClienteGForUpdate"]
		usinaId := r.Form["UsinaGForUpdate"]
		concessionariaId := r.Form["ConcessionariaGForUpdate"]
		contratoConcessionaria := r.FormValue("ContratoConcessionariaGForUpdate")
		unidadeConsumidora := r.FormValue("UnidadeConsumidoraGForUpdate")
		VencimentoEm := r.FormValue("VencimentoEmForUpdate")
		AssinaturaEm := r.FormValue("AssinaturaEmForUpdate")
		sqlStatement := "UPDATE contratos_geracao" +
			" SET concessionaria_id = $1, cliente_id= $2, usina_id=$3" +
			" contrato_concessionaria=$4, unidade_consumidora=$5," +
			" vencimento=$6, assinatura_em=$7 WHERE id=$8"
		updtForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		updtForm.Exec(clienteId[0], usinaId[0], concessionariaId[0],  contratoConcessionaria,
			 unidadeConsumidora, VencimentoEm, AssinaturaEm, id)
		log.Println("UPDATE: Id: " + id)
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
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ContratosGeracaoRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListContratosGeracaoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Contratos Geracao")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT " +
			"a.id," +
			"a.cliente_id, " +
			"a.usina_id, " +
			"a.concessionaria_id," +
			"coalesce(c.name,'') as concessionaria_nome, " +
			"coalesce(d.name,'') as usina_nome, " +
			"coalesce(b.name,'') as cliente_nome, " +
			"coalesce(a.contrato_concessionaria,''), " +
			"coalesce(a.unidade_consumidora,''), " +
			"a.vencimento, " +
			"coalesce(to_char(a.assinatura_em,'DD/MM/YYYY'),'') " +
			"FROM contratos_geracao a " +
			"LEFT JOIN clientes b ON b.id = a.cliente_id " +
			"LEFT JOIN usinas d ON d.id = a.usina_id " +
			"LEFT JOIN concessionarias c ON c.id = a.concessionaria_id " +
			"ORDER BY a.assinatura_em DESC" 
		log.Println("sql: " + query)
		rows, _ := Db.Query(query)
		var contratosGeracao []mdl.ContratoGeracao
		var contratoGeracao mdl.ContratoGeracao
		var i = 1
		for rows.Next() {
			rows.Scan(
				&contratoGeracao.Id,
				&contratoGeracao.ClienteId,
				&contratoGeracao.UsinaId,
				&contratoGeracao.ConcessionariaId,
				&contratoGeracao.ConcessionariaName,
				&contratoGeracao.UsinaName,
				&contratoGeracao.ClienteName,
				&contratoGeracao.ContratoConcessionaria,
				&contratoGeracao.UnidadeConsumidora,
				&contratoGeracao.VencimentoEm,
				&contratoGeracao.AssinaturaEm)
			contratoGeracao.Order = i
			i++
			log.Println(contratoGeracao)
			contratosGeracao = append(contratosGeracao, contratoGeracao)
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
		query = "SELECT id, name FROM public.usinas"
		// log.Println("Query cliente: " + query)
		rows, _ = Db.Query(query)
		var usinas []mdl.Usina
		var usina mdl.Usina
		i = 1
		for rows.Next() {
			rows.Scan(&usina.Id,
				&usina.Name)
			usina.Order = i
			i++
			usinas = append(usinas, usina)
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
		var page mdl.PageContratosGeracao
		page.AppName = mdl.AppName
		page.Clientes = clientes
		page.Usinas = usinas
		log.Println(len(usinas))
		page.Concessionarias = concessionarias
		page.Title = "Contratos de Geracao"
		page.ContratosGeracao = contratosGeracao
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/contratosgeracao/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-ContratosGeracao", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
