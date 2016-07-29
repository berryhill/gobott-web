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
	"encoding/json"
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

		collection = session.DB("test").C("reports")
		if err != nil {
			return err
		}

		results :=  []models.Report{}
		err = collection.Find(bson.M{"_id": bson.M{"$exists": 1}}).All(&results)
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(results, "", "    ")
		fmt.Println("Number of Reports: ", len(results))
		session.Close()

		return Respond(c, err, json)
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
