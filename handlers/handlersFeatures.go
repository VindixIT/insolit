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

func CreateFeatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Create Feature")
		name := r.FormValue("Name")
		code := r.FormValue("Code")
		sqlStatement := "INSERT INTO features(name, code) VALUES ($1, $2) RETURNING id"
		id := 0
		err := Db.QueryRow(sqlStatement, name, code).Scan(&id)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		log.Println("INSERT: Id: " + strconv.Itoa(id) + " | Name: " + name + " | Code: " + code)
		http.Redirect(w, r, route.FeaturesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func UpdateFeatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Update Feature")
		id := r.FormValue("Id")
		name := r.FormValue("Name")
		code := r.FormValue("Code")
		sqlStatement := "UPDATE features SET name=$1, code=$2 WHERE id=$3"
		updtForm, err := Db.Prepare(sqlStatement)
		sec.CheckInternalServerError(err, w)
		if err != nil {
			panic(err.Error())
		}
		sec.CheckInternalServerError(err, w)
		updtForm.Exec(name, code, id)
		log.Println("UPDATE: Id: " + id + " | Name: " + name + " | Code: " + code)
		http.Redirect(w, r, route.FeaturesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteFeatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && sec.IsAuthenticated(w, r) {
		log.Println("Delete Feature")
		id := r.FormValue("Id")
		sqlStatement := "DELETE FROM features WHERE id=$1"
		deleteForm, err := Db.Prepare(sqlStatement)
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		sec.CheckInternalServerError(err, w)
		log.Println("DELETE: Id: " + id)
		http.Redirect(w, r, route.FeaturesRoute, 301)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

func DeleteFeaturesByRoleHandler(roleId string) {
	sqlStatement := "DELETE FROM features_roles WHERE role_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	deleteForm.Exec(roleId)
	log.Println("DELETE features_roles in Role Id: " + roleId)
}
func DeleteFeaturesHandler(diffDB []mdl.Feature) {
	sqlStatement := "DELETE FROM features_roles WHERE feature_id=$1"
	deleteForm, err := Db.Prepare(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	for n := range diffDB {
		deleteForm.Exec(strconv.FormatInt(int64(diffDB[n].Id), 10))
		log.Println("DELETE: Feature Id: " + strconv.FormatInt(int64(diffDB[n].Id), 10))
	}
}

func ListFeaturesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Features")
	if sec.IsAuthenticated(w, r) {
		rows, err := Db.Query("SELECT id, name, code FROM features order by id asc")
		sec.CheckInternalServerError(err, w)
		var features []mdl.Feature
		var feature mdl.Feature
		var i = 1
		for rows.Next() {
			err = rows.Scan(&feature.Id, &feature.Name, &feature.Code)
			sec.CheckInternalServerError(err, w)
			feature.Order = i
			i++
			features = append(features, feature)
		}
		var page mdl.PageFeatures
		page.AppName = mdl.AppName
		page.Features = features
		page.Title = "Funcionalidades"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/features/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Features", page)
		sec.CheckInternalServerError(err, w)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}

// AJAX
func ListFeaturesByRoleIdHandler(roleId string) []mdl.Feature {
	log.Println("List Features By Role Id")
	sql := "SELECT feature_id" +
		" FROM features_roles WHERE role_id= $1"
	log.Println(sql)
	rows, _ := Db.Query(sql, roleId)
	var features []mdl.Feature
	var feature mdl.Feature
	for rows.Next() {
		rows.Scan(&feature.Id)
		features = append(features, feature)
		log.Println(feature, feature)
	}
	return features
}
