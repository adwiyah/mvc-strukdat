package model

import (
	"proy/database"
	"proy/entity"
)

var Rek int = 2334

func TambahNasabah(Username string, Saldo int) int {

	dummy := entity.Nasabah{
		Data: entity.Data{
			Saldo:         Saldo,
			NomerRekening: Rek,
			Username:      Username,
		},
	}

	Data := &database.Nasabah

	if Data.Next == nil {
		Data.Next = &dummy
	} else {
		dummy.Next = Data.Next
		Data.Next = &dummy
	}

	return Rek
}

func Login(noRek int, Username string) *int {
	temp := database.Nasabah.Next

	for temp != nil {
		if Username == temp.Data.Username {
			return &temp.Data.NomerRekening
		}
		temp = temp.Next
	}

	return nil
}

func TarikSaldo(noRek int, tarik int) string {
	temp := database.Nasabah.Next

	for temp != nil {
		if noRek == temp.Data.NomerRekening {
			if tarik > temp.Data.Saldo {
				return "Saldo anda tidak cukup"
			}
			temp.Data.Saldo = temp.Data.Saldo - tarik
			break
		}
		temp = temp.Next
	}

	return "Saldo anda telah di tarik"
}

func TambahSaldo(noRek int, Saldo int) string {
	temp := database.Nasabah.Next

	for temp != nil {
		if noRek == temp.Data.NomerRekening {
			if Saldo == 0 {
				return "Saldo tidak boleh 0"
			}

			temp.Data.Saldo += Saldo
			break
		}
		temp = temp.Next
	}

	return "Saldo anda sudah di tambahkan"
}

func CekSaldo(noRek int) int {
	temp := database.Nasabah.Next

	for temp != nil {
		if temp.Data.NomerRekening == noRek {
			return temp.Data.Saldo
		}
		temp = temp.Next
	}

	return 0
}

func HapusAtm(token int) string {
	temp := &database.Nasabah

	for temp.Next != nil {
		if temp.Data.NomerRekening == token {
			temp.Next = temp.Next.Next
		}
		temp = temp.Next
	}

	return "Akun anda sudah di hapus"
}

func CekUser() []entity.Data {
	temp := database.Nasabah.Next

	var hasil []entity.Data

	for temp != nil {
		hasil = append(hasil, temp.Data)
		temp = temp.Next
	}

	return hasil
}
