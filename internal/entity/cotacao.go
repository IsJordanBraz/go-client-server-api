package entity

import (
	"github.com/IsJordanBraz/go-client-server-api/internal/dto"
	"github.com/google/uuid"
)

type Cotacao struct {
	Id         string `json:"id"`
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func NewCotacao(dto *dto.CotacaoDolarDto) *Cotacao {
	return &Cotacao{
		Id:         uuid.New().String(),
		Code:       dto.Code,
		Codein:     dto.Codein,
		Name:       dto.Name,
		High:       dto.High,
		Low:        dto.Low,
		VarBid:     dto.VarBid,
		PctChange:  dto.PctChange,
		Bid:        dto.Bid,
		Ask:        dto.Ask,
		Timestamp:  dto.Timestamp,
		CreateDate: dto.CreateDate,
	}
}
