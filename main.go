package main

import "fmt"

type Peserta struct {
	ID            int
	Nama          string
	TanggalDaftar string
	BidangMinat   string
	StatusAktif   bool
}

var peserta []Peserta
var nextID = 10001

func initDataPeserta() {
	peserta = []Peserta{
		{10001, "Andi", "01-01-2025", "Backend Engineering", true},
		{10002, "Budi", "02-01-2025", "Data Science & Artificial Intelligence", true},
		{10003, "Citra", "03-01-2025", "Mobile Development", true},
		{10004, "Dimas", "04-01-2025", "Frontend Engineering", true},
		{10005, "Eka", "05-01-2025", "Backend Engineering", true},
		{10006, "Fajar", "06-01-2025", "Mobile Development", true},
		{10007, "Gita", "07-01-2025", "Data Science & Artificial Intelligence", true},
		{10008, "Hana", "08-01-2025", "Frontend Engineering", true},
		{10009, "Iqbal", "09-01-2025", "Backend Engineering", true},
		{10010, "Joko", "10-01-2025", "Mobile Development", false},
	}

	nextID = 10011
}

func main() {
	initDataPeserta()
	
	var pilihan int

	for {
		fmt.Println("\n===== MENU UTAMA =====")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Ubah Peserta")
		fmt.Println("3. Hapus Peserta")
		fmt.Println("4. Informasi Peserta")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahPeserta()

		case 2:
			ubahPeserta()

		case 3:
			hapusPeserta()

		case 4:
			menuInformasi()

		case 0:
			fmt.Println("Program selesai")
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahPeserta() {
	var p Peserta
	var status int
	var minat int

	fmt.Println("\n===== TAMBAH PESERTA =====")

	p.ID = nextID
	nextID++

	fmt.Println("ID Peserta :", p.ID)

	fmt.Print("Nama : ")
	fmt.Scan(&p.Nama)

	fmt.Print("Tanggal Daftar : ")
	fmt.Scan(&p.TanggalDaftar)

	fmt.Println("Bidang Minat")
	fmt.Println("1. Backend Engineering")
	fmt.Println("2. Mobile Development")
	fmt.Println("3. Data Science & Artificial Intelligence")
	fmt.Println("4. Frontend Engineering")
	fmt.Print("Pilih : ")
	fmt.Scan(&minat)

	switch minat {
	case 1:
		p.BidangMinat = "Backend Engineering"
	case 2:
		p.BidangMinat = "Mobile Development"
	case 3:
		p.BidangMinat = "Data Science & Artificial Intelligence"
	case 4:
		p.BidangMinat = "Frontend Engineering"
	default:
		p.BidangMinat = "Tidak Ada"
	}

	fmt.Println("Status Peserta")
	fmt.Println("1. Aktif")
	fmt.Println("2. Tidak Aktif")
	fmt.Print("Pilih : ")
	fmt.Scan(&status)

	if status == 1 {
		p.StatusAktif = true
	} else {
		p.StatusAktif = false
	}

	peserta = append(peserta, p)

	fmt.Println("Peserta berhasil ditambahkan")
}

func ubahPeserta() {
	var id int
	var status int
	var minat int

	fmt.Print("\nMasukkan ID Peserta : ")
	fmt.Scan(&id)

	for i := 0; i < len(peserta); i++ {

		if peserta[i].ID == id {

			fmt.Print("Nama Baru : ")
			fmt.Scan(&peserta[i].Nama)

			fmt.Print("Tanggal Daftar Baru : ")
			fmt.Scan(&peserta[i].TanggalDaftar)

			fmt.Println("Bidang Minat")
			fmt.Println("1. Backend Engineering")
			fmt.Println("2. Mobile Development")
			fmt.Println("3. Data Science & Artificial Intelligence")
			fmt.Println("4. Frontend Engineering")
			fmt.Print("Pilih : ")
			fmt.Scan(&minat)

			switch minat {
			case 1:
				peserta[i].BidangMinat = "Backend Engineering"
			case 2:
				peserta[i].BidangMinat = "Mobile Development"
			case 3:
				peserta[i].BidangMinat = "Data Science & Artificial Intelligence"
			case 4:
				peserta[i].BidangMinat = "Frontend Engineering"
			}

			fmt.Println("Status Peserta")
			fmt.Println("1. Aktif")
			fmt.Println("2. Tidak Aktif")
			fmt.Print("Pilih : ")
			fmt.Scan(&status)

			if status == 1 {
				peserta[i].StatusAktif = true
			} else {
				peserta[i].StatusAktif = false
			}

			fmt.Println("Data berhasil diubah")
			return
		}
	}

	fmt.Println("ID tidak ditemukan")
}

func hapusPeserta() {
	var id int

	fmt.Print("\nMasukkan ID Peserta : ")
	fmt.Scan(&id)

	for i := 0; i < len(peserta); i++ {

		if peserta[i].ID == id {

			peserta = append(peserta[:i], peserta[i+1:]...)

			fmt.Println("Data berhasil dihapus")
			return
		}
	}

	fmt.Println("ID tidak ditemukan")
}

func menuInformasi() {
	var pilih int

	for {
		fmt.Println("\n===== INFORMASI PESERTA =====")
		fmt.Println("1. Seluruh Peserta")
		fmt.Println("2. Cari Peserta")
		fmt.Println("3. Statistik Peserta")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			menuSeluruhPeserta()

		case 2:
			menuCariPeserta()

		case 3:
			statistikPeserta()

		case 0:
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menuSeluruhPeserta() {
	var pilih int

	fmt.Println("\n===== SELURUH PESERTA =====")
	fmt.Println("1. Urutkan Berdasarkan Nama")
	fmt.Println("2. Urutkan Berdasarkan ID")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	switch pilih {

	case 1:
		insertionSortNama()
		tampilkanPeserta()

	case 2:
		selectionSortID()
		tampilkanPeserta()
	}
}

func menuCariPeserta() {
	var pilih int

	fmt.Println("\n===== CARI PESERTA =====")
	fmt.Println("1. Cari Berdasarkan ID")
	fmt.Println("2. Cari Berdasarkan Nama")
	fmt.Println("3. Cari Berdasarkan Bidang Minat")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	switch pilih {

	case 1:
		cariID()

	case 2:
		cariNama()

	case 3:
		cariBidangMinat()
	}
}

func tampilkanPeserta() {

	if len(peserta) == 0 {
		fmt.Println("Belum ada data")
		return
	}

	fmt.Println("\n===== DATA PESERTA =====")

	for i := 0; i < len(peserta); i++ {

		status := "Tidak Aktif"

		if peserta[i].StatusAktif {
			status = "Aktif"
		}

		fmt.Println("ID            :", peserta[i].ID)
		fmt.Println("Nama          :", peserta[i].Nama)
		fmt.Println("Tanggal Daftar:", peserta[i].TanggalDaftar)
		fmt.Println("Bidang Minat  :", peserta[i].BidangMinat)
		fmt.Println("Status        :", status)
		fmt.Println()
	}
}

func insertionSortNama() {

	for i := 1; i < len(peserta); i++ {

		temp := peserta[i]
		j := i - 1

		for j >= 0 && peserta[j].Nama > temp.Nama {
			peserta[j+1] = peserta[j]
			j--
		}

		peserta[j+1] = temp
	}
}

func selectionSortID() {

	for i := 0; i < len(peserta)-1; i++ {

		min := i

		for j := i + 1; j < len(peserta); j++ {

			if peserta[j].ID < peserta[min].ID {
				min = j
			}
		}

		peserta[i], peserta[min] = peserta[min], peserta[i]
	}
}

func cariID() {
	var id int

	selectionSortID()

	fmt.Print("Masukkan ID : ")
	fmt.Scan(&id)

	kiri := 0
	kanan := len(peserta) - 1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		if peserta[tengah].ID == id {

			fmt.Println("Data ditemukan")
			fmt.Println(peserta[tengah])
			return
		}

		if peserta[tengah].ID < id {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func cariNama() {
	var nama string
	found := false

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&nama)

	for i := 0; i < len(peserta); i++ {

		if peserta[i].Nama == nama {

			fmt.Println("Data ditemukan")
			fmt.Println(peserta[i])
			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func cariBidangMinat() {
	var pilih int
	var bidang string
	found := false

	fmt.Println("\nPilih Bidang Minat")
	fmt.Println("1. Backend Engineering")
	fmt.Println("2. Mobile Development")
	fmt.Println("3. Data Science & Artificial Intelligence")
	fmt.Println("4. Frontend Engineering")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		bidang = "Backend Engineering"
	case 2:
		bidang = "Mobile Development"
	case 3:
		bidang = "Data Science & Artificial Intelligence"
	case 4:
		bidang = "Frontend Engineering"
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	for i := 0; i < len(peserta); i++ {

		if peserta[i].BidangMinat == bidang {

			fmt.Println(peserta[i])
			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func statistikPeserta() {

	backend := 0
	mobile := 0
	dataAI := 0
	frontend := 0
	aktif := 0
	tidakAktif := 0

	for i := 0; i < len(peserta); i++ {

		switch peserta[i].BidangMinat {

		case "Backend Engineering":
			backend++

		case "Mobile Development":
			mobile++

		case "Data Science & Artificial Intelligence":
			dataAI++

		case "Frontend Engineering":
			frontend++
		}

		if peserta[i].StatusAktif {
			aktif++
		} else {
			tidakAktif++
		}
	}

	fmt.Println("\n===== STATISTIK PESERTA =====")
	fmt.Println("Backend Engineering :", backend)
	fmt.Println("Mobile Development :", mobile)
	fmt.Println("Data Science & AI :", dataAI)
	fmt.Println("Frontend Engineering :", frontend)
	fmt.Println("Peserta Aktif :", aktif)
	fmt.Println("Peserta Tidak Aktif :", tidakAktif)
}
