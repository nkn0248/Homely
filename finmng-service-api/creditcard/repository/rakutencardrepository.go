package repository

import (
	"github.com/labstack/echo"
	"github.com/nkn0248/Homely/finmng-service-api/creditcard/domain"
)

type RakutenCardRepository struct {
	client string
}

func NewRakutenCardRepository(client string) domain.RakutenCardRepository {
	return RakutenCardRepository{
		client: client,
	}
}

// Get implements domain.RakutenCardRepository
func (RakutenCardRepository) Get(ctx echo.Context, user string, svc string) (domain.RakutenCard, error) {
	panic("unimplemented")
}
