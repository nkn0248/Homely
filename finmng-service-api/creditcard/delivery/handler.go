package delivery

import (
	"github.com/labstack/echo"
	domain "github.com/nkn0248/Homely/finmng-service-api/creditcard/domain"
)

// FinMngHandlerはUsecaseのInterfaceを保持
type CreditCardHandler struct {
	CardUsecase        domain.CardUsecase
	RakutenCardUsecase domain.RakutenCardUsecase
}

func NewCreditCardHandler(e *echo.Echo, cu domain.CardUsecase, ru domain.RakutenCardUsecase) {
	handler := &CreditCardHandler{
		CardUsecase:        cu,
		RakutenCardUsecase: ru,
	}

	e.GET("/api/v1/card", handler.Search)
}

func (h CreditCardHandler) Search(c echo.Context) error {
	issuer := c.QueryParam("issuer")
	brand := c.QueryParam("brand")
	ctx := c.Request().Context()

	res, err := h.CardUsecase.Search(ctx, issuer, brand)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, res)
}

func (h CreditCardHandler) Get(c echo.Context) error {
	res, err := h.CardUsecase.Get(c, "nakano", "rakuten")
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, res)
}
