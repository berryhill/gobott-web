package controllers

import (
	"github.com/labstack/echo"

	"github.com/gobott-web/store"
)

func GetAllReports(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := store.RetrieveAllFromDb([]byte("reports"), []byte("report"))

		return Respond(c, err, []byte("FOUND"))
	}
}