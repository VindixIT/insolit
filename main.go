package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	dpk "insolit/db"
	hd "insolit/handlers"
	route "insolit/routes"
	sec "insolit/security"
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
	sec.Store = sessions.NewCookieStore([]byte("vindixit123581321"))
	hd.Db = dbConn()
	// injeta	ndo a variável Authenticated
	dpk.Initialize()
	r := mux.NewRouter()

	r.HandleFunc("/", hd.IndexHandler).Methods("GET")
	r.HandleFunc("/login", hd.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", hd.LogoutHandler).Methods("GET")
	r.HandleFunc("/inicio", hd.InicioHandler).Methods("GET")
	// ----------------- ACTIONS
	r.HandleFunc(route.ActionsRoute, hd.ListActionsHandler).Methods("GET")
	r.HandleFunc("/createAction", hd.CreateActionHandler).Methods("POST")
	r.HandleFunc("/updateAction", hd.UpdateActionHandler).Methods("POST")
	r.HandleFunc("/deleteAction", hd.DeleteActionHandler).Methods("POST")
	// ----------------- CLIENTES
	r.HandleFunc(route.ClientesRoute, hd.ListClientesHandler).Methods("GET")
	r.HandleFunc("/createCliente", hd.CreateClienteHandler).Methods("POST")
	r.HandleFunc("/updateCliente", hd.UpdateClienteHandler).Methods("POST")
	r.HandleFunc("/deleteCliente", hd.DeleteClienteHandler).Methods("POST")
	// ----------------- CONCESSIONARIAS
	r.HandleFunc(route.ConcessionariasRoute, hd.ListConcessionariasHandler).Methods("GET")
	r.HandleFunc("/createConcessionaria", hd.CreateConcessionariaHandler).Methods("POST")
	r.HandleFunc("/updateConcessionaria", hd.UpdateConcessionariaHandler).Methods("POST")
	r.HandleFunc("/deleteConcessionaria", hd.DeleteConcessionariaHandler).Methods("POST")
	// ----------------- CONCILIAÇÕES
	r.HandleFunc(route.ConciliacoesRoute, hd.ListConciliacoesHandler).Methods("GET")
	r.HandleFunc("/createConciliacao", hd.CreateConciliacaoHandler).Methods("POST")
	r.HandleFunc("/updateConciliacao", hd.UpdateConciliacaoHandler).Methods("POST")
	r.HandleFunc("/deleteConciliacao", hd.DeleteConciliacaoHandler).Methods("POST")
	// ----------------- CONTRATOS
	r.HandleFunc(route.ContratosRoute, hd.ListContratosHandler).Methods("GET")
	r.HandleFunc("/createContrato", hd.CreateContratoHandler).Methods("POST")
	r.HandleFunc("/deleteContrato", hd.DeleteContratoHandler).Methods("POST")
	// ----------------- CONTRATOS CONSUMO
	r.HandleFunc(route.ContratosConsumoRoute, hd.ListContratosConsumoHandler).Methods("GET")
	r.HandleFunc("/createContratoConsumo", hd.CreateContratoConsumoHandler).Methods("POST")
	r.HandleFunc("/updateContratoConsumo", hd.UpdateContratoConsumoHandler).Methods("POST")
	r.HandleFunc("/deleteContratoConsumo", hd.DeleteContratoConsumoHandler).Methods("POST")
	// ----------------- CONTRATOS GERAÇÃO
	r.HandleFunc(route.ContratosGeracaoRoute, hd.ListContratosGeracaoHandler).Methods("GET")
	r.HandleFunc("/createContratoGeracao", hd.CreateContratoGeracaoHandler).Methods("POST")
	r.HandleFunc("/updateContratoGeracao", hd.UpdateContratoGeracaoHandler).Methods("POST")
	r.HandleFunc("/deleteContratoGeracao", hd.DeleteContratoGeracaoHandler).Methods("POST")
	// ----------------- CONTRATOS REPARTICAO
	r.HandleFunc(route.ContratosReparticaoRoute, hd.ListContratosReparticaoHandler).Methods("GET")
	r.HandleFunc("/createContratoReparticao", hd.CreateContratoReparticaoHandler).Methods("POST")
	r.HandleFunc("/updateContratoReparticao", hd.UpdateContratoReparticaoHandler).Methods("POST")
	r.HandleFunc("/deleteContratoReparticao", hd.DeleteContratoReparticaoHandler).Methods("POST")
	// ----------------- CREDITOS
	r.HandleFunc(route.CreditosRoute, hd.ListCreditosHandler).Methods("GET")
	r.HandleFunc("/createCredito", hd.CreateCreditoHandler).Methods("POST")
	r.HandleFunc("/updateCredito", hd.UpdateCreditoHandler).Methods("POST")
	r.HandleFunc("/deleteCredito", hd.DeleteCreditoHandler).Methods("POST")
	// ----------------- FATURAS CONSUMO
	r.HandleFunc(route.FaturasConsumoRoute, hd.ListFaturasConsumoHandler).Methods("GET")
	r.HandleFunc("/createFaturaConsumo", hd.CreateFaturaConsumoHandler).Methods("POST")
	r.HandleFunc("/updateFaturaConsumo", hd.UpdateFaturaConsumoHandler).Methods("POST")
	r.HandleFunc("/deleteFaturaConsumo", hd.DeleteFaturaConsumoHandler).Methods("POST")
	// ----------------- FATURAS GERACAO
	r.HandleFunc(route.FaturasGeracaoRoute, hd.ListFaturasGeracaoHandler).Methods("GET")
	r.HandleFunc("/createFaturaGeracao", hd.CreateFaturaGeracaoHandler).Methods("POST")
	r.HandleFunc("/updateFaturaGeracao", hd.UpdateFaturaGeracaoHandler).Methods("POST")
	r.HandleFunc("/deleteFaturaGeracao", hd.DeleteFaturaGeracaoHandler).Methods("POST")
	// ----------------- FATURAS REPARTICAO
	r.HandleFunc(route.FaturasReparticaoRoute, hd.ListFaturasReparticaoHandler).Methods("GET")
	r.HandleFunc("/createFaturaReparticao", hd.CreateFaturaReparticaoHandler).Methods("POST")
	r.HandleFunc("/updateFaturaReparticao", hd.UpdateFaturaReparticaoHandler).Methods("POST")
	r.HandleFunc("/deleteFaturaReparticao", hd.DeleteFaturaReparticaoHandler).Methods("POST")
	// ----------------- INVERSORES
	r.HandleFunc(route.InversoresRoute, hd.ListInversoresHandler).Methods("GET")
	r.HandleFunc("/createInversor", hd.CreateInversorHandler).Methods("POST")
	r.HandleFunc("/updateInversor", hd.UpdateInversorHandler).Methods("POST")
	r.HandleFunc("/deleteInversor", hd.DeleteInversorHandler).Methods("POST")
	// ----------------- MÓDULOS
	r.HandleFunc(route.ModulosRoute, hd.ListModulosHandler).Methods("GET")
	r.HandleFunc("/createModulo", hd.CreateModuloHandler).Methods("POST")
	r.HandleFunc("/updateModulo", hd.UpdateModuloHandler).Methods("POST")
	r.HandleFunc("/deleteModulo", hd.DeleteModuloHandler).Methods("POST")
	// ----------------- USINAS
	r.HandleFunc(route.UsinasRoute, hd.ListUsinasHandler).Methods("GET")
	r.HandleFunc("/createUsina", hd.CreateUsinaHandler).Methods("POST")
	r.HandleFunc("/updateUsina", hd.UpdateUsinaHandler).Methods("POST")
	r.HandleFunc("/deleteUsina", hd.DeleteUsinaHandler).Methods("POST")
	// ----------------- WORKFLOWS
	r.HandleFunc(route.WorkflowsRoute, hd.ListWorkflowsHandler).Methods("GET")
	r.HandleFunc("/createWorkflow", hd.CreateWorkflowHandler).Methods("POST")
	r.HandleFunc("/updateWorkflow", hd.UpdateWorkflowHandler).Methods("POST")
	r.HandleFunc("/deleteWorkflow", hd.DeleteWorkflowHandler).Methods("POST")
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
	// ----------------- PARQUES
	r.HandleFunc(route.ParquesRoute, hd.ListParquesHandler).Methods("GET")
	r.HandleFunc("/createParque", hd.CreateParqueHandler).Methods("POST")
	r.HandleFunc("/updateParque", hd.UpdateParqueHandler).Methods("POST")
	r.HandleFunc("/deleteParque", hd.DeleteParqueHandler).Methods("POST")
	// ----------------- USERS
	r.HandleFunc(route.UsersRoute, hd.ListUsersHandler).Methods("GET")
	r.HandleFunc("/createUser", hd.CreateUserHandler).Methods("POST")
	r.HandleFunc("/updateUser", hd.UpdateUserHandler).Methods("POST")
	r.HandleFunc("/deleteUser", hd.DeleteUserHandler).Methods("POST")
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
