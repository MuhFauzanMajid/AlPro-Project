package main

import "fmt"

const NMAX int = 100

type camaba struct {
	nilaiTest     float64
	nama, jurusan string
	status        bool
}

type arrCamaba [NMAX]camaba

func main() {
	var ncamaba int
	var TabCamaba arrCamaba

	HomePage(&TabCamaba, &ncamaba)
}

func HomePage(TabCamaba *arrCamaba, n *int) {
	var pilihan int

	fmt.Println("======================================")
	fmt.Println("   Aplikasi Pendaftaran Mahasiswa     ")
	fmt.Println("   1. Login Calon Mahasiswa           ")
	fmt.Println("   2. Login Admin                     ")
	fmt.Println("======================================")

	fmt.Printf("\nPilihan: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		Menu(*TabCamaba, *n, "camaba")
	} else {
		Menu(*TabCamaba, *n, "Admin")

	}
}

func Menu(TabCamaba arrCamaba, n int, role string) {
	var pilihan int
	var jurusan string

	fmt.Println("============ Menu ==============")
	fmt.Println(" 1. Tampilkan Data Mahasiswa    ")
	if role == "Admin" {
		fmt.Println(" 2. Menambahkan Data Mahasiswa ")
		fmt.Println(" 3. Mengubah Data Mahasiswa    ")
		fmt.Println(" 4. Menghapus Data Mahasiswa   ")
	}
	fmt.Println(" 5. Log Out              ")
	fmt.Println("===============================")

	fmt.Printf("\nPilihan: ")
	fmt.Scan(&pilihan)
	for pilihan < 1 || pilihan > 5 {
		fmt.Println("Pilihan tidak sesuai, mohon masukan pilihan yang ada pada Menu.")
		fmt.Printf("\nPilihan: ")
		fmt.Scan(&pilihan)
	}
	if pilihan == 1 {
		fmt.Println(" 1. Tampilkan Berdasarkan Jurusan  ")
		fmt.Println(" 2. Diurutkan Berdasarkan Nilai Calon Mahasiswa ")
		fmt.Println(" 3. Diurutkan Berdasarkan Jurusan")
		fmt.Println(" 4. Diurutkan Berdasarkan Nama Calon Mahasiswa ")
		fmt.Printf("\nPilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Printf("\nData Mahasiswa untuk Jurusan: ")
			fmt.Scan(&jurusan)
			TampilkanMahasiswa(TabCamaba, jurusan, n)
		} else if pilihan == 2 {
			SortNilai(&TabCamaba, n)
			ReadCamaba(TabCamaba, n)
		} else if pilihan == 3 {
			SortJurusan(&TabCamaba, n)
			ReadCamaba(TabCamaba, n)
		} else {
			SortNama(&TabCamaba, n)
			ReadCamaba(TabCamaba, n)
		}
	} else if pilihan == 2 {
		fmt.Println(" 1. Menambahkan Data Mahasiswa ")
		fmt.Println(" 2. Menambahkan Nilai Mahasiswa ")
		fmt.Printf("\nPilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			CreateCamaba(&TabCamaba, &n)
		} else {
			AddNilai(&TabCamaba, n)
		}
	} else if pilihan == 3 {
		fmt.Println(" 1. Mengubah Data Mahasiswa ")
		fmt.Println(" 2. Mengubah Nilai Mahasiswa ")
		fmt.Printf("\nPilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			EditCamaba(&TabCamaba, n)
		} else {
			EditNilaiCamaba(&TabCamaba, n)
		}
	} else if pilihan == 4 {
		fmt.Println(" 1. Menghapus Data Mahasiswa ")
		fmt.Println(" 2. Menghapus Nilai Mahasiswa ")
		fmt.Printf("\nPilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			DeleteCamaba(&TabCamaba, &n)
		} else {
			DeleteNilai(&TabCamaba, n)
		}
	} else if pilihan == 5 {
		HomePage(&TabCamaba, &n)
	}
	Menu(TabCamaba, n, role)
}

/// CRUD

// Create
func CreateCamaba(TabCamaba *arrCamaba, n *int) {
	fmt.Print("Nama Camaba: ")
	fmt.Scan(&TabCamaba[*n].nama)
	fmt.Printf("Jurusan: ")
	fmt.Scan(&TabCamaba[*n].jurusan)
	*n++
}

