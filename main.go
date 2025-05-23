package main

import "fmt"

const MAXDATA = 14

type tidur struct {
	id                                                                         int
	tanggal, kualitas                                                          string
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
		fmt.Println("    APLIKASI MYSLEEP")
		fmt.Println("=========================")
		fmt.Println("1. Tambah Riwayat Tidur")
		fmt.Println("2. Ubah Riwayat Tidur")
		fmt.Println("3. Hapus Riwayat Tidur")
		fmt.Println("4. Pengecekan Durasi")
		fmt.Println("5. Cetak Semua Data")
		fmt.Println("6. Cari Riwayat Tidur")
		fmt.Println("7. Exit")
		fmt.Println("-------------------------")
		fmt.Print("Pilih [1/2/3/4/5/6/7]: ")
		fmt.Scan(&pilih)
		fmt.Println("-------------------------")

		switch pilih {
		case 1:
			tambah_data(&ts, &ns)
		case 2:
			if ns == 0 {
				fmt.Println("Data kosong")
			} else {
				ubah_data(&ts, ns)
			}
		case 3:
			if ns == 0 {
				fmt.Println("Data kosong")
			} else {
				hapus_data(&ts, &ns)
			}
		case 4:
			if ns == 0 {
				fmt.Println("Data kosong")
			} else {
				cek_durasi(&ts, ns)
			}
		case 5:
			cetak_data(ts, ns)
		case 6:
			// cariRiwayat(ts, ns)
			var cari string
			fmt.Print("Masukkan tanggal yang ingin dicari (TTTT-BB-HH): ")
			fmt.Scan(&cari)
			hasil := cari_tanggal(ts, ns, cari)
			if hasil != -1 {
				fmt.Printf("Data ditemukan di indeks ke-%d\n", hasil+1)
				fmt.Println(" ")
			} else {
				fmt.Println("Data tidak ditemukan.")
				fmt.Println(" ")
			}
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
		m[*n].id = *n + 1
		fmt.Print("Tanggal (TTTT-BB-HH): ")
		fmt.Scan(&m[*n].tanggal)
		fmt.Print("Umur (T): ")
		fmt.Scan(&m[*n].umur)
		fmt.Print("Jam Tidur (JJ MM): ")
		fmt.Scan(&m[*n].jamTidur, &m[*n].menitTidur)
		fmt.Print("Jam Bangun (JJ MM): ")
		fmt.Scan(&m[*n].jamBangun, &m[*n].menitBangun)

		*n++
		fmt.Println("Data berhasil ditambahkan!")
		fmt.Println(" ")
	} else {
		fmt.Println("Data sudah penuh")
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
	var cari, x string

	fmt.Println("Masukan tanggal yang ingin di cek")
	fmt.Print("(TTTT-BB-HH): ")
	fmt.Scan(&cari)
	k = cari_tanggal(*m, n, cari)
	fmt.Println("Apakah tidur anda sudah nyenyak/nyaman? (Iya/Tidak)")
	fmt.Scan(&x)

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

		if m[k].umur > 1 && m[k].umur <= 2 {
			if m[k].durasiJam < 11 {
				fmt.Println("Waktu tidur kurang")
			} else if m[k].durasiJam > 14 {
				fmt.Println("Waktu tidur anda berlebih")
			} else {
				fmt.Println("Waktu tidur anda cukup")
			}
		} else if m[k].umur >= 0 && m[k].umur <= 1 {
			if m[k].durasiJam < 13 {
				fmt.Println("Waktu tidur kurang")
			} else if m[k].durasiJam > 16 {
				fmt.Println("Waktu tidur anda berlebih")
			} else {
				fmt.Println("Waktu tidur anda cukup")
			}
		} else if m[k].umur >= 3 && m[k].umur <= 5 {
			if m[k].durasiJam < 11 {
				fmt.Println("Waktu tidur kurang")
			} else if m[k].durasiJam > 13 {
				fmt.Println("Waktu tidur anda berlebih")
			} else {
				fmt.Println("Waktu tidur anda cukup")
			}
		} else if m[k].umur >= 6 && m[k].umur <= 13 {
			if m[k].durasiJam < 9 {
				fmt.Println("Waktu tidur kurang")
			} else if m[k].durasiJam > 11 {
				fmt.Println("Waktu tidur anda berlebih")
			} else {
				fmt.Println("Waktu tidur anda cukup")
			}
		} else if m[k].umur >= 14 && m[k].umur <= 17 {
			if m[k].durasiJam < 8 {
				fmt.Println("Waktu tidur kurang")
			} else if m[k].durasiJam > 10 {
				fmt.Println("Waktu tidur anda berlebih")
			} else {
				fmt.Println("Waktu tidur anda cukup")
			}
		} else if m[k].umur >= 18 && m[k].umur <= 40 {
			if m[k].durasiJam < 7 {
				fmt.Println("Waktu tidur kurang")
			} else if m[k].durasiJam > 9 {
				fmt.Println("Waktu tidur anda berlebih")
			} else {
				fmt.Println("Waktu tidur anda cukup")
			}
		} else if m[k].umur >= 18 && m[k].umur <= 64 {
			if m[k].durasiJam < 7 {
				fmt.Println("Waktu tidur kurang")
			} else if m[k].durasiJam > 9 {
				fmt.Println("Waktu tidur anda berlebih")
			} else {
				fmt.Println("Waktu tidur anda cukup")
			}
		} else if m[k].umur >= 65 {
			if m[k].durasiJam <= 7 {
				fmt.Println("Waktu tidur kurang")
			} else if m[k].durasiJam >= 8 {
				fmt.Println("Waktu tidur anda berlebih")
			} else {
				fmt.Println("Waktu tidur anda cukup")
			}
		} else {
			fmt.Println("Data tidak ditemukan.")
			fmt.Println(" ")
		}
		if x == "Iya" || x == "iya" || x == "IYA" {
			fmt.Println("Pola tidur anda sudah tepat")
		} else if x == "Tidak" || x == "tidak" || x == "TIDAK" {
			fmt.Println("Mungkin ini beberapa saran yang bisa membantu tidur anda lebih baik :")
			fmt.Println("1. Membuat jadwal tidur")
			fmt.Println("2. Menciptakan lingkungan tidur yang nyaman")
			fmt.Println("3. Menghindari bermain ponsel menjelang tidur")
			fmt.Println("4. Melakukan meditasi sebelum tidur")
			fmt.Println("5. Membatasi waktu tidur siang")
			fmt.Println("6. Berolahraga secara rutin")
			fmt.Println("7. Menghentikan kebiasaan makan menjelang tidur")
			fmt.Println("Semoga tips ini membantu anda")
			fmt.Println(" ")
		} else {
			fmt.Println("Masukan pola tidur tidak dapat dibaca")
		}
	}
}

func cetak_data(m tabtidur, n int) {
	fmt.Println("--- DATA RIWAYAT TIDUR ---")
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("%d. %s | Umur: %d | Tidur: %02d:%02d | Bangun: %02d:%02d | Durasi: %d jam %d menit\n", i+1, m[i].tanggal, m[i].umur, m[i].jamTidur, m[i].menitTidur, m[i].jamBangun, m[i].menitBangun, m[i].durasiJam, m[i].durasiMenit)
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
