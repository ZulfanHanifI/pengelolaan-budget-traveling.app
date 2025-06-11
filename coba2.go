package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pengeluaran struct {
	ID       int
	Nama     string
	Jumlah   int
	Kategori string
}

var (
	pengeluaranList []Pengeluaran
	idCounter       = 1
	totalBudget     = 0
	reader          = bufio.NewReader(os.Stdin)
)

func main() {
	fmt.Print("Masukkan total budget perjalanan: ")
	input := bacaInput()
	totalBudget, _ = strconv.Atoi(input)

	for {
		tampilkanMenu()
		fmt.Print("Pilih: ")
		input := bacaInput()
		pilihan, _ := strconv.Atoi(input)

		switch pilihan {
		case 1:
			for {
				tambahPengeluaran()
				if !tanyaTambahLagi() {
					break
				}
			}
		case 2:
			lihatSemua()
		case 3:
			ubahPengeluaran()
		case 4:
			hapusPengeluaran()
		case 5:
			cariPengeluaran()
		case 6:
			urutkanPengeluaran()
		case 7:
			tampilkanLaporan()
		case 8:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tampilkanMenu() {
	fmt.Println("\n--- MENU ---")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Lihat Semua Pengeluaran")
	fmt.Println("3. Ubah Pengeluaran")
	fmt.Println("4. Hapus Pengeluaran")
	fmt.Println("5. Cari Pengeluaran (Sequential/Binary)")
	fmt.Println("6. Urutkan Pengeluaran (Selection/Insertion)")
	fmt.Println("7. Laporan Anggaran")
	fmt.Println("8. Keluar")
}

func tanyaTambahLagi() bool {
	fmt.Println("\nTambah Lagi? (Y / N)")
	fmt.Print("Pilih: ")
	input := bacaInput()
	return input == "Y" || input == "y"
}

func bacaInput() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func tambahPengeluaran() {
	fmt.Print("Nama Pengeluaran: ")
	nama := bacaInput()
	fmt.Print("Harga  (Rp): ")
	jumlah, _ := strconv.Atoi(bacaInput())
	fmt.Print("Kategori (transportasi, akomodasi, makanan, hiburan): ")
	kategori := strings.ToLower(bacaInput())

	p := Pengeluaran{idCounter, nama, jumlah, kategori}
	pengeluaranList = append(pengeluaranList, p)
	idCounter++
	fmt.Println("Pengeluaran berhasil ditambahkan!")
}

func lihatSemua() {
	if len(pengeluaranList) == 0 {
		fmt.Println("Belum ada data pengeluaran.")
		return
	}
	for _, p := range pengeluaranList {
		fmt.Printf("ID: %d | Nama: %s | Jumlah: Rp%d | Kategori: %s\n", p.ID, p.Nama, p.Jumlah, p.Kategori)
	}
}

func ubahPengeluaran() {
	fmt.Print("Masukkan ID pengeluaran yang ingin diubah: ")
	id, _ := strconv.Atoi(bacaInput())
	for i, p := range pengeluaranList {
		if p.ID == id {
			fmt.Print("Nama baru: ")
			p.Nama = bacaInput()
			fmt.Print("Jumlah baru: ")
			p.Jumlah, _ = strconv.Atoi(bacaInput())
			fmt.Print("Kategori baru: ")
			p.Kategori = bacaInput()
			pengeluaranList[i] = p
			fmt.Println("Data berhasil diperbarui.")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func hapusPengeluaran() {
	fmt.Print("Masukkan ID pengeluaran yang ingin dihapus: ")
	id, _ := strconv.Atoi(bacaInput())
	for i, p := range pengeluaranList {
		if p.ID == id {
			pengeluaranList = append(pengeluaranList[:i], pengeluaranList[i+1:]...)
			fmt.Println("Data berhasil dihapus.")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func cariPengeluaran() {
	fmt.Println("[1] Sequential Search | [2] Binary Search")
	pilihan := bacaInput()
	fmt.Print("Masukkan nama yang dicari: ")
	nama := strings.ToLower(bacaInput())

	if pilihan == "1" {
		for _, p := range pengeluaranList {
			if strings.ToLower(p.Nama) == nama {
				fmt.Printf("Ditemukan: %+v\n", p)
				return
			}
		}
		fmt.Println("Tidak ditemukan.")
	} else if pilihan == "2" {
		sort.Slice(pengeluaranList, func(i, j int) bool {
			return pengeluaranList[i].Nama < pengeluaranList[j].Nama
		})
		low, high := 0, len(pengeluaranList)-1
		for low <= high {
			mid := (low + high) / 2
			if strings.ToLower(pengeluaranList[mid].Nama) == nama {
				fmt.Printf("Ditemukan: %+v\n", pengeluaranList[mid])
				return
			} else if nama < strings.ToLower(pengeluaranList[mid].Nama) {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		fmt.Println("Tidak ditemukan.")
	}
}

func urutkanPengeluaran() {
	fmt.Println("[1] Berdasarkan Jumlah (Selection Sort)")
	fmt.Println("[2] Berdasarkan Kategori (Insertion Sort)")
	pilihan := bacaInput()

	if pilihan == "1" {
		selectionSortJumlah(pengeluaranList)
	} else if pilihan == "2" {
		insertionSortKategori(pengeluaranList)
	}

	fmt.Println("Pengeluaran setelah diurutkan:")
	lihatSemua()
}

func selectionSortJumlah(arr []Pengeluaran) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].Jumlah < arr[minIdx].Jumlah {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func insertionSortKategori(arr []Pengeluaran) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Kategori > key.Kategori {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func tampilkanLaporan() {
	fmt.Println("\n--- LAPORAN ANGGARAN ---")
	totalPengeluaran := 0
	kategoriMap := make(map[string]int)

	for _, p := range pengeluaranList {
		totalPengeluaran += p.Jumlah
		kategoriMap[p.Kategori] += p.Jumlah
	}

	for kategori, total := range kategoriMap {
		fmt.Printf("Kategori: %s | Total: Rp%d\n", kategori, total)
	}

	fmt.Printf("\nTotal Anggaran: Rp%d\n", totalBudget)
	fmt.Printf("Total Pengeluaran: Rp%d\n", totalPengeluaran)
	fmt.Printf("Sisa Anggaran: Rp%d\n", totalBudget-totalPengeluaran)

	if totalPengeluaran > totalBudget {
		fmt.Println("Peringatan: Pengeluaran melebihi anggaran!")
	}
}
