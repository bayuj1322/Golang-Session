package session6

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Pegawai struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Umur int     `json:"umur"`
	Gaji float64 `json:"gaji"`
}

var posisi = [3]string{"Staff", "Supervisor", "Manager"}
var pegawais []Pegawai
var pegawaiMap = make(map[int]Pegawai)

func getPosisi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posisi)
}

func getPegawais(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pegawais)
}

func createPegawai(w http.ResponseWriter, r *http.Request) {
	newPegawai := Pegawai{}
	err := json.NewDecoder(r.Body).Decode(&newPegawai)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pegawais = append(pegawais, newPegawai)
	pegawaiMap[newPegawai.ID] = newPegawai
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPegawai)
}

func getPegawaiByID(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID dari URL
	idStr := r.URL.Path[len("/pegawais/"):]

	// Parse ID ke integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
		return
	}

	// Cari pegawai berdasarkan ID
	idpegawai, exists := pegawaiMap[id]
	if !exists {
		http.Error(w, "Pegawai tidak ditemukan", http.StatusNotFound)
		return
	}

	// Kirimkan data pegawai dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(idpegawai)
}

func deletePegawaiByID(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari URL
	idStr := r.URL.Path[len("/pegawais/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
		return
	}

	// Cek apakah pegawai ada
	_, exists := pegawaiMap[id]
	if !exists {
		http.Error(w, "Pegawai tidak ditemukan", http.StatusNotFound)
		return
	}

	// Hapus dari map
	delete(pegawaiMap, id)

	// Hapus dari slice
	for i, p := range pegawais {
		if p.ID == id {
			pegawais = append(pegawais[:i], pegawais[i+1:]...)
			break
		}
	}

	// Re-indexing ID
	for i := range pegawais {
		pegawais[i].ID = i + 1 // Menyesuaikan ID agar berurutan mulai dari 1
	}

	// Perbarui map dengan ID yang baru
	pegawaiMap = make(map[int]Pegawai)
	for _, p := range pegawais {
		pegawaiMap[p.ID] = p // Simpan pegawai dengan ID yang baru
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Pegawai dengan ID %d berhasil dihapus", id)
}

func Webmain() {
	http.HandleFunc("/posisi", getPosisi) // Route untuk mendapatkan posisi
	http.HandleFunc("/pegawais", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getPegawais(w, r)
		} else if r.Method == http.MethodPost {
			createPegawai(w, r)
		} else {
			http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/pegawais/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getPegawaiByID(w, r)
		} else if r.Method == http.MethodDelete {
			deletePegawaiByID(w, r)
		} else {
			http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	// Mulai server di port 8082
	fmt.Println("Server berjalan di http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))

}
