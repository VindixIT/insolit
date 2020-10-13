package models

import (
	"github.com/lib/pq"
	//	"time"
)

var AppName = "Insolit"

type Action struct {
	Order         int
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	OriginId      int64  `json:"originid"`
	Origin        string `json:"originName"`
	DestinationId int64  `json:"destinationid"`
	Destination   string `json:"destinationName"`
	OtherThan     bool   `json:"otherthan"`
	Roles         []Role
}

type Cliente struct {
	Order    int
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Endereco string `json:"endereco"`
	Cidade   string `json:"cidade"`
	Estado   string `json:"estado"`
	Cnpj     string `json:"cnpj"`
}

type Concessionaria struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Cnpj  string `json:"cnpj"`
}

type Conciliacao struct {
	Order         int
	Id            int64  `json:"id"`
	Credito       int64  `json:"creditoId"`
	FaturaConsumo int64  `json:"faturaConsumoId"`
	Name          string `json:"name"`
}

type ContratoConsumo struct {
	Order                  int
	Id                     int64    `json:"id"`
	ClienteId              int64    `json:"clienteId"`
	ClienteName            string   `json:"clienteName"`
	ConcessionariaId       int64    `json:"concessionariaId"`
	ConcessionariaName     string   `json:"concessionariaName"`
	UnidadeConsumidora     string   `json:"unidadeConsumidora"`
	ContratoConcessionaria string   `json:"contratoConcessionaria"`
	VencimentoEm           NullTime `json:"vencimentoEm"`
	AssinaturaEm           NullTime `json:"assinaturaEm"`
}

type ContratoGeracao struct {
	Order                  int
	Id                     int64    `json:"id"`
	ClienteId              int64    `json:"clienteId"`
	ClienteName            string   `json:"clienteName"`
	UsinaId                int64    `json:"usinaId"`
	UsinaName              string   `json:"usinaName"`
	UnidadeConsumidora     string   `json:"unidadeConsumidora"`
	ConcessionariaId       int64    `json:"concessionariaId"`
	ConcessionariaName     string   `json:"concessionariaName"`
	ContratoConcessionaria string   `json:"contratoConcessionaria"`
	VencimentoEm           NullTime `json:"vencimentoEm"`
	AssinaturaEm           NullTime `json:"assinaturaEm"`
}

type ContratoReparticao struct {
	Order              int
	Id                 int64    `json:"id"`
	ClienteId          int64    `json:"clienteId"`
	ClienteName        string   `json:"clienteName"`
	ProdutoId          int64    `json:"usinaId"`
	ProdutoName        string   `json:"usinaName"`
	UsinaId            int64    `json:"usinaId"`
	UsinaName          string   `json:"usinaName"`
	ContratoInvestidor string   `json:"unidadeConsumidora"`
	VencimentoEm       NullTime `json:"vencimentoEm"`
	AssinaturaEm       NullTime `json:"assinaturaEm"`
}

type Credito struct {
	Order                int
	Id                   int64    `json:"id"`
	ContratoReparticaoId string   `json:"contratoReparticaoId"`
	EnergiaGerada        float64  `json:"energiaGerada"`
	Inicio               NullTime `json:"inicio"`
	Termino              NullTime `json:"termino"`
}

type FaturaConsumo struct {
	Order              int
	Id                 int64    `json:"id"`
	ContratoConsumoId  int64    `json:"contratoConsumo"`
	ClienteName        string   `json:"clienteName"`
	ConcessionariaName string   `json:"concessionariaName"`
	ConsumoMes         float64  `json:"consumoMes"`
	CreditoMes         float64  `json:"creditoMes"`
	SaldoAcumulado     float64  `json:"saldoAcumulado"`
	VencimentoEm       NullTime `json:"vencimentoEm"`
	EmissaoEm          NullTime `json:"emissaoEm"`
	Valor              float64  `json:"valor"`
	LeituraAnterior    float64  `json:"leituraAnterior"`
	LeituraAtual       float64  `json:"leituraAtual"`
	Pago               bool     `json:"pago"`
}

type FaturaGeracao struct {
	Order             int
	Id                int64    `json:"id"`
	ContratoGeracaoId int64    `json:"contratoGeracao"`
	ConsumoMes        float64  `json:"consumoMes"`
	CreditoMes        float64  `json:"creditoMes"`
	LeituraAnterior   float64  `json:"leituraAnterior"`
	LeituraAtual      float64  `json:"leituraAtual"`
	Valor             float64  `json:"valor"`
	VencimentoEm      NullTime `json:"vencimentoEm"`
	EmissaoEm         NullTime `json:"emissaoEm"`
	Pago              bool     `json:"pago"`
}

type FaturaReparticao struct {
	Order                int
	Id                   int64    `json:"id"`
	ContratoReparticaoId int64    `json:"contratoReparticao"`
	ConciliacaoId        int64    `json:"conciliacaoId"`
	ConsumoMes           float64  `json:"consumoMes"`
	GeracaoMes           float64  `json:"geracaoMes"`
	LeituraAnterior      float64  `json:"leituraAnterior"`
	LeituraAtual         float64  `json:"leituraAtual"`
	Valor                float64  `json:"valor"`
	VencimentoEm         NullTime `json:"vencimentoEm"`
	EmissaoEm            NullTime `json:"emissaoEm"`
	Pago                 bool     `json:"pago"`
}

type Feature struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
}

