package domain

import (
	"context"

	"github.com/labstack/echo"
)

type Card struct {
	Brand  string `json:"brand"`
	Issuer string `json:"issuer"`
}

type CardUsecase interface {
	Search(ctx context.Context, issuer string, brand string) (*Card, error)
	Get(ctx echo.Context, user string, svc string) (Card, error)
}

type CardRepository interface {
	Search(ctx context.Context, issuer string, brand string) (*Card, error)
	Get(ctx echo.Context, user string, svc string) (Card, error)
}
