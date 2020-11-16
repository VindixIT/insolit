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

func CreateUsinaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Contratos Consumo")
		//id := r.FormValue("Id")
		r.ParseForm()
		parqueId := r.Form["ParqueInsert"]
		moduloId := r.Form["ModuloInsert"]
		inversorId := r.Form["InversorInsert"]
		log.Println("ParqueInsert" + parqueId[0])
		nome := r.FormValue("Name")
		potencia := r.FormValue("Potencia")
		potenciaNominal := r.FormValue("PotenciaNominal")
		energiaMedia := r.FormValue("EnergiaMedia")
		sqlStatement := "INSERT INTO usinas(" +
			" name,parque_id, inversor_id, modulo_id, " +
			" potencia, potencia_nominal, energia_media) " +
			" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
		log.Println(sqlStatement)
		// id = 0
		Db.QueryRow(sqlStatement, parqueId[0], moduloId[0], inversorId[0],
			nome, potencia, potenciaNominal, energiaMedia)
		//if err != nil {
		//panic(err.Error())
		http.Redirect(w, r, route.UsinasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateUsinaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Usina")
		id := r.FormValue("Id")
		r.ParseForm()
		parqueId := r.Form["ParqueForInsert"]
		moduloId := r.Form["ModuloForInsert"]
		inversorId := r.Form["InversorForInsert"]
		//log.Println("ParqueForInsert" + parqueId[0])
		nome := r.FormValue("Name")
		potencia := r.FormValue("Potencia")
		potenciaNominal := r.FormValue("PotenciaNominal")
		energiaMedia := r.FormValue("EnergiaMedia")
		sqlStatement := "UPDATE usinas " +
		" SET name=$1, parque_id=$2, inversor_id=$3, " +
		" modulo_id=$4, potencia=$5, " +
		" potencia_nominal=$6, energia_media=$7 WHERE id=$8"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(id, nome, parqueId, inversorId,
			           moduloId , potencia, potenciaNominal,
			           energiaMedia, id )
		log.Println("UPDATE: Id: " + id) // + " | Name: " + name + " | Endereco: " + endereco + )"
		//" | Cidade: " + cidade+ " | Estado: " + estado)
		http.Redirect(w, r, route.UsinasRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteUsinaHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Delete Usina")
	if r.Method == "POST" {
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM usinas WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
	}
	http.Redirect(w, r, route.UsinasRoute, 301)
}

func ListUsinasHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Usinas")
	if sec.IsAuthenticated(w, r) {
		query := "SELECT " +
		" a.id, " +
		" coalesce(a.name,'') as UsinaNome," +
		" coalesce(b.name,'') as ParqueNome, " +
		" coalesce(c.modelo,'') as InversorNome, " +
		" coalesce(d.modelo,'') as moduloNome," +
		" a.potencia, d.potencia_pico " +
		" c.potencia_nominal, a.energia_media " +
		" FROM usinas a " +
		" LEFT JOIN parques b ON a.id = b.id " +
		" LEFT JOIN inversores c ON a.id = c.id" +
		" LEFT JOIN modulos d ON a.id = d.id" +
		" WHERE a.id = $1 " +
		" order by id asc"
		log.Println(query)
		var usinas []mdl.Usina
		var usina mdl.Usina
		rows, _ := Db.Query(query)
		var i = 1
		for rows.Next() {
			_ = rows.Scan(
				&usina.Id, 
				&usina.Name,
				&usina.ParqueNome,
				&usina.InversorNome,
				&usina.ModuloNome,
				&usina.Potencia,
				&usina.PotenciaPico,
				&usina.PotenciaNominal,
				&usina.EnergiaMedia)
			usina.Order = i
			i++
			usinas = append(usinas, usina)
		}
		query = "SELECT id, modelo, fabricante, potencia_pico FROM modulos order by id asc"
		log.Println("List Modulo -> Query: " + query)
		rows, _ = Db.Query(query)
		var modulos []mdl.Modulo
		var modulo mdl.Modulo
		 i = 1
		for rows.Next() {
			_ = rows.Scan(&modulo.Id, &modulo.Modelo)
			modulo.Order = i
			i++
			modulos = append(modulos, modulo)
		}
		query = "SELECT id, modelo, fabricante, potencia_nominal FROM inversores order by id asc"
		log.Println("List Modulo -> Query: " + query)
		rows, _ = Db.Query(query)
		var inversores []mdl.Inversor
		var inversor mdl.Inversor
		i = 1
		for rows.Next() {
			_ = rows.Scan(&inversor.Id, &inversor.Modelo)
			inversor.Order = i
			i++
			inversores = append(inversores, inversor)
		}
		var page mdl.PageUsinas
		page.AppName = mdl.AppName
		page.Usinas = usinas
		page.Modulos = modulos
		page.Inversores = inversores
		page.Title = "Usinas"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/usinas/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Usinas", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
