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
	createFeaturesRoles()
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
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Workflows', 'listWorkflows' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listWorkflows')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Workflow', 'createWorkflow' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createWorkflow')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Usuários', 'listUsers' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listUsers')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Usuário', 'createUser' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createUser')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Papéis', 'listRoles' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listRoles')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Papel', 'createRole' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createRole')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Status', 'listStatus' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listStatus')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Status', 'createStatus' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createStatus')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Funcionalidades', 'listFeatures' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listFeatures')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Funcionalidade', 'createFeature' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createFeature')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Ações', 'listActions' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listActions')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Ação', 'createAction' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createAction')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Produtos', 'listProdutos' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listProdutos')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Produto', 'createProduto' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createProduto')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Parques', 'listParques' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listParques')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Parque', 'createParque' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createParque')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Concessionarias', 'listConcessionarias' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listConcessionarias')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Concessionaria', 'createConcessionaria' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createConcessionaria')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Clientes', 'listClientes' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listClientes')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Cliente', 'createCliente' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createCliente')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Contratos de Consumo', 'listContratosConsumo' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listContratosConsumo')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Contrato de Consumo', 'createContratoConsumo' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createContratoConsumo')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Contratos de Geração', 'listContratosGeracao' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listContratosGeracao')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Contrato de Geração', 'createContratoGeracao' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createContratoGeracao')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Contratos de Repartição', 'listContratosReparticao' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listContratosReparticao')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Contrato de Repartição', 'createContratoReparticao' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createContratoReparticao')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Créditos', 'listCreditos' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listCreditos')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Crédito', 'createCredito' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createCredito')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Faturas de Consumo', 'listFaturasConsumo' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listFaturasConsumo')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Fatura de Consumo', 'createFaturaConsumo' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createFaturaConsumo')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Faturas de Geração', 'listFaturasGeracao' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listFaturasGeracao')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Fatura de Geração', 'createFaturaGeracao' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createFaturaGeracao')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Faturas de Repartição', 'listFaturasReparticao' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listFaturasReparticao')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Fatura de Repartição', 'createFaturaReparticao' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createFaturaReparticao')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Inversores', 'listInversores' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listInversores')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Inversor', 'createInversor' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createInversor')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Módulos', 'listModulos' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listModulos')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Módulo', 'createModulo' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createModulo')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Listar Usinas', 'listUsinas' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listUsinas')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Criar Usina', 'createUsina' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createUsina')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Menu Acordos', 'acordos' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'acordos')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Menu Finanças', 'financeiros' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'financeiros')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Menu Relatórios', 'relatorios' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'relatorios')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Menu Investimentos', 'investimentos' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'investimentos')")
	db.Exec("INSERT INTO public.features (name, code) SELECT 'Menu Administração do Sistema', 'admin' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'admin')")
}

