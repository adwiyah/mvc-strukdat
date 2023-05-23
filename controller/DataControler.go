package controller

import (
	"proy/entity"
	"proy/model"
)

func Daftar(Username string, Saldo int) int {
	return model.TambahNasabah(Username, Saldo)
}

func Masuk(rekening int, username string) *int {
	return model.Login(rekening, username)
}

func AmbilUang(noRek int, tarik int) string {
	return model.TarikSaldo(noRek, tarik)
}

func MasukanUang(noRek int, Saldo int) string {
	return model.TambahSaldo(noRek, Saldo)
}

func CekUang(noRek int) int {
	return model.CekSaldo(noRek)
}

func BlokirRekening(noRek int) string {
	return model.HapusAtm(noRek)
}

func SemuaNasabah() []entity.Data {
	return model.CekUser()
}

func GetNasabah() []entity.Data {
	return model.CekUser()
}
