package main

import "fmt"

func main() {
	var pilihan, pilihan2 int
	var saldo, jumlah int
	var nama, masukNama string
	var pin, masukPIN int
	var loop bool = true
	var menuATM bool

	fmt.Println("\n=== Selamat Datang Di Bank VINESHYT ===")

	for loop {

		fmt.Println("\n1. Buat kartu")
		fmt.Println("2. Masuk kartu")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih opsi = ")
		fmt.Scan(&pilihan)
		fmt.Scanln()

		if pilihan == 1 {

			var namaBaru string
			var pinBaru int

			fmt.Print("Buat Nama = ")
			fmt.Scan(&namaBaru)
			fmt.Scanln()

			fmt.Print("Masukkan PIN (100000 - 999999) = ")
			fmt.Scan(&pinBaru)
			fmt.Scanln()

			if pinBaru >= 100000 && pinBaru <= 999999 {

				if nama == "" {
					nama = namaBaru
					pin = pinBaru
					saldo = 0
					fmt.Println("\n=== Berhasil Membuat Kartu ===")
				} else if namaBaru == nama {
					fmt.Println("!!! Kartu sudah terdaftar !!!")
				} else {
					fmt.Println("!!! ATM ini hanya mendukung satu kartu !!!")
				}

			} else {
				fmt.Println("PIN harus 6 digit angka dan tidak boleh diawali 0")
			}

		} else if pilihan == 2 {

			fmt.Print("Nama = ")
			fmt.Scan(&masukNama)
			fmt.Scanln()

			fmt.Print("Masukkan PIN = ")
			fmt.Scan(&masukPIN)
			fmt.Scanln()

			if masukNama == nama && masukPIN == pin && nama != "" {

				fmt.Println("\n=== Login Berhasil ===")
				menuATM = true

				for menuATM {

					fmt.Println("\n1. Setor tunai")
					fmt.Println("2. Tarik tunai")
					fmt.Println("3. Cek saldo")
					fmt.Println("4. Ganti PIN")
					fmt.Println("5. Keluar akun")
					fmt.Print("Pilih transaksi = ")
					fmt.Scan(&pilihan2)
					fmt.Scanln()

					if pilihan2 == 1 {
						fmt.Print("Jumlah setor = ")
						fmt.Scan(&jumlah)
						fmt.Scanln()

						if jumlah > 0 {
							saldo += jumlah
							fmt.Println("Saldo =", saldo)
						} else {
							fmt.Println("Jumlah tidak valid")
						}

					} else if pilihan2 == 2 {
						fmt.Print("Jumlah tarik = ")
						fmt.Scan(&jumlah)
						fmt.Scanln()

						if jumlah > 0 && jumlah <= saldo {
							saldo -= jumlah
							fmt.Println("Saldo =", saldo)
						} else {
							fmt.Println("Saldo tidak cukup")
						}

					} else if pilihan2 == 3 {
						fmt.Println("Saldo saat ini =", saldo)

					} else if pilihan2 == 4 {
						var pinLama, pinBaru int

						fmt.Print("Masukkan PIN lama = ")
						fmt.Scan(&pinLama)
						fmt.Scanln()

						if pinLama == pin {

							fmt.Print("Masukkan PIN baru = ")
							fmt.Scan(&pinBaru)
							fmt.Scanln()

							if pinBaru >= 100000 && pinBaru <= 999999 {
								pin = pinBaru
								fmt.Println("PIN berhasil diubah")
							} else {
								fmt.Println("PIN harus 6 digit angka dan tidak boleh diawali 0")
							}

						} else {
							fmt.Println("PIN lama salah")
						}

					} else if pilihan2 == 5 {
						fmt.Println("Berhasil keluar dari akun")
						menuATM = false
					} else {
						fmt.Println("Pilih sesuai menu")
					}
				}

			} else {
				fmt.Println("Nama atau PIN salah")
			}

		} else if pilihan == 3 {
			fmt.Println("Terima kasih telah menggunakan ATM")
			loop = false
		} else {
			fmt.Println("Pilih sesuai menu")
		}
	}
}