func createFeaturesRoles() {
	query := " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 1 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '11')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 2 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '12')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 3 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '13')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 4 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '14')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 5 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '15')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 6 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '16')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 7 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '17')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 8 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '18')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 9 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '19')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 10 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '110')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 11 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '111')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 12 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '112')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 13 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '113')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 14 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '114')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 15 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '115')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 16 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '116')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 17 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '117')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 18 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '118')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 19 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '119')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 20 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '120')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 21 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '121')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 22 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '122')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 23 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '123')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 24 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '124')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 25 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '125')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 26 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '126')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 27 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '127')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 28 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '128')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 29 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '129')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 30 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '130')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 31 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '131')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 32 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '132')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 33 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '133')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 34 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '134')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 35 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '135')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 36 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '136')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 37 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '137')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 38 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '138')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 39 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '139')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 40 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '140')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 41 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '141')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 42 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '142')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 43 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '143')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 44 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '144')"
	db.Exec(query)
	query = " INSERT INTO features_roles (role_id, feature_id) SELECT 1, 45 WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE text(role_id)||text(feature_id) = '145')"
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

	db.Exec("ALTER TABLE ONLY public.activities_roles " +
		" ADD CONSTRAINT activities_fkey FOREIGN KEY (activity_id)" +
		" REFERENCES public.activities (id) MATCH SIMPLE" +
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

	db.Exec("ALTER TABLE ONLY public.activities_roles " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES public.roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.conciliacoes " +
		" ADD CONSTRAINT creditos_fkey FOREIGN KEY (credito_id)" +
		" REFERENCES public.creditos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.conciliacoes " +
		" ADD CONSTRAINT faturas_consumo_fkey FOREIGN KEY (contrato_consumo_id)" +
		" REFERENCES public.faturas_consumo (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.contratos_consumo " +
		" ADD CONSTRAINT concessionarias_fkey FOREIGN KEY (concessionaria_id)" +
		" REFERENCES public.concessionarias (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.contratos_consumo " +
		" ADD CONSTRAINT clientes_fkey FOREIGN KEY (cliente_id)" +
		" REFERENCES public.clientes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.contratos_geracao " +
		" ADD CONSTRAINT concessionarias_fkey FOREIGN KEY (concessionaria_id)" +
		" REFERENCES public.concessionarias (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.contratos_geracao " +
		" ADD CONSTRAINT clientes_fkey FOREIGN KEY (cliente_id)" +
		" REFERENCES public.clientes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.contratos_geracao " +
		" ADD CONSTRAINT usinas_fkey FOREIGN KEY (usina_id)" +
		" REFERENCES public.usinas (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.contratos_reparticao " +
		" ADD CONSTRAINT produtos_fkey FOREIGN KEY (produto_id)" +
		" REFERENCES public.produtos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.contratos_reparticao " +
		" ADD CONSTRAINT clientes_fkey FOREIGN KEY (cliente_id)" +
		" REFERENCES public.clientes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.contratos_reparticao " +
		" ADD CONSTRAINT usinas_fkey FOREIGN KEY (usina_id)" +
		" REFERENCES public.usinas (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.creditos " +
		" ADD CONSTRAINT contratos_reparticao_fkey FOREIGN KEY (contrato_reparticao_id)" +
		" REFERENCES public.contratos_reparticao (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.faturas_consumo " +
		" ADD CONSTRAINT contratos_consumo_fkey FOREIGN KEY (contrato_consumo_id)" +
		" REFERENCES public.contratos_consumo (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.faturas_geracao " +
		" ADD CONSTRAINT contratos_geracao_fkey FOREIGN KEY (contrato_geracao_id)" +
		" REFERENCES public.contratos_geracao (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.faturas_reparticao " +
		" ADD CONSTRAINT contratos_reparticao_fkey FOREIGN KEY (contrato_reparticao_id)" +
		" REFERENCES public.contratos_reparticao (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.faturas_reparticao " +
		" ADD CONSTRAINT conciliacoes_fkey FOREIGN KEY (conciliacao_id)" +
		" REFERENCES public.conciliacoes (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.usinas " +
		" ADD CONSTRAINT modulos_fkey FOREIGN KEY (modulo_id)" +
		" REFERENCES public.modulos (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY public.usinas " +
		" ADD CONSTRAINT inversores_fkey FOREIGN KEY (inversor_id)" +
		" REFERENCES public.inversores (id) MATCH SIMPLE" +
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

	db.Exec("ALTER TABLE ONLY public.status " +
		" ADD CONSTRAINT workflows_fkey FOREIGN KEY (workflow_id)" +
		" REFERENCES public.workflows (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")
}

func createPKey() {
	db.Exec("ALTER TABLE ONLY public.actions ADD CONSTRAINT actions_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.actions_status ADD CONSTRAINT actions_status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.clientes ADD CONSTRAINT clientes_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.concessionarias ADD CONSTRAINT concessionarias_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.conciliacoes ADD CONSTRAINT conciliacoes_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.contratos_consumo ADD CONSTRAINT contratos_consumo_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.contratos_geracao ADD CONSTRAINT contratos_geracao_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.contratos_reparticao ADD CONSTRAINT contratos_reparticao_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.creditos ADD CONSTRAINT creditos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.faturas_consumo ADD CONSTRAINT faturas_consumo_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.faturas_geracao ADD CONSTRAINT faturas_geracao_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.faturas_reparticao ADD CONSTRAINT faturas_reparticao_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.features ADD CONSTRAINT features_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.inversores ADD CONSTRAINT inversores_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.modulos ADD CONSTRAINT modulos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.parques ADD CONSTRAINT parques_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.produtos ADD CONSTRAINT produtos_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.roles ADD CONSTRAINT roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.status ADD CONSTRAINT status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.users ADD CONSTRAINT users_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.usinas ADD CONSTRAINT usinas_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.workflows ADD CONSTRAINT workflows_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY public.workflows_entities ADD CONSTRAINT workflows_entities_pkey PRIMARY KEY (entity_name, workflow_id)")
}

func createAdmin() {
	query := "INSERT INTO users (id, username, password, email, mobile, name, role_id)" +
		" SELECT 1, 'aria', '$2a$14$C1DIYDsmE0QHjje4wR5uwOAC7m8/YAUe8DYw/yuKIAQgRDibeCDMy'," +
		" 'aria@vindixit.com', '61 984385415', 'Ária Ohashi', 1 " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'aria')"
	log.Println(query)
	db.Exec(query)
	query = "INSERT INTO users (id, username, password, email, mobile, name, role_id)" +
		" SELECT 2, 'marcelo', '$2a$10$RGZRB6DZBcMa/pMe4MrPbekkAMs2s9/eDm5Aa6nDauv5t75Yty.xO'," +
		" 'marcelo@vindixit.com', '61 984385415', 'Marcelo Serra Silva', 1 " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'marcelo')"
	log.Println(query)
	db.Exec(query)
}

