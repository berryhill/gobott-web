package controllers

import (
	"fmt"
	"github.com/labstack/echo"

	"github.com/gobott-web/store"
)

func GetReports(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := store.RetrieveAllFromDb([]byte("reports"), []byte("report"))

		fmt.Println("Get Reports")

		return Respond(c, err, []byte("FOUND"))
	}
}