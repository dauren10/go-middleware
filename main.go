package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Server is started")
	s := echo.New()

	s.GET("/status", Handler)

	err := s.Start(":1325")
	if err != nil {
		log.Fatal("Error here")
	}
}
func Handler(c echo.Context) error {

	d := time.Date(2025, time.January, 1, 22, 10, 0, 0, time.UTC)

	dur := time.Until(d)
	s := fmt.Sprintf("Количество дней %d", int64(dur.Hours()/24))
	err := c.String(http.StatusOK, s)

	if err != nil {
		return err
	}
	return nil
}
