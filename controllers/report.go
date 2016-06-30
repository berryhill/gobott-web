package controllers

import (
	"fmt"
	"log"
	//"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	//"github.com/gobott-web/store"
	"github.com/gobott-web/models"
	"github.com/gobott-web/mqtt"
	"gopkg.in/mgo.v2"
)

func GetReports(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		var session *mgo.Session
		var collection *mgo.Collection
		var err error

		session, err = mgo.Dial("localhost:27017")
		defer session.Close()
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			return err
		}

		collection = session.DB("test").C("reports")
		if err != nil {
			return err
		}

		var results []models.Report
		err = collection.Find(bson.M{"name": "green1"}).Sort("-timestamp").All(&results)
		if err != nil {
			return err
		}
		fmt.Println("Results All: ", results)

		session.Close()

		return err
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
		seconds := c.Param("seconds")
		message := ("timer " + seconds)
		err := mqtt.Send([]byte(string(message)))

		return Respond(c, err, []byte("SetTimer"))
	}
}
