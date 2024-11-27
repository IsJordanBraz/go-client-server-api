package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/IsJordanBraz/go-client-server-api/internal/dto"
	"github.com/IsJordanBraz/go-client-server-api/internal/entity"
	"github.com/IsJordanBraz/go-client-server-api/internal/infra/database"
)

type CotacaoHandler struct {
	Db database.CotacaoInterface
}

func NewCotacaoHandler(db database.CotacaoInterface) *CotacaoHandler {
	return &CotacaoHandler{
		Db: db,
	}
}

func (h *CotacaoHandler) Create(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while requesting Economia AwesomeApi"))
		fmt.Println("error while requesting Economia AwesomeApi: " + err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while requesting Economia AwesomeApi"))
		fmt.Println("error while requesting Economia AwesomeApi: " + err.Error())
		return
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while reading Body"))
		fmt.Println("Error while reading Body: " + err.Error())
		return
	}

	var data dto.CotacaoMoedaDto
	err = json.Unmarshal(res, &data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while Unmarshal Body"))
		fmt.Println("Error while Unmarshal Body: " + err.Error())
		return
	}

	cotacao := entity.NewCotacao(&data.USDBRL)

	err = h.Db.Save(cotacao)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Println("error while requesting Economia AwesomeApi" + err.Error())
		return
	}

	var bid dto.CotacaoOutput
	bid.Bid = data.USDBRL.Bid

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bid)
}
