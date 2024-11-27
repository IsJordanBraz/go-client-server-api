package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/IsJordanBraz/go-client-server-api/internal/entity"
)

type CotacaoDb struct {
	DB *sql.DB
}

func NewCotacaoDb(db *sql.DB) *CotacaoDb {
	return &CotacaoDb{DB: db}
}

func (c *CotacaoDb) Save(cotacao *entity.Cotacao) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()

	stmt, err := c.DB.Prepare("INSERT INTO cotacoes(id, code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Error Sql Prepare: " + err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, cotacao.Id, cotacao.Code, cotacao.Codein, cotacao.Name, cotacao.High, cotacao.Low, cotacao.VarBid, cotacao.PctChange, cotacao.Bid, cotacao.Ask, cotacao.Timestamp, cotacao.CreateDate)
	if err != nil {
		fmt.Println("Error Sql ExecContext: " + err.Error())
		return err
	}
	return nil
}

func (c *CotacaoDb) CreateTable() error {
	_, err := c.DB.Exec("CREATE TABLE IF NOT EXISTS cotacoes (id TEXT, code TEXT, codein TEXT, name TEXT, high TEXT, low TEXT, varBid TEXT, pctChange TEXT, bid TEXT, ask TEXT, timestamp TEXT, create_date TEXT)")
	if err != nil {
		fmt.Println("Error CreateTable: " + err.Error())
		return err
	}
	return nil
}