func createSeq() {
	// Sequence ACTIONS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.actions_id_seq " +
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
	// Sequence ACTIVITIES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.activities_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence ACTIVITIES_ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.activities_roles_id_seq " +
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
	// Sequence CONCESSIONARIAS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.concessionarias_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CONCILIACOES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.conciliacoes_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CONTRATOS_CONSUMO_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.contratos_consumo_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CONTRATOS_GERACAO_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.contratos_geracao_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CONTRATOS_REPARTICAO_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.contratos_reparticao_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence CREDITOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.creditos_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FATURAS_CONSUMO_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.faturas_consumo_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FATURAS_GERACAO_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.faturas_geracao_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence FATURAS_REPARTICAO_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.faturas_reparticao_id_seq " +
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
	// Sequence FEATURES_ROLES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.features_roles_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence INVERSORES_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.inversores_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence MODULOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.modulos_id_seq " +
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
	// Sequence PRODUTOS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.produtos_id_seq " +
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
	// Sequence STATUS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.status_id_seq " +
		" START WITH 1" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence USERS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.users_id_seq " +
		" START WITH 3" +
		" INCREMENT BY 1" +
		" NO MINVALUE" +
		" NO MAXVALUE" +
		" CACHE 1")
	// Sequence USINAS_ID_SEQ
	db.Exec("CREATE SEQUENCE IF NOT EXISTS public.usinas_id_seq " +
		" START WITH 3" +
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
}

