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

func CreateProdutoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Produtos")
		name := r.FormValue("Name")
		sqlStatement := "INSERT INTO produtos(name) VALUES ($1) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		//		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Qtd: " + qtd + " | Price: " + price)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name)
		http.Redirect(w, r, route.ProdutosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateProdutoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Produto")
		id := r.FormValue("Id")
		name := r.FormValue("Name")

		sqlStatement := "UPDATE produtos SET name=$1 WHERE id=$2"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name)
		http.Redirect(w, r, route.ProdutosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteProdutoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete Produto")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM produtos WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.ProdutosRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func ListProdutosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Produtos")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, name FROM produtos order by id asc")
		sec.CheckInternalServerError(err, w)
		var produtos []mdl.Produto
		var produto mdl.Produto
		var i = 1
		for rows.Next() {
			err = rows.Scan(&produto.Id, &produto.Name)
			sec.CheckInternalServerError(err, w)
			produto.Order = i
			i++
			produtos = append(produtos, produto)
		}
		var page mdl.PageProdutos
		page.AppName = mdl.AppName
		page.Produtos = produtos
		page.Title = "Produtos"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/produtos/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Produtos", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
