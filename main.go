package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Data Mahasiswa
// 1. GET semua data mahasiswa
// 2. POST data mahasiswa (menambahkan data mahasiswa baru)
// 4. PUT data mahasiswa berdasarkan ID (mengupdate data mahasiswa)
// 5. DELETE data mahasiswa berdasarkan ID (menhapus data mahasiswa)

type Mahasiswa struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

var listMahasiswa = []Mahasiswa{
	{ID: 1, Name: "Budi", Address: "Jakarta", Email: "budi@mail.com"},
	{ID: 2, Name: "Siti", Address: "Bandung", Email: "siti@mail.com"},
	{ID: 3, Name: "Andi", Address: "Surabaya", Email: "andi@mail.com"},
}

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listMahasiswa)
}

func addMahasiswa(w http.ResponseWriter, r *http.Request) {
	var newMahasiswa Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&newMahasiswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMahasiswa.ID = len(listMahasiswa) + 1
	listMahasiswa = append(listMahasiswa, newMahasiswa)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMahasiswa)
}

func updateMahasiswa(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	var id int

	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var uptdatedMahasiswa Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&uptdatedMahasiswa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uptdatedMahasiswa.ID = id

	for i, mahasiswa := range listMahasiswa {
		if mahasiswa.ID == id {
			listMahasiswa[i] = uptdatedMahasiswa
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(uptdatedMahasiswa)
			return
		}
	}
}

func deleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	var id int

	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, mahasiswa := range listMahasiswa {
		if mahasiswa.ID == id {
			listMahasiswa = append(listMahasiswa[:i], listMahasiswa[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Mahasiswa not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getMahasiswa(w, r)
		case http.MethodPost:
			addMahasiswa(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateMahasiswa(w, r)
		case http.MethodDelete:
			deleteMahasiswa(w, r)
		}
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
