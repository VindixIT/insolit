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
	createStatusZERO()
	createAdmin()
	createBeer()
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
	db.Exec("ALTER TABLE ONLY public.beers ADD CONSTRAINT beers_pkey PRIMARY KEY (id)")
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
	query := "INSERT INTO users (id, username, password, email, mobile, name)" +
		" SELECT 1, 'aria', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy', 'aria@vindixit.com', '61 984385415', 'Ária Ohashi'" +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'aria')"
	log.Println(query)
	db.Exec(query)
}

func createBeer() {
	query := "INSERT INTO beers (id, name, qtd, price)" +
		" SELECT 1, 'Molson', 100, 100" +
		" WHERE NOT EXISTS (SELECT name FROM beers WHERE name = 'Molson')"
	log.Println(query)
	db.Exec(query)

}

func createSeq() {
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
	// Table ACTIONS_ROLES
	db.Exec(
		" CREATE TABLE public.activities_roles (" +
			" id integer DEFAULT nextval('activities_roles_id_seq'::regclass)," +
			" activities_id integer," +
			" role_id integer)")
	// Table FEATURES_ROLES
	db.Exec(
		" CREATE TABLE public.features_roles (" +
			" id integer DEFAULT nextval('features_roles_id_seq'::regclass)," +
			" feature_id integer," +
			" role_id integer)")
	// Table WORKFLOWS_ENTITIES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.beers  (" +
			" id integer DEFAULT nextval('public.beers_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" qtd integer," +
			" price double precision)")
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
	// Table ITEMS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.items (" +
			" id integer DEFAULT nextval('public.items_id_seq'::regclass) NOT NULL," +
			" quantity double precision," +
			" beer_id integer," +
			" price double precision," +
			" item_value double precision," +
			" order_id integer)")
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
