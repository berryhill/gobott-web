package controllers

import (
	"github.com/labstack/echo"

	"github.com/gobott-web/store"
	"github.com/gobott-web/models"
)

func GetAllReports(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		//err, reports := store.RetrieveAllFromDb(models.Report{}, []byte("reports"))
		err := store.RetrieveAllFromDb(models.Report{}, []byte("reports"))

		return Respond(c, err, []byte("REPORTS"))
	}
}