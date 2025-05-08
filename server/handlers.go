package server

import (
	"farhan0167/mem-db/database"
	"log"
	"net/http"

	"github.com/google/uuid"
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

func HandleGetTable(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		table := &database.Table{}

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

		var err error
		if id != "" {
			table, err = db.GetTableById(id)
			if err != nil {
				encode(w, r, http.StatusNotFound, database.Table{})
				return
			}
		} else if name != "" {
			table, err = db.GetTableByName(name)
			if err != nil {
				encode(w, r, http.StatusNotFound, database.Table{})
				return
			}
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

type Attributes map[string]any

type AddItemRequest struct {
	Table      string     `json:"name"`
	Attributes Attributes `json:"attributes"`
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
		encode(w, r, http.StatusCreated, "Item added successfully")
	})
}

type GetTableItemsResponse struct {
	Items []Attributes `json:"items"`
}

func HandleGetItems(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		empty_items := make([]Attributes, 0)
		query := r.URL.Query()
		table_name := query.Get("table_name")
		if table_name == "" {
			encode(w, r, http.StatusBadRequest, "table name is required")
			return
		}
		table, err := db.GetTableByName(table_name)
		if err != nil {
			encode(w, r, http.StatusNotFound, empty_items)
			return
		}
		items := table.GetItems()
		res := make([]Attributes, len(items))
		for i, item := range items {
			res[i] = make(Attributes)
			for _, attr := range item.Attribute {
				res[i]["key"] = item.Key
				res[i][attr.Name] = attr.Value
			}
		}
		encode(w, r, http.StatusOK, res)
	})
}
