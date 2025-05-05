package server

import (
	"farhan0167/mem-db/database"
	"log/slog"
	"net/http"
)

func addRoutes(mux *http.ServeMux, logger *slog.Logger, db *database.DB) {

	mux.Handle("GET /tables", LoggingMiddleware(logger, HandleGetTables(db)))
	mux.Handle("GET /tables/id/{id}", LoggingMiddleware(logger, HandleGetTableById(db)))
	mux.Handle("GET /tables/name/{name}", LoggingMiddleware(logger, HandleGetTableByName(db)))
	mux.Handle("POST /tables", LoggingMiddleware(logger, HandleAddTable(db)))
}
