package controllers

import (
	"github.com/labstack/echo"

	"github.com/gobott-web/store"
	"github.com/gobott-web/models"
	"github.com/gobott-web/mqtt"
)

func GetAllReports(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		//err, reports := store.RetrieveAllFromDb(models.Report{}, []byte("reports"))
		err := store.RetrieveAllFromDb(models.Report{}, []byte("reports"))

		return Respond(c, err, []byte("REPORTS"))
	}
}

func ResumeReport(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := mqtt.Send([]byte("start_bot"))

		return Respond(c, err, []byte("Resume Report"))
	}
}

func HaltReport(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := mqtt.Send([]byte("stop_bot"))

		return Respond(c, err, []byte("Halt Report"))
	}
}