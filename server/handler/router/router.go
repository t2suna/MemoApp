package router

import (
	"database/sql"
	"net/http"

	"github.com/t2suna/memo_server/handler"
)

func NewRouter(serviceNameDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/post", handler.NewMemoHandler())
	return mux
}
