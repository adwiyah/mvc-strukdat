package view

import (
	"fmt"
	"proy/controller"
	"proy/model"
)

func Start() {
	var menu int
	var username string
	var saldo int
	var tarik int
	var tambah int
	// var nomer int
	var pilihAtm int
	var isi int
	var noRek int

	for {
		fmt.Println("MAIN MENU")
		fmt.Println("1. Daftar akun")
		fmt.Println("2. Masuk")
		fmt.Println("3. Cek semua nasabah")
		fmt.Println("4. Exit")
		fmt.Print("pilih : ")

		fmt.Scan(&menu)

		if menu == 1 {

			fmt.Print("Masukan nama anda : ")
			fmt.Scan(&username)
			fmt.Print("Masukan saldo anda : ")
			fmt.Scan(&saldo)

			hasilDaftar := controller.Daftar(username, saldo)
			model.Rek++

			fmt.Println("Akun anda sudah di buat, ini nomer rekening anda : ", hasilDaftar)

		} else if menu == 2 {

			fmt.Print("Masukan username : ")
			fmt.Scan(&username)
			fmt.Print("Masukan nomer Rekening : ")
			fmt.Scan(&noRek)

			rekening := controller.Masuk(noRek, username)
			fmt.Println("akun anda sudah ditemukan, ini nomor rekening anda : ", noRek)

			if rekening != nil {
				for {
					fmt.Println("1. tarik saldo")
					fmt.Println("2. tambah saldo")
					fmt.Println("3. cek saldo")
					fmt.Println("4. Hapus akun")
					fmt.Println("5. exit")
					fmt.Print("Pilih : ")
					fmt.Scan(&pilihAtm)

					if pilihAtm == 1 {

						fmt.Print("Masukan saldo yang mau anda tarik : ")
						fmt.Scan(&tarik)

						fmt.Println(controller.AmbilUang(*rekening, tarik))

					} else if pilihAtm == 2 {

						fmt.Print("Masukan saldo yang mau anda tambah : ")
						fmt.Scan(&tambah)

						fmt.Println(controller.MasukanUang(*rekening, tambah))
					} else if pilihAtm == 3 {

						isi = controller.CekUang(*rekening)
						fmt.Println("Saldo anda sebesar Rp.", isi)

					} else if pilihAtm == 4 {
						fmt.Println(controller.BlokirRekening(*rekening))
						break
					} else if pilihAtm == 5 {
						break
					} else {
						fmt.Println("Format tidak ditemukan")
					}
				}
			} else {
				fmt.Println("Akun tidak di temukan")
			}
		} else if menu == 3 {
			isi := controller.SemuaNasabah()
			for _, val := range isi {
				fmt.Println(val)
			}
		} else if menu == 4 {
			break
		} else {
			fmt.Println("Format tidak di temukan")
		}
	}
}
