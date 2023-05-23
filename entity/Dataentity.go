package entity

type Data struct {
	Saldo         int
	NomerRekening int
	Username      string
}

type Nasabah struct {
	Data Data
	Next *Nasabah
}