type Inversor struct {
	Order           int
	Id              int64  `json:"id"`
	Modelo          string `json:"modelo"`
	Fabricante      string `json:"fabricante"`
	PotenciaNominal string `json:"potenciaNominal"`
}

type LoggedUser struct {
	User          User
	HasPermission func(string) bool
}

type Modulo struct {
	Order        int
	Id           int64  `json:"id"`
	Modelo       string `json:"modelo"`
	Fabricante   string `json:"fabricante"`
	PotenciaPico string `json:"potenciaPico"`
}

type NullTime struct {
	pq.NullTime
}

type Parque struct {
	Order    int
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Endereco string `json:"endereco"`
	Cidade   string `json:"cidade"`
	Estado   string `json:"estado"`
}

type Produto struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

type Role struct {
	Order    int
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Selected bool
	Features []Feature
}

type Status struct {
	Order      int
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Stereotype string `json:"stereotype"`
}

type User struct {
	Order    int       `json:"order"`
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Mobile   string    `json:"mobile"`
	Role     int64     `json:"role"`
	RoleName string    `json:"rolename"`
	Features []Feature `json:"features"`
	Selected bool      `json:"selected"`
}

type Usina struct {
	Order           int
	Id              int64  `json:"id"`
	ParqueId        int64  `json:"parqueId"`
	ParqueNome      string `json:"parqueNome"`
	Potencia        string `json:"potencia"`
	PotenciaNominal string `json:"potenciaNominal"`
	EnergiaMedia    string `json:"energiaMedia"`
	InversorId      int64  `json:"inversorId"`
	InversorNome    string `json:"inversorNome"`
	ModuloId        int64  `json:"moduloId"`
	ModuloNome      string `json:"moduloNome"`
}

type Workflow struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

type PageActions struct {
	AppName    string
	Title      string
	Statuss    []Status
	Actions    []Action
	LoggedUser LoggedUser
}

type PageClientes struct {
	AppName    string
	Title      string
	Clientes   []Cliente
	LoggedUser LoggedUser
}

type PageConcessionarias struct {
	AppName         string
	Title           string
	Concessionarias []Concessionaria
	LoggedUser      LoggedUser
}

type PageConciliacoes struct {
	AppName        string
	Title          string
	Conciliacoes   []Conciliacao
	Creditos       []Credito
	FaturasConsumo []FaturaConsumo
	LoggedUser     LoggedUser
}

type PageContratosConsumo struct {
	AppName          string
	Title            string
	ContratosConsumo []ContratoConsumo
	Concessionarias  []Concessionaria
	Clientes         []Cliente
	LoggedUser       LoggedUser
}

type PageContratosGeracao struct {
	AppName          string
	Title            string
	ContratosGeracao []ContratoGeracao
	Concessionarias  []Concessionaria
	Usinas           []Usina
	Clientes         []Cliente
	LoggedUser       LoggedUser
}

type PageContratosReparticao struct {
	AppName             string
	Title               string
	ContratosReparticao []ContratoReparticao
	Concessionarias     []Concessionaria
	Usinas              []Usina
	Clientes            []Cliente
	LoggedUser          LoggedUser
}

type PageCreditos struct {
	AppName             string
	Title               string
	Creditos            []Credito
	ContratosReparticao []ContratoReparticao
	LoggedUser          LoggedUser
}

type PageFaturasConsumo struct {
	AppName          string
	Title            string
	ContratosConsumo []ContratoConsumo
	FaturasConsumo   []FaturaConsumo
	LoggedUser       LoggedUser
}

type PageFaturasGeracao struct {
	AppName          string
	Title            string
	ContratosGeracao []ContratoGeracao
	FaturasGeracao   []FaturaGeracao
	LoggedUser       LoggedUser
}

type PageFaturasReparticao struct {
	AppName             string
	Title               string
	ContratosReparticao []ContratoReparticao
	FaturasReparticao   []FaturaReparticao
	LoggedUser          LoggedUser
}

type PageFeatures struct {
	AppName    string
	Title      string
	Features   []Feature
	LoggedUser LoggedUser
}

type PageInicio struct {
	AppName    string
	Title      string
	LoggedUser LoggedUser
}

type PageInversores struct {
	AppName    string
	Title      string
	Inversores []Inversor
	LoggedUser LoggedUser
}

type PageModulos struct {
	AppName    string
	Title      string
	Modulos    []Modulo
	LoggedUser LoggedUser
}

type PageParques struct {
	AppName    string
	Title      string
	Parques    []Parque
	LoggedUser LoggedUser
}

type PageProdutos struct {
	AppName    string
	Title      string
	Produtos   []Produto
	LoggedUser LoggedUser
}

type PageRoles struct {
	AppName    string
	Title      string
	Roles      []Role
	Features   []Feature
	LoggedUser LoggedUser
}

type PageStatus struct {
	AppName    string
	Title      string
	Status     []Status
	LoggedUser LoggedUser
}

type PageUsers struct {
	AppName    string
	Title      string
	Users      []User
	Roles      []Role
	LoggedUser LoggedUser
}

type PageUsinas struct {
	AppName    string
	Title      string
	Inversores []Inversor
	Modulos    []Modulo
	Parques    []Parque
	Usinas     []Usina
	LoggedUser LoggedUser
}

type PageWorkflows struct {
	AppName    string
	Title      string
	Actions    []Action
	Roles      []Role
	Workflows  []Workflow
	LoggedUser LoggedUser
}
