package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/t2suna/memo_server/db"
	"github.com/t2suna/memo_server/handler/router"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func realMain() error {
	// config values
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// set up sqlite3
	memoDB, err := db.NewDB(dbPath)
	if err != nil {
		return err
	}
	defer memoDB.Close()
	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := router.NewRouter(memoDB)

	// TODO: サーバーをlistenする
	//http.ListenAndServeTLS(port, "server.crt", "server.key", mux)
	http.ListenAndServe(port, mux)
	return nil
}
