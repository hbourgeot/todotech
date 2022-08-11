package crud

import "time"

type Clients struct {
	Username string
	Password string
}

type ClientsInfo struct {
	CompleteName string
	Phone        string
	Birth        time.Time
	Country      string
}

type Products struct {
	Cod            int
	Name           string
	Brand          string
	Description    string
	Price          float32
	InventoryCount int
}

type Orders struct {
	ID         int
	ProductCod int
	ClientDNI  int
}