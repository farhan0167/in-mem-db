package server

import (
	"farhan0167/mem-db/database"
	"net/http"
)

func HandleGetTables(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		empty_tables := []database.Table{}
		tables := db.GetTables()
		if tables == nil {
			encode(w, r, http.StatusNotFound, empty_tables)
			return
		}
		encode(w, r, http.StatusOK, tables)
	})
}

func HandleGetTableById(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		table, err := db.GetTableById(id)
		if err != nil {
			encode(w, r, http.StatusNotFound, database.Table{})
			return
		}
		encode(w, r, http.StatusOK, table)
	})
}

func HandleGetTableByName(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		table, err := db.GetTableByName(name)
		if err != nil {
			encode(w, r, http.StatusNotFound, database.Table{})
			return
		}
		encode(w, r, http.StatusOK, table)
	})
}

func HandleAddTable(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		table, err := decode[database.Table](r)
		if err != nil {
			encode(w, r, http.StatusBadRequest, err)
			return
		}
		err = db.AddTable(table)
		if err != nil {
			encode(w, r, http.StatusBadRequest, err)
			return
		}
		encode(w, r, http.StatusCreated, "Table added successfully")
	})
}
