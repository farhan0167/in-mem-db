package server

import (
	"farhan0167/mem-db/database"
	"farhan0167/mem-db/service"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorMessageResponse struct {
	Error string `json:"error"`
}

func HandleGetTables(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tables := service.GetTables(db)
		encode(w, r, http.StatusOK, tables)
	})
}

func HandleGetTable(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		id := query.Get("id")
		name := query.Get("name")

		if id == "" && name == "" {
			encode(w, r, http.StatusBadRequest, "id or name is required")
			return
		}

		if id != "" && name != "" {
			encode(w, r, http.StatusBadRequest, "id and name cannot be used together")
			return
		}

		params := service.GetTableParams{
			Id:   id,
			Name: name,
		}
		table, err := service.GetTable(db, params)
		if err != nil {
			encode(w, r, http.StatusNotFound, ErrorMessageResponse{Error: err.Error()})
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
		err = service.AddTable(db, table)
		if err != nil {
			encode(w, r, http.StatusBadRequest, ErrorMessageResponse{Error: err.Error()})
			return
		}
		encode(w, r, http.StatusCreated, MessageResponse{Message: "table created"})
	})
}

type AddItemRequest struct {
	Table      string             `json:"name"`
	Attributes service.Attributes `json:"attributes"`
}

func HandleAddItem(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request_body, err := decode[AddItemRequest](r)
		if err != nil {
			log.Println(err)
			encode(w, r, http.StatusInternalServerError, err)
			return
		}
		table, err := db.GetTableByName(request_body.Table)
		if err != nil {
			log.Println(err)
			encode(w, r, http.StatusInternalServerError, err)
			return
		}
		item := database.Item{
			Key: uuid.NewString(),
			Ttl: 3600,
		}
		for k, v := range request_body.Attributes {
			item.Attribute = append(item.Attribute, database.Attribute{
				Name:  k,
				Value: v,
			})
		}
		err = table.AddItem(item)
		if err != nil {
			encode(w, r, http.StatusInternalServerError, err)
			return
		}
		encode(w, r, http.StatusCreated, MessageResponse{Message: "item added successfully"})
	})
}

type GetTableItemsResponse struct {
	Items []service.Attributes `json:"items"`
}

func HandleGetItems(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		table_name := query.Get("table_name")
		if table_name == "" {
			encode(w, r, http.StatusBadRequest, ErrorMessageResponse{Error: "table name is required"})
			return
		}
		items, err := service.GetItems(db, table_name)
		if err != nil {
			encode(w, r, http.StatusNotFound, ErrorMessageResponse{Error: err.Error()})
			return
		}

		encode(w, r, http.StatusOK, items)
	})
}
