package main

import (
	"database/sql"
	"fmt"
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
	http.HandleFunc("/", hd.IndexHandler)
	http.HandleFunc("/login", hd.LoginHandler)
	http.HandleFunc("/logout", hd.LogoutHandler)
	// ----------------- WORKFLOWS
	http.HandleFunc("/workflows2", hd.ListWorkflowsHandler)
	http.HandleFunc(route.WorkflowsRoute, hd.ListWorkflowsHandler)
	http.HandleFunc("/createWorkflow", hd.CreateWorkflowHandler)
	http.HandleFunc("/updateWorkflow", hd.UpdateWorkflowHandler)
	http.HandleFunc("/deleteWorkflow", hd.DeleteWorkflowHandler)
	// ----------------- ACTIONS
	http.HandleFunc("/actions2", hd.ListActionsHandler)
	http.HandleFunc(route.ActionsRoute, hd.ListActionsHandler)
	http.HandleFunc("/createAction", hd.CreateActionHandler)
	http.HandleFunc("/updateAction", hd.UpdateActionHandler)
	http.HandleFunc("/deleteAction", hd.DeleteActionHandler)
	// ----------------- STATUS
	http.HandleFunc("/status2", hd.ListStatusHandler)
	http.HandleFunc(route.StatusRoute, hd.ListStatusHandler)
	http.HandleFunc("/createStatus", hd.CreateStatusHandler)
	http.HandleFunc("/updateStatus", hd.UpdateStatusHandler)
	http.HandleFunc("/deleteStatus", hd.DeleteStatusHandler)
	// ----------------- FEATURES
	http.HandleFunc("/features2", hd.ListFeaturesHandler)
	http.HandleFunc(route.FeaturesRoute, hd.ListFeaturesHandler)
	http.HandleFunc("/createFeature", hd.CreateFeatureHandler)
	http.HandleFunc("/updateFeature", hd.UpdateFeatureHandler)
	http.HandleFunc("/deleteFeature", hd.DeleteFeatureHandler)
	// ----------------- ROLES
	http.HandleFunc("/roles2", hd.ListRolesHandler)
	http.HandleFunc(route.RolesRoute, hd.ListRolesHandler)
	http.HandleFunc("/createRole", hd.CreateRoleHandler)
	http.HandleFunc("/updateRole", hd.UpdateRoleHandler)
	http.HandleFunc("/deleteRole", hd.DeleteRoleHandler)
	// ----------------- PRODUTOS
	http.HandleFunc("/produtos2", hd.ListProdutosHandler)
	http.HandleFunc(route.ProdutosRoute, hd.ListProdutosHandler)
	http.HandleFunc("/createProduto", hd.CreateProdutoHandler)
	http.HandleFunc("/updateProduto", hd.UpdateProdutoHandler)
	http.HandleFunc("/deleteProduto", hd.DeleteProdutoHandler)
	// ----------------- CLIENTES
	http.HandleFunc("/clientes2", hd.ListClientesHandler)
	http.HandleFunc(route.ClientesRoute, hd.ListClientesHandler)
	http.HandleFunc("/createCliente", hd.CreateClienteHandler)
	http.HandleFunc("/updateCliente", hd.UpdateClienteHandler)
	http.HandleFunc("/deleteCliente", hd.DeleteClienteHandler)
	// ----------------- PARQUES
	http.HandleFunc("/parques2", hd.ListParquesHandler)
	http.HandleFunc(route.ParquesRoute, hd.ListParquesHandler)
	http.HandleFunc("/createParque", hd.CreateParqueHandler)
	http.HandleFunc("/updateParque", hd.UpdateParqueHandler)
	http.HandleFunc("/deleteParque", hd.DeleteParqueHandler)
	// ----------------- CONCESSIONARIAS
	http.HandleFunc("/concessionarias2", hd.ListConcessionariasHandler)
	http.HandleFunc(route.ConcessionariasRoute, hd.ListConcessionariasHandler)
	http.HandleFunc("/createConcessionaria", hd.CreateConcessionariaHandler)
	http.HandleFunc("/updateConcessionaria", hd.UpdateConcessionariaHandler)
	http.HandleFunc("/deleteConcessionaria", hd.DeleteConcessionariaHandler)
	// ----------------- USERS
	http.HandleFunc("/users2", hd.ListUsersHandler)
	http.HandleFunc(route.UsersRoute, hd.ListUsersHandler)
	http.HandleFunc("/createUser", hd.CreateUserHandler)
	http.HandleFunc("/updateUser", hd.UpdateUserHandler)
	http.HandleFunc("/deleteUser", hd.DeleteUserHandler)
	// ----------------- ORDERS
	http.HandleFunc(route.OrdersRoute, hd.ListOrdersHandler)
	http.HandleFunc("/createOrder", hd.CreateOrderHandler)
	http.HandleFunc("/updateOrder", hd.UpdateOrderHandler)
	http.HandleFunc("/deleteOrder", hd.DeleteOrderHandler)
	// ----------------- ITEMS
	http.HandleFunc("/loadItemsByOrderId", hd.LoadItemsByOrderId)
	http.HandleFunc("/loadFeaturesByRoleId", hd.LoadFeaturesByRoleId)
	http.HandleFunc("/loadRolesByActionId", hd.LoadRolesByActionId)
	// ----------------- STATICS
	http.Handle("/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))),
	)
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
