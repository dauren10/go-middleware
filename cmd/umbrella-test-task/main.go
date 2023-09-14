package main

import (
	"log"
	"myapp/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal("error check")
	}
	err = a.Run()
	if err != nil {
		log.Fatal("error check2 %v", err)
	}
}

// func Handler(c echo.Context) error {

// 	d := time.Date(2025, time.January, 1, 22, 10, 0, 0, time.UTC)

// 	dur := time.Until(d)
// 	s := fmt.Sprintf("Количество дней %d", int64(dur.Hours()/24))
// 	err := c.String(http.StatusOK, s)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
