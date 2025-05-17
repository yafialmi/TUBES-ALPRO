package main

import "fmt"

const MAXDATA = 100

type tidur struct {
	tanggal   string
	jamTidur  string
	jamBangun string
}

type tabtidur [MAXDATA]tidur

var ts tabtidur
var ns int = 0

func main() {
	menu_utama()
}

func menu_utama() {
	var pilih int
	for {
		fmt.Println("=========================")
		fmt.Println("    APLIKASI MYSLEEP")
		fmt.Println("=========================")
		fmt.Println("1. Tambah Riwayat Tidur")
		fmt.Println("2. Ubah Riwayat Tidur")
		fmt.Println("3. Hapus Riwayat Tidur")
		fmt.Println("4. Cetak Semua Data")
		fmt.Println("5. Exit")
		fmt.Println("-------------------------")
		fmt.Println("Pilih [1/2/3/4/5]: ")
		fmt.Scan(&pilih)
		fmt.Println("-------------------------")

		switch pilih {
		case 1:
			tambah_data(&ts, &ns)
		case 2:
			ubah_data(&ts, ns)
		case 3:
			hapus_data(&ts, &ns)
		case 4:
			cetak_data(ts, ns)
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi MySleep!")
			fmt.Println(" ")
		default:
			fmt.Println("Pilihan tidak valid")
			fmt.Println(" ")
		}
		if pilih == 5 {
			break
		}
	}
}

func tambah_data(m *tabtidur, n *int) {
	fmt.Println("Tambah Riwayat Tidur")
	if *n < MAXDATA {
		fmt.Print("Tanggal (TTTT-BB-HH): ")
		fmt.Scan(&m[*n].tanggal)
		fmt.Print("Jam Tidur (JJ:MM): ")
		fmt.Scan(&m[*n].jamTidur)
		fmt.Print("Jam Bangun (JJ:MM): ")
		fmt.Scan(&m[*n].jamBangun)
		*n++
		fmt.Println("Data berhasil ditambahkan!")
		fmt.Println(" ")
	}
}

func ubah_data(m *tabtidur, n int) {
	var cari string
	var k int
	fmt.Println("Masukkan tanggal yang ingin diubah")
	fmt.Print("(TTTT-BB-HH): ")
	fmt.Scan(&cari)
	k = cari_tanggal(*m, n, cari)
	if k != -1 {
		fmt.Print("JamTidurBaru (JJ:MM): ")
		fmt.Scan(&m[k].jamTidur)
		fmt.Print("JamBangunBaru (JJ:MM): ")
		fmt.Scan(&m[k].jamBangun)
		fmt.Println("Data berhasil diubah")
		fmt.Println(" ")
	} else {
		fmt.Println("Data tidak ditemukan")
		fmt.Println(" ")
	}
}

func hapus_data(m *tabtidur, n *int) {
	var cari string
	var k, i int
	fmt.Println("Masukkan tanggal yang ingin dihapus")
	fmt.Print("(TTTT-BB-HH): ")
	fmt.Scan(&cari)
	if k != -1 {
		for i = k; i < *n-1; i++ {
			m[i] = m[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus.")
		fmt.Println(" ")
	} else {
		fmt.Println("Data tidak ditemukan.")
		fmt.Println(" ")
	}
}

func cetak_data(m tabtidur, n int) {
	fmt.Println("--- DATA RIWAYAT TIDUR ---")
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s | Tidur: %s | Bangun: %s\n", i+1, m[i].tanggal, m[i].jamTidur, m[i].jamBangun)
		fmt.Println(" ")
	}
	if n == 0 {
		fmt.Println("Belum ada data.")
		fmt.Println(" ")
	}

}

func cari_tanggal(m tabtidur, n int, cari string) int {
	var i int
	var idx int = -1
	i = 0
	for i < n && idx == -1 {
		if m[i].tanggal == cari {
			idx = i
		}
		i++
	}
	return idx
}
