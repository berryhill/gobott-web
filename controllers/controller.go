package controllers

import (
	"fmt"
	"net/http"
	//"log"

	"github.com/labstack/echo"
	//"github.com/boltdb/bolt"

	"github.com/gobott-web/models"
	"github.com/gobott-web/store"
)

/*
func init() {
	db, err := bolt.Open("my.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
*/

func respond(c echo.Context, err error, result interface{}) error {
	var msg string

	if err != nil {
		msg = fmt.Sprint(err)
	}

	statusCode := http.StatusOK

	/*
		switch err {
		case gorm.RecordNotFound:
			statusCode = http.StatusNotFound
		case gorm.InvalidSql, gorm.NoNewAttrs, gorm.NoValidTransaction, gorm.CantStartTransaction:
			statusCode = http.StatusInternalServerError
		}
	*/

	fmt.Println(result)

	return c.JSON(statusCode, map[string]interface{} {
		"result":  result,
		"error":   err != nil,
		"message": msg,
	})
}

func Ping(my interface{}) echo.HandlerFunc {
	var err error

	return func(c echo.Context) error {
		err = nil
		str := "hey"
		return respond(c, err, str)
	}
}

func AddPerson(my interface{}) echo.HandlerFunc {
	var err error
	var json []byte

	return func(c echo.Context) error {
		person := models.NewPerson(c.Param("name"))
		json, err = person.MarshalJson()

		store.AddToDb([]byte("people"), []byte("asdf"), json)

		return respond(c, err, json)
	}
}

/*
func GetPeople() echo.HandlerFunc {
	var err error

	return func(c echo.Context) error {
		people := []models.Person
	}
}
*/

func GetPerson(my interface{}) echo.HandlerFunc {
	var err error

	return func(c echo.Context) error {
		err = store.RetrieveFromDb([]byte(c.Param("bucket")), []byte(c.Param("key")))

		return respond(c, err, []byte("FOUND"))
	}
}
