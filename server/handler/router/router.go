package router

import (
	"database/sql"
	"net/http"

	"github.com/t2suna/memo_server/handler"
	"github.com/t2suna/memo_server/handler/middleware"
)

func NewRouter(serviceNameDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/post", middleware.CORS(handler.NewMemoHandler()))
	return mux
}
