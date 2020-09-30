package db

import (
	"database/sql"
	hd "insolit/handlers"
	"log"
)

var db *sql.DB

func Initialize() {
	db = hd.Db
	createSeq()
	createTable()
	createRoleAdmin()
	createFeatures()
	createStatusZERO()
	createRoleFeatures()
	createAdmin()
	createPKey()
	createFKey()
	createUniqueKey()
}

func createRoleAdmin() {
	query := " INSERT INTO roles (id, name) " +
		" SELECT 1, 'Admin' " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Admin')"
	log.Println(query)
	db.Exec(query)
}

func createFeatures() {
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (1, 'Listar Workflows', 'listWorkflows')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (2, 'Criar Workflow', 'createWorkflow')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (3, 'Listar Usuários', 'listUsers')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (4, 'Criar Usuário', 'createUser')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (5, 'Listar Papéis', 'listRoles')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (6, 'Criar Papel', 'createRole')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (7, 'Listar Status', 'listStatus')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (8, 'Criar Status', 'createStatus')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (9, 'Listar Funcionalidades', 'listFeatures')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (10, 'Criar Funcionalidade', 'createFeature')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (11, 'Listar Ações', 'listActions')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (12, 'Criar Ação', 'createAction')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (13, 'Listar Produtos', 'listProdutos')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (14, 'Criar Produto', 'createProduto')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (15, 'Listar Parques', 'listParques')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (16, 'Criar Parque', 'createParque')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (17, 'Listar Concessionarias', 'listConcessionarias')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (18, 'Criar Concessionaria', 'createConcessionaria')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (19, 'Listar Clientes', 'listClientes')")
	db.Exec("INSERT INTO public.features (id, name, code) VALUES (20, 'Criar Cliente', 'createCliente')")
}

func createRoleFeatures() {
	query := " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 1) "
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 2) "
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 3) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 4) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 5) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 6) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 7) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 8) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1, 9) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,10) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,11) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,12) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,13) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,14) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,15) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,16) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,17) "
	log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,18) "
	//log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,19) "
	log.Println(query)
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) VALUES (1,20) "
	log.Println(query)
	db.Exec(query)
}

func createStatusZERO() {
	query := "INSERT INTO status (id, name, stereotype)" +
		" SELECT 0, '-', '' " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE id = 0)"
	log.Println(query)
	db.Exec(query)
}

func createUniqueKey() {
	db.Exec(" ALTER TABLE ONLY public.actions_status" +
		" ADD CONSTRAINT action_status_unique_key UNIQUE (action_id, origin_status_id, destination_status_id)")

	db.Exec(" ALTER TABLE ONLY public.features_roles" +
		" ADD CONSTRAINT feature_role_unique_key UNIQUE (feature_id, role_id)")

	db.Exec(" ALTER TABLE ONLY public.users" +
		" ADD CONSTRAINT username_unique_key UNIQUE (username)")

	db.Exec(" ALTER TABLE ONLY public.workflows_entities" +
		" ADD CONSTRAINT entity_unique_index UNIQUE (entity_name)")

	db.Exec(" ALTER TABLE ONLY public.activities_roles" +
		" ADD CONSTRAINT action_role_unique_key UNIQUE (action_id, role_id)")
}

func createFKey() {
	db.Exec("ALTER TABLE ONLY public.items" +
		" ADD CONSTRAINT beers_fkey FOREIGN KEY (beer_id)" +
		" REFERENCES public.beers(id) ON UPDATE RESTRICT ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.items" +
		" ADD CONSTRAINT orders_fkey FOREIGN KEY (order_id)" +
		" REFERENCES public.orders(id) ON UPDATE RESTRICT ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.orders" +
		" ADD CONSTRAINT users_fkey FOREIGN KEY (user_id)" +
		" REFERENCES public.users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY public.actions " +
		" ADD CONSTRAINT destination_status_fkey FOREIGN KEY (destination_status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions " +
		" ADD CONSTRAINT origin_status_fkey FOREIGN KEY (origin_status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions " +
		" ADD CONSTRAINT workflows_fkey FOREIGN KEY (workflow_id)" +
		" REFERENCES public.workflows (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.features_roles " +
		" ADD CONSTRAINT features_fkey FOREIGN KEY (feature_id)" +
		" REFERENCES public.features (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.features_roles " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES public.roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.activities_roles " +
		" ADD CONSTRAINT activities_fkey FOREIGN KEY (activity_id)" +
		" REFERENCES public.activities (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.activities_roles " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES public.roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.status " +
		" ADD CONSTRAINT workflows_fkey FOREIGN KEY (workflow_id)" +
		" REFERENCES public.workflows (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions_status " +
		" ADD CONSTRAINT actions_fkey FOREIGN KEY (action_id)" +
		" REFERENCES public.actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions_status " +
		" ADD CONSTRAINT origin_status_fkey FOREIGN KEY (origin_status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.actions_status " +
		" ADD CONSTRAINT destination_status_fkey FOREIGN KEY (destination_status_id)" +
		" REFERENCES public.status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")
}

