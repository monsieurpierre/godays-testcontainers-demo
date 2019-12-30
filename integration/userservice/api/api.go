package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type User struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

//noinspection ALL
func PostUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		if res := db.QueryRow(`INSERT INTO users (name) VALUES ($1) RETURNING *`, user.Name); res.Scan(&user.ID, &user.Name) != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

//noinspection ALL
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		var user User
		if res := db.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id); res.Scan(&user.ID, &user.Name) != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}