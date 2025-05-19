package main

import "fmt"

const MAXDATA = 100

type tidur struct {
	tanggal                                                                    string
	umur, jamTidur, menitTidur, jamBangun, menitBangun, durasiJam, durasiMenit int
}

type tabtidur [MAXDATA]tidur

func main() {
	menu_utama()
}

func menu_utama() {
	var ts tabtidur
	var ns int = 0
	var pilih int
	for {
		fmt.Println("=========================")
		fmt.Println("    APLIKASI MYSLEEP Azmi")
		fmt.Println("=========================")
		fmt.Println("1. Tambah Riwayat Tidur")
		fmt.Println("2. Ubah Riwayat Tidur")
		fmt.Println("3. Hapus Riwayat Tidur")
		fmt.Println("4. Pengecekan Durasi")
		fmt.Println("5. Cetak Semua Data")
		fmt.Println("6. Cari Riwayat Tidur")
		fmt.Println("7. Exit")
		fmt.Println("-------------------------")
		fmt.Println("Pilih [1/2/3/4/5/6/7]: ")
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
			cek_durasi(&ts, ns)
		case 5:
			cetak_data(ts, ns)
		case 6:
			// cariRiwayat(ts, ns)
		case 7:
			fmt.Println("Terima kasih telah menggunakan aplikasi MySleep!")
			fmt.Println(" ")
		default:
			fmt.Println("Pilihan tidak valid")
			fmt.Println(" ")
		}
		if pilih == 7 {
			break
		}
	}
}

func tambah_data(m *tabtidur, n *int) {
	fmt.Println("Tambah Riwayat Tidur")
	if *n < MAXDATA {
		fmt.Print("Tanggal (TTTT-BB-HH): ")
		fmt.Scan(&m[*n].tanggal)
		fmt.Print("Umur (T): ")
		fmt.Scan(&m[*n].umur)
		fmt.Print("Jam Tidur (JJ:MM): ")
		fmt.Scan(&m[*n].jamTidur, &m[*n].menitTidur)
		fmt.Print("Jam Bangun (JJ:MM): ")
		fmt.Scan(&m[*n].jamBangun, &m[*n].menitBangun)
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
		fmt.Scan(&m[k].jamTidur, &m[k].menitTidur)
		fmt.Print("JamBangunBaru (JJ:MM): ")
		fmt.Scan(&m[k].jamBangun, &m[k].menitBangun)
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
	k = cari_tanggal(*m, *n, cari)
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

func cek_durasi(m *tabtidur, n int) {
	var k int
	var cari string

	fmt.Println("Masukan tanggal yang ingin di cek")
	fmt.Print("(TTTT-BB-HH): ")
	fmt.Scan(&cari)
	k = cari_tanggal(*m, n, cari)

	if k != -1 {
		*&m[k].durasiJam = 23 - m[k].jamTidur + m[k].jamBangun
		*&m[k].durasiMenit = 60 - m[k].menitTidur + m[k].menitBangun
		if m[k].durasiJam >= 24 {
			m[k].durasiJam = m[k].durasiJam - 24
		}
		if m[k].durasiMenit >= 60 {
			m[k].durasiMenit = m[k].durasiMenit - 60
			m[k].durasiJam++
		}
		fmt.Printf("Berdasarkan pengecekan, durasi tidur anda adalah %d jam, %d menit\n", m[k].durasiJam, m[k].durasiMenit)
		if m[k].durasiJam < 7 {
			fmt.Println("Waktu tidur anda kurang cukup")
		} else if m[k].durasiJam > 11 {
			fmt.Println("Waktu tidur anda lebih dari cukup")
		} else {
			fmt.Println("Waktu tidur anda cukup")
		}
	} else {
		fmt.Println("Data tidak ditemukan.")
		fmt.Println(" ")
	}
}

func cetak_data(m tabtidur, n int) {
	fmt.Println("--- DATA RIWAYAT TIDUR ---")
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s | Umur: %d | Tidur: %d:%d | Bangun: %d:%d\n", i+1, m[i].tanggal, m[i].umur, m[i].jamTidur, m[i].menitTidur, m[i].jamBangun, m[i].menitBangun)
		fmt.Println(" ")
	}
	if n == 0 {
		fmt.Println("Belum ada data.")
		fmt.Println(" ")
	}

}

//TODO: UPDATE CARI TANGGAL JADI SEQUENTIAL SEARCH
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

//TODO: SORT MENAIK MENGGUNAKAN INSERTION SORT
func sortMenaik() {}

//TODO: SORT MENURUN MENGGUNAKAN SORT BIASA
func sortMenurun() {}
