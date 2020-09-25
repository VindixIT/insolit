package models

import (
	"github.com/lib/pq"
	"time"
)

var AppName = "InSolIT"

type Role struct {
	Order    int
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Selected bool
	Features []Feature
}

type Feature struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
}

type Workflow struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

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

type Status struct {
	Order      int
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Stereotype string `json:"stereotype"`
}

type Beer struct {
	Order int
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Qtd   int     `json:"qtd"`
	Price float64 `json:"price"`
}

type Item struct {
	Order     int
	Id        int64   `json:"id"`
	IdOrder   int64   `json:"orderId"`
	BeerId    int64   `json:"beerId"`
	BeerName  string  `json:"beerName"`
	Qtt       float64 `json:"qtd"`
	Price     float64 `json:"price"`
	ItemValue float64 `json:"value"`
}

type Order struct {
	Order            int
	Id               int64     `json:"id"`
	UserId           int64     `json:"clientId"`
	UserName         string    `json:"clientName`
	OrderedAt        time.Time `json:"orderedAt`
	TakeOutAt        time.Time `json:"endAt"`
	COrderedDateTime string    `json:"corderedDateTime`
	CTakeOutDateTime string    `json:"ctakeOutDateTime`
	COrderedAt       string    `json:"corderedAt`
	CTakeOutAt       string    `json:"ctakeOutAt`
}

type NullTime struct {
	pq.NullTime
}

// NÃO SEI SE VOU USAR AINDA
type LoggedUser struct {
	User          User
	HasPermission func(string) bool
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

type Measure struct {
	Order int
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

type PageMeasures struct {
	AppName    string
	Title      string
	Measures   []Measure
	LoggedUser LoggedUser
}

type PageOrders struct {
	AppName    string
	Title      string
	UserId     int
	Orders     []Order
	Beers      []Beer
	Users      []User
	LoggedUser LoggedUser
}

type PageUsers struct {
	AppName    string
	Title      string
	Users      []User
	Roles      []Role
	LoggedUser LoggedUser
}

type PageBeers struct {
	AppName    string
	Title      string
	Beers      []Beer
	LoggedUser LoggedUser
}

type PageRoles struct {
	AppName    string
	Title      string
	Roles      []Role
	Features   []Feature
	LoggedUser LoggedUser
}

type PageFeatures struct {
	AppName    string
	Title      string
	Features   []Feature
	LoggedUser LoggedUser
}

type PageStatus struct {
	AppName    string
	Title      string
	Status     []Status
	LoggedUser LoggedUser
}

type PageAction struct {
	AppName    string
	Title      string
	Statuss    []Status
	Actions    []Action
	LoggedUser LoggedUser
}

type PageWorkflow struct {
	AppName    string
	Title      string
	Actions    []Action
	Roles      []Role
	Workflows  []Workflow
	LoggedUser LoggedUser
}