func createTable() {
	// Table ACTIONS
	db.Exec(" CREATE TABLE IF NOT EXISTS public.actions (" +
		"id integer DEFAULT nextval('public.actions_id_seq'::regclass) NOT NULL, " +
		"name character varying(255) NOT NULL, " +
		"origin_status_id integer, " +
		"destination_status_id integer, " +
		"other_than boolean)")
	// Table ACTIONS_STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.actions_status (" +
			" id integer DEFAULT nextval('actions_status_id_seq'::regclass)," +
			" action_id integer," +
			" origin_status_id integer," +
			" destination_status_id integer)")
	// Table ACTIVITIES
	db.Exec(" CREATE TABLE public.activities (" +
		" id integer NOT NULL DEFAULT nextval('activities_id_seq'::regclass)," +
		" workflow_id integer," +
		" action_id integer," +
		" expiration_action_id integer," +
		" expiration_time_days integer," +
		" start_at timestamp without time zone," +
		" end_at timestamp without time zone)")
	// Table ACTIVITIES_ROLES
	db.Exec(
		" CREATE TABLE public.activities_roles (" +
			" id integer DEFAULT nextval('activities_roles_id_seq'::regclass)," +
			" activity_id integer," +
			" role_id integer)")
	// Table CLIENTES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.clientes (" +
			" id integer DEFAULT nextval('clientes_id_seq'::regclass)," +
			" name character varying(255)," +
			" endereco character varying(255)," +
			" cidade character varying(255)," +
			" estado character varying(255)," +
			" cnpj character varying(255))")
	// Table CONCESSIONARIAS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.concessionarias (" +
			" id integer DEFAULT nextval('concessionarias_id_seq'::regclass)," +
			" name character varying(255) NOT NULL," +
			" cnpj character varying(255) )")
	// Table CONCILIACOES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.conciliacoes (" +
			" id integer DEFAULT nextval('conciliacoes_id_seq'::regclass)," +
			" creditos_id," +
			" fatura_id," +
			" data timestamp without time zone )")
	// Table CONTRATOS CONSUMO
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.contratos_consumo (" +
			" id integer DEFAULT nextval('contratos_consumo_id_seq'::regclass)," +
			" concessionaria_id integer," +
			" cliente_id integer," +
			" contrato_concessionaria character varying(255)," +
			" unidade_consumidora character varying(255)," +
			" endereco_uc character varying(255)," +
			" vencimento integer," +
			" assinatura_em character varying(255) )")
	// Table CONTRATOS GERACAO
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.contratos_geracao (" +
			" id integer DEFAULT nextval('contratos_geracao_id_seq'::regclass)," +
			" cliente_id integer," +
			" usina_id integer," +
			" concessionaria_id integer," +
			" contrato_concessionaria character varying(255)," +
			" unidade_consumidora character varying(255)," +
			" vencimento integer," +
			" assinatura_em timestamp without time zone)")
	// Table CONTRATOS REPARTICAO
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.contratos_reparticao (" +
			" id integer DEFAULT nextval('contratos_reparticao_id_seq'::regclass)," +
			" cliente_id integer," +
			" produto_id integer," +
			" usina_id integer," +
			" contrato_investidor character varying(255)," +
			" vencimento integer," +
			" assinatura_em timestamp without time zone)")
	// Table CREDITOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.creditos (" +
			" id integer DEFAULT nextval('creditos_id_seq'::regclass)," +
			" contrato_reparticao_id integer," +
			" produto_id integer," +
			" usina_id integer," +
			" energia_gerada double precision," +
			" inicio timestamp without time zone," +
			" termino timestamp without time zone)")
	// Table FATURAS CONSUMO
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.faturas_consumo (" +
			" id integer DEFAULT nextval('faturas_consumo_id_seq'::regclass)," +
			" contrato_consumo_id integer," +
			" consumo_mes double precision," +
			" credito_mes double precision," +
			" saldo_acumulado double precision," +
			" valor double precision," +
			" vencimento_em timestamp without time zone," +
			" emissao_em timestamp without time zone," +
			" leitura_anterior timestamp without time zone," +
			" leitura_atual timestamp without time zone," +
			" pago boolean)")
	// Table FATURAS GERACAO
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.faturas_geracao (" +
			" id integer DEFAULT nextval('faturas_geracao_id_seq'::regclass)," +
			" contrato_geracao_id integer," +
			" consumo_mes double precision," +
			" credito_mes double precision," +
			" valor double precision," +
			" vencimento_em timestamp without time zone," +
			" emissao_em timestamp without time zone," +
			" leitura_anterior timestamp without time zone," +
			" leitura_atual timestamp without time zone," +
			" pago boolean)")
	// Table FATURAS REPARTICAO
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.faturas_reparticao (" +
			" id integer DEFAULT nextval('faturas_reparticao_id_seq'::regclass)," +
			" contrato_reparticao_id integer," +
			" geracao_mes double precision," +
			" valor double precision," +
			" vencimento_em timestamp without time zone," +
			" emissao_em timestamp without time zone," +
			" leitura_anterior timestamp without time zone," +
			" leitura_atual timestamp without time zone," +
			" pago boolean)")
	// Table FEATURES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.features  (" +
			" id integer DEFAULT nextval('public.features_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" code character varying(255) NOT NULL)")
	// Table FEATURES_ROLES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.features_roles (" +
			" id integer DEFAULT nextval('features_roles_id_seq'::regclass)," +
			" feature_id integer," +
			" role_id integer)")
	// Table INVERSORES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.inversores (" +
			" id integer DEFAULT nextval('inversores_id_seq'::regclass)," +
			" fabricante character varying(255)," +
			" modelo character varying(255)," +
			" potencia_nominal double precision)")
	// Table MODULOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.modulos (" +
			" id integer DEFAULT nextval('modulos_id_seq'::regclass)," +
			" fabricante character varying(255)," +
			" modelo character varying(255)," +
			" potencia_pico double precision)")
	// Table PARQUES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.parques (" +
			" id integer DEFAULT nextval('parques_id_seq'::regclass)," +
			" name character varying(255) NOT NULL," +
			" endereco character varying(255)," +
			" cidade character varying(255)," +
			" estado character varying(255))")
	// Table PRODUTOS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.produtos (" +
			" id integer DEFAULT nextval('produtos_id_seq'::regclass)," +
			" name character varying(255) NOT NULL)")
	// Table ROLES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.roles  (" +
			" id integer DEFAULT nextval('public.roles_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL)")
	// Table STATUS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.status  (" +
			" id integer DEFAULT nextval('public.status_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL," +
			" stereotype character varying(255) NULL)")
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
	// Table USINAS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.usinas (" +
			" id integer DEFAULT nextval('usinas_id_seq'::regclass)," +
			" parque_id integer," +
			" inversor_id integer," +
			" modulo_id integer," +
			" potencia double precision," +
			" potencia_nominal double precision," +
			" energia_media double precision)")
	// Table WORKFLOWS
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.workflows  (" +
			" id integer DEFAULT nextval('public.workflows_id_seq'::regclass) NOT NULL," +
			" name character varying(255) NOT NULL)")
	// Table WORKFLOWS_ENTITIES
	db.Exec(
		" CREATE TABLE IF NOT EXISTS public.workflows_entities (" +
			" entity_name character varying(255) COLLATE pg_catalog.'default' NOT NULL," +
			" workflow_id integer NOT NULL)")
}
