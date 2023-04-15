package domain

import (
	"github.com/labstack/echo"
)

type RakutenCard struct {
	Id    int64  `json:"id"`
	User  string `json:"user"`
	Money int64  `json:"money"`
}

type RakutenCardUsecase interface {
	Get(ctx echo.Context, user string, svc string) (RakutenCard, error)
}

type RakutenCardRepository interface {
	Get(ctx echo.Context, user string, svc string) (RakutenCard, error)
}