func AddNilai(TabCamaba *arrCamaba, n int) {
	var pilihan int

	fmt.Println(" ===== Tambahkan Nilai Tes Camaba ======")
	ReadCamaba(*TabCamaba, n)
	fmt.Printf("\nPilih Camaba: ")
	fmt.Scan(&pilihan)
	fmt.Print("\nNilai Test Camaba: ")
	fmt.Scan(&TabCamaba[pilihan-1].nilaiTest)
}

// Read
func TampilkanMahasiswa(TabCamaba arrCamaba, jurusan string, n int) {
	var j int = 1

	fmt.Println(" ===== Data Camaba ======")
	for i := 0; i < n; i++ {
		if TabCamaba[i].jurusan == jurusan {
			fmt.Println(j, TabCamaba[i].nama)
			j++
		}
	}
	// F.S Menampilkan Data Mahasiswa sesuai dengan jurusan yang dicari
}

func ReadCamaba(TabCamaba arrCamaba, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, TabCamaba[i].nama, "-", TabCamaba[i].jurusan, "-", TabCamaba[i].nilaiTest)
	}
}

// Update
func EditCamaba(TabCamaba *arrCamaba, n int) {
	var pilihan int

	fmt.Println(" ===== Edit Data Camaba ======")

	ReadCamaba(*TabCamaba, n)

	fmt.Printf("\nData Mahasiswa yang mau diubah: ")
	fmt.Scan(&pilihan)
	fmt.Printf("\nNama Mahasiswa: ")
	fmt.Scan(&TabCamaba[pilihan-1].nama)
	fmt.Printf("\nJurusan Mahasiswa: ")
	fmt.Scan(&TabCamaba[pilihan-1].jurusan)
	fmt.Println("Data Mahasiswa Telah Diubah.")
}

func EditNilaiCamaba(TabCamaba *arrCamaba, n int) {
	var pilihan int

	fmt.Println(" ===== Edit Data Camaba ======")
	ReadCamaba(*TabCamaba, n)

	fmt.Printf("\nData Mahasiswa yang mau diubah: ")
	fmt.Scan(&pilihan)
	fmt.Printf("\nNilai Test %v: ", TabCamaba[pilihan-1].nama)
	fmt.Scan(&TabCamaba[pilihan-1].nilaiTest)
	fmt.Println("Data Mahasiswa Telah Diubah.")
}

// Delete
func DeleteCamaba(TabCamaba *arrCamaba, n *int) {
	var pilihan int

	fmt.Println(" ===== Delete Data Camaba ======")
	ReadCamaba(*TabCamaba, *n)
	fmt.Printf("\nData Mahasiswa yang akan dihapus: ")
	fmt.Scan(&pilihan)
	for i := pilihan - 1; i < *n; i++ {
		TabCamaba[i] = TabCamaba[i+1]
	}
	*n--
}

func DeleteNilai(TabCamaba *arrCamaba, n int) {
	var pilihan int

	fmt.Println(" ===== Delete Nilai Tes Camaba ======")
	ReadCamaba(*TabCamaba, n)
	fmt.Printf("\nNilai Tes Mahasiswa yang akan dihapus: ")
	fmt.Scan(&pilihan)
	for i := pilihan - 1; i < n; i++ {
		TabCamaba[i].nilaiTest = 0
	}
}

/// Sorting

func SortNilai(TabCamaba *arrCamaba, n int) {
	var i, j int = 1, 0
	var temp camaba

	for i <= n-1 {
		j = i
		temp = TabCamaba[i]
		for j > 0 && temp.nilaiTest > TabCamaba[j-1].nilaiTest {
			TabCamaba[j] = TabCamaba[j-1]
			j--
		}
		TabCamaba[j] = temp
		i++
	}
	// Menggunakan Insertion Sort
}

func SortJurusan(TabCamaba *arrCamaba, n int) {

	// Menggunakan Selection Sort
	var pass, idx, i int
	var temp camaba

	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if TabCamaba[idx].jurusan[0] > TabCamaba[i].jurusan[0] {
				idx = i
			}
			i++
		}
		temp = TabCamaba[pass-1]
		TabCamaba[pass-1] = TabCamaba[idx]
		TabCamaba[idx] = temp
		pass++
	}
}

func SortNama(TabCamaba *arrCamaba, n int) {

	// Menggunakan Selection Sort
	var pass, idx, i int
	var temp camaba

	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if TabCamaba[idx].nama[0] > TabCamaba[i].nama[0] {
				idx = i
			}
			i++
		}
		temp = TabCamaba[pass-1]
		TabCamaba[pass-1] = TabCamaba[idx]
		TabCamaba[idx] = temp
		pass++
	}
}
