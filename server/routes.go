package server

import (
	"farhan0167/mem-db/database"
	"log/slog"
	"net/http"
)

func addRoutes(mux *http.ServeMux, logger *slog.Logger, db *database.DB) {
	// Table Routes
	mux.Handle("GET /tables", LoggingMiddleware(logger, HandleGetTables(db)))
	mux.Handle("GET /table", LoggingMiddleware(logger, HandleGetTable(db)))
	mux.Handle("POST /tables", LoggingMiddleware(logger, HandleAddTable(db)))

	// Item Routes
	mux.Handle("POST /item", LoggingMiddleware(logger, HandleAddItem(db)))
	mux.Handle("GET /items", LoggingMiddleware(logger, HandleGetItems(db)))
}
