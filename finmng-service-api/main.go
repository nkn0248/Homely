package main

import (
	"github.com/labstack/echo"
	"github.com/nkn0248/Homely/finmng-service-api/creditcard/delivery"
	"github.com/nkn0248/Homely/finmng-service-api/creditcard/repository"
	"github.com/nkn0248/Homely/finmng-service-api/creditcard/usecase"
)

func main() {
	e := echo.New()

	addRakutenCardRoute(e, "hoge")
	e.Logger.Fatal(e.Start(":1323"))
}

func addRakutenCardRoute(e *echo.Echo, str string) {
	client := "psql"
	rr := repository.NewRakutenCardRepository(client)
	ru := usecase.NewRakutenCardUsecase(rr)

	cr := repository.NewCardRepository(client)
	cu := usecase.NewCardUsecase(cr)
	delivery.NewCreditCardHandler(e, cu, ru)
}