func createPKey() {
	db.Exec("ALTER TABLE ONLY public.produtos ADD CONSTRAINT produtos_pkey PRIMARY KEY (id)")
	//db.Exec("ALTER TABLE ONLY public.beers ADD CONSTRAINT beers_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.roles ADD CONSTRAINT roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.features ADD CONSTRAINT features_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.items ADD CONSTRAINT items_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.orders ADD CONSTRAINT order_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.users ADD CONSTRAINT users_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.workflows ADD CONSTRAINT workflows_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.status ADD CONSTRAINT status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.actions ADD CONSTRAINT actions_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.actions_status ADD CONSTRAINT actions_status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.workflows_entities ADD CONSTRAINT workflows_entities_pkey PRIMARY KEY (entity_name, workflow_id)")
}

func createAdmin() {
	query := "INSERT INTO users (id, username, password, email, mobile, name, role_id)" +
		" SELECT 1, 'aria', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy'," +
		" 'aria@vindixit.com', '61 984385415', 'Ária Ohashi', 1 " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'aria')"
	log.Println(query)
	db.Exec(query)
}

//func createBeer() {
//	query := "INSERT INTO beers (id, name, qtd, price)" +
//		" SELECT 1, 'Molson', 100, 100" +
//		" WHERE NOT EXISTS (SELECT name FROM beers WHERE name = 'Molson')"
//	log.Println(query)
//	db.Exec(query)
//
//}

func createSeq() {
	// Sequence PRODUTOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.produtos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence PARQUES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.parques_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CONCESSIONARIAS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.concessionarias_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CLIENTES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.clientes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIONS_STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.actions_status_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.features_roles_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence WORKFLOWS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.workflows_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.status_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIONS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.actions_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.roles_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FEATURES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.features_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence BEERS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.beers_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence USERS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.users_id_seq " +
		" START WITH 2" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ITEMS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.items_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ORDERS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.orders_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
}

func createTable() {
	// Table PRODUTOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.produtos (" +
			" id integer DEFAULT nextval('produtos_id_seq'::regclass)," +
			" name character varying(255) NOT NULL)")
	// Table CLIENTES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.clientes (" +
			" id integer DEFAULT nextval('clientes_id_seq'::regclass)," +
			" name character varying(255)," +
			" endereco character varying(255)," +
			" capacidade character varying(255)," +
			" cnpj character varying(255))")
	// Table PARQUES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.parques (" +
			" id integer DEFAULT nextval('parques_id_seq'::regclass)," +
			" name character varying(255) NOT NULL," +
			" endereco character varying(255)," +
			" cidade character varying(255)," +
			" estado character varying(255))")
	// Table CONCESSIONARIAS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.concessionarias (" +
			" id integer DEFAULT nextval('concessionarias_id_seq'::regclass)," +
			" name character varying(255) NOT NULL," +
			" cnpj character varying(255) )")
	// Table ACTIONS_ROLES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.activities_roles (" +
			" id integer DEFAULT nextval('activities_roles_id_seq'::regclass)," +
			" activities_id integer," +
			" role_id integer)")
	// Table FEATURES_ROLES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.features_roles (" +
			" id integer DEFAULT nextval('features_roles_id_seq'::regclass)," +
			" feature_id integer," +
			" role_id integer)")

	// Table WORKFLOWS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.workflows  (" +
			" id integer DEFAULT nextval('public.workflows_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL)")
	// Table ACTIONS
	db.Exec(" CREATE TABLE IF NOT EXISTS public.actions (" +
		"id integer DEFAULT nextval('public.actions_id_seq'::regclass) NOT NULL, " +
		"name character varying(255) NOT NULL, " +
		"origin_status_id integer, " +
		"destination_status_id integer, " +
		"other_than boolean)")
	// Table STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.status  (" +
			" id integer DEFAULT nextval('public.status_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" stereotype character varying(255) NULL)")
	// Table ROLES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.roles  (" +
			" id integer DEFAULT nextval('public.roles_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL)")
	// Table FEATURES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.features  (" +
			" id integer DEFAULT nextval('public.features_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" code character varying(255) NOT NULL)")
	// Table USERS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.users (" +
			" id integer DEFAULT nextval('public.users_id_seq'::regclass) NOT NULL," +
			" username character varying(255) NOT NULL," +
			" password character varying(255) NOT NULL," +
			" email character varying(255) NOT NULL," +
			" mobile character varying(255) NOT NULL," +
			" name character varying(255)," +
			" role_id integer)")

	// Table ORDERS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.orders (" +
			" id integer DEFAULT nextval('public.orders_id_seq'::regclass) NOT NULL," +
			" user_id integer," +
			" ordered_at timestamp without time zone," +
			" take_out_at timestamp without time zone)")
	// Table WORKFLOWS_ENTITIES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.workflows_entities (" +
			" entity_name character varying(255) COLLATE pg_catalog.'default' NOT NULL," +
			" workflow_id integer NOT NULL)")
	// Table ACTIONS_STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.actions_status (" +
			" id integer DEFAULT nextval('actions_status_id_seq'::regclass)," +
			" action_id integer," +
			" origin_status_id integer," +
			" destination_status_id integer)")
}
