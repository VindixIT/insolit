package handlers

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	mdl "insolit/models"
	route "insolit/routes"
	sec "insolit/security"
	"log"
	"net/http"
)

var Db *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if sec.IsAuthenticated(w, r) {
		http.Redirect(w, r, route.BeersRoute, 200)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Logout Handler")
	session, _ := sec.Store.Get(r, sec.CookieName)
	delete(session.Values, "user")
	session.Options.MaxAge = -1
	_ = session.Save(r, w)
	http.ServeFile(w, r, "tmpl/login.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "tmpl/login.html")
		return
	}
	username := r.FormValue("usrname")
	password := r.FormValue("psw")
	var user mdl.User
	// bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := Db.QueryRow("SELECT id, "+
		" username, password, COALESCE(role_id, 0)"+
		" FROM users WHERE username=$1", &username).Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	sec.CheckInternalServerError(err, w)
	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("erro do /login")
		http.Redirect(w, r, "/login", 301)
	}

	query := "SELECT " +
		"A.feature_id, B.code FROM features_roles A, features B " +
		"WHERE A.feature_id = B.id AND A.role_id = $1"
	log.Println("Query: " + query)
	rows, _ := Db.Query(query, user.Role)
	var features []mdl.Feature
	var feature mdl.Feature
	for rows.Next() {
		rows.Scan(&feature.Id, &feature.Code)
		features = append(features, feature)
		log.Println(feature)
	}
	user.Features = features

	AddUserInCookie(w, r, user)
	// Abrindo o Cookie
	savedUser := GetUserInCookie(w, r)
	log.Println("MAIN Saved User is " + savedUser.Username)
	http.Redirect(w, r, route.OrdersRoute, 301)
}

func GetUserInCookie(w http.ResponseWriter, r *http.Request) mdl.User {
	session, _ := sec.Store.Get(r, sec.CookieName)
	var savedUser mdl.User
	sessionUser := session.Values["user"]
	if sessionUser != nil {
		strUser := sessionUser.(string)
		json.Unmarshal([]byte(strUser), &savedUser)
	}
	return savedUser
}

func AddUserInCookie(w http.ResponseWriter, r *http.Request, user mdl.User) {
	session, _ := sec.Store.Get(r, sec.CookieName)
	bytesUser, _ := json.Marshal(&user)
	session.Values["user"] = string(bytesUser)
	session.Save(r, w)
}

func BuildLoggedUser(user mdl.User) mdl.LoggedUser {
	var loggedUser mdl.LoggedUser
	loggedUser.User = user
	loggedUser.HasPermission = func(feature string) bool {
		for _, value := range user.Features {
			if value.Code == feature {
				// log.Println("PASSOU: " + feature)
				return true
			}
		}
		return false
	}
	return loggedUser
}
