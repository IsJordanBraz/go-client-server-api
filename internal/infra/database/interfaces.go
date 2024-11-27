package database

import "github.com/IsJordanBraz/go-client-server-api/internal/entity"

type CotacaoInterface interface {
	Save(cotacao *entity.Cotacao) error
}
