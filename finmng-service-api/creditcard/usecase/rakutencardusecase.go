package usecase

import (
	"github.com/labstack/echo"
	domain "github.com/nkn0248/Homely/finmng-service-api/creditcard/domain"
)

type RakutenCardUsecase struct {
	rakutencardRepo domain.RakutenCardRepository
}

func NewRakutenCardUsecase(rr domain.RakutenCardRepository) domain.RakutenCardUsecase {
	return RakutenCardUsecase{
		rakutencardRepo: rr,
	}
}

// Get implements domain.RakutenCardUsecase
func (ru RakutenCardUsecase) Get(ctx echo.Context, user string, svc string) (domain.RakutenCard, error) {
	panic("unimplemented")
}
