package usecase

import (
	"context"

	"github.com/labstack/echo"
	domain "github.com/nkn0248/Homely/finmng-service-api/creditcard/domain"
)

type CardUsecase struct {
	cardRepo domain.CardRepository
}

func NewCardUsecase(cr domain.CardRepository) domain.CardUsecase {
	return &CardUsecase{
		cardRepo: cr,
	}
}

// Get implements domain.CardUsecase
func (cu *CardUsecase) Search(ctx context.Context, issuer string, brand string) (*domain.Card, error) {
	// business logic
	res, err := cu.cardRepo.Search(ctx, issuer, brand)
	if err != nil {
		return nil, err
	}
	return res, nil

}

// Get implements domain.CardUsecase
func (*CardUsecase) Get(ctx echo.Context, user string, svc string) (domain.Card, error) {
	panic("unimplemented")
}
