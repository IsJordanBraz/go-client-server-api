package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/IsJordanBraz/go-client-server-api/internal/infra/database"
	"github.com/IsJordanBraz/go-client-server-api/internal/infra/webserver/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		fmt.Println("Error Sql Open: " + err.Error())
		return
	}

	defer db.Close()

	cotacaoDB := database.NewCotacaoDb(db)
	cotacaoDB.CreateTable()
	cotacaoHandler := handlers.NewCotacaoHandler(cotacaoDB)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao", cotacaoHandler.Create)
	http.ListenAndServe(":8080", mux)
}
