package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(c echo.Context) error {

	d := e.s.DaysLeft()
	s := fmt.Sprintf("Количество дней %d", d)
	err := c.String(http.StatusOK, s)

	if err != nil {
		return err
	}
	return nil
}
