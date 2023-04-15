package repository

import (
	"context"
	"fmt"

	"github.com/labstack/echo"
	"github.com/nkn0248/Homely/finmng-service-api/creditcard/domain"
)

type CardRepository struct {
	client string
}

func NewCardRepository(client string) domain.CardUsecase {
	return &CardRepository{
		client: client,
	}
}

func (cr *CardRepository) Search(ctx context.Context, issuer string, brand string) (*domain.Card, error) {
	// db access
	fmt.Printf("client(%v) access ¥n", cr.client)
	fmt.Printf("queryparam : issuer = %v, brand = %v ¥n", issuer, brand)
	card := &domain.Card{
		Issuer: "rakuten",
		Brand:  "JCB",
	}
	return card, nil
}

func (cr *CardRepository) Get(ctx echo.Context, user string, svc string) (domain.Card, error) {
	panic("unimplemented")
}
