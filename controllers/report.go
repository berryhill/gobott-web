package controllers

import (
	"github.com/labstack/echo"

	"github.com/gobott-web/store"
	"github.com/gobott-web/models"
	"github.com/gobott-web/mqtt"
	//"strconv"
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

func SetTimer(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		//seconds, _ := strconv.Atoi(c.Param("seconds"))
		seconds := c.Param("seconds")
		message := ("timer " + seconds)
		err := mqtt.Send([]byte(string(message)))

		return Respond(c, err, []byte("SetTimer"))
	}
}

/*
func SetTimer(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		timer := new(models.Timer)
		seconds, _ := strconv.Atoi(c.Param("seconds"))
		timer.Seconds = seconds

		json, err := json.Marshal(timer)

		if err != nil {
			return Respond(c, err, []byte("ERROR"))
		}

		err = mqtt.Send(json)

		return Respond(c, err, []byte("Halt Report"))
	}
}
*/
