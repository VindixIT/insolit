package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	dpk "insolit/db"
	hd "insolit/handlers"
	route "insolit/routes"
	"log"
	"net/http"
	"os"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func dbConn() *sql.DB {
	dbase, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	log.Println(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	return dbase
}

func main() {

	hd.Db = dbConn()
	// injeta	ndo a variável Authenticated
	dpk.Initialize()
	r := mux.NewRouter()
	r.HandleFunc("/", hd.IndexHandler).Methods("GET")
	r.HandleFunc("/login", hd.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", hd.LogoutHandler).Methods("GET")
	// ----------------- WORKFLOWS
	r.HandleFunc(route.WorkflowsRoute, hd.ListWorkflowsHandler).Methods("GET")
	r.HandleFunc("/createWorkflow", hd.CreateWorkflowHandler).Methods("POST")
	r.HandleFunc("/updateWorkflow", hd.UpdateWorkflowHandler).Methods("POST")
	r.HandleFunc("/deleteWorkflow", hd.DeleteWorkflowHandler).Methods("POST")
	// ----------------- ACTIONS
	r.HandleFunc(route.ActionsRoute, hd.ListActionsHandler).Methods("GET")
	r.HandleFunc("/createAction", hd.CreateActionHandler).Methods("POST")
	r.HandleFunc("/updateAction", hd.UpdateActionHandler).Methods("POST")
	r.HandleFunc("/deleteAction", hd.DeleteActionHandler).Methods("POST")
	// ----------------- STATUS
	r.HandleFunc(route.StatusRoute, hd.ListStatusHandler).Methods("GET")
	r.HandleFunc("/createStatus", hd.CreateStatusHandler).Methods("POST")
	r.HandleFunc("/updateStatus", hd.UpdateStatusHandler).Methods("POST")
	r.HandleFunc("/deleteStatus", hd.DeleteStatusHandler).Methods("POST")
	// ----------------- FEATURES
	r.HandleFunc(route.FeaturesRoute, hd.ListFeaturesHandler).Methods("GET")
	r.HandleFunc("/createFeature", hd.CreateFeatureHandler).Methods("POST")
	r.HandleFunc("/updateFeature", hd.UpdateFeatureHandler).Methods("POST")
	r.HandleFunc("/deleteFeature", hd.DeleteFeatureHandler).Methods("POST")
	// ----------------- ROLES
	r.HandleFunc(route.RolesRoute, hd.ListRolesHandler).Methods("GET")
	r.HandleFunc("/createRole", hd.CreateRoleHandler).Methods("POST")
	r.HandleFunc("/updateRole", hd.UpdateRoleHandler).Methods("POST")
	r.HandleFunc("/deleteRole", hd.DeleteRoleHandler).Methods("POST")
	// ----------------- PRODUTOS
	r.HandleFunc(route.ProdutosRoute, hd.ListProdutosHandler).Methods("GET")
	r.HandleFunc("/createProduto", hd.CreateProdutoHandler).Methods("POST")
	r.HandleFunc("/updateProduto", hd.UpdateProdutoHandler).Methods("POST")
	r.HandleFunc("/deleteProduto", hd.DeleteProdutoHandler).Methods("POST")
	// ----------------- CLIENTES
	r.HandleFunc(route.ClientesRoute, hd.ListClientesHandler).Methods("GET")
	r.HandleFunc("/createCliente", hd.CreateClienteHandler).Methods("POST")
	r.HandleFunc("/updateCliente", hd.UpdateClienteHandler).Methods("POST")
	r.HandleFunc("/deleteCliente", hd.DeleteClienteHandler).Methods("POST")
	// ----------------- PARQUES
	r.HandleFunc(route.ParquesRoute, hd.ListParquesHandler).Methods("GET")
	r.HandleFunc("/createParque", hd.CreateParqueHandler).Methods("POST")
	r.HandleFunc("/updateParque", hd.UpdateParqueHandler).Methods("POST")
	r.HandleFunc("/deleteParque", hd.DeleteParqueHandler).Methods("POST")
	// ----------------- CONCESSIONARIAS
	r.HandleFunc(route.ConcessionariasRoute, hd.ListConcessionariasHandler).Methods("GET")
	r.HandleFunc("/createConcessionaria", hd.CreateConcessionariaHandler).Methods("POST")
	r.HandleFunc("/updateConcessionaria", hd.UpdateConcessionariaHandler).Methods("POST")
	r.HandleFunc("/deleteConcessionaria", hd.DeleteConcessionariaHandler).Methods("POST")
	// ----------------- USERS
	r.HandleFunc(route.UsersRoute, hd.ListUsersHandler).Methods("GET")
	r.HandleFunc("/createUser", hd.CreateUserHandler).Methods("POST")
	r.HandleFunc("/updateUser", hd.UpdateUserHandler).Methods("POST")
	r.HandleFunc("/deleteUser", hd.DeleteUserHandler).Methods("POST")
	// ----------------- ORDERS
	r.HandleFunc(route.OrdersRoute, hd.ListOrdersHandler).Methods("GET")
	r.HandleFunc("/createOrder", hd.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/updateOrder", hd.UpdateOrderHandler).Methods("POST")
	r.HandleFunc("/deleteOrder", hd.DeleteOrderHandler).Methods("POST")
	// ----------------- ITEMS
	r.HandleFunc("/loadItemsByOrderId", hd.LoadItemsByOrderId)
	r.HandleFunc("/loadFeaturesByRoleId", hd.LoadFeaturesByRoleId)
	r.HandleFunc("/loadRolesByActionId", hd.LoadRolesByActionId)
	// ----------------- STATICS
	http.Handle("/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))),
	)
	http.Handle("/", r)
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
	defer hd.Db.Close()
}
