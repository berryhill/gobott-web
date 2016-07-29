package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/gobott-web/models"
	"github.com/gobott-web/store"
)

func Respond(c echo.Context, err error, result interface{}) error {
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

		return Respond(c, err, str)
	}
}

func AddUser(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.NewUser(c.Param("name"))
		json, err := user.Save()

		return Respond(c, err, json)
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

func GetUser(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := store.RetrieveFromDb([]byte(c.Param("bucket")), []byte(c.Param("key")))

		return Respond(c, err, []byte("FOUND"))
	}
}

func GetUsers(my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := store.RetrieveAllFromDb([]byte("people"), []byte("asdf"))

		fmt.Println("Get People")

		return Respond(c, err, []byte("FOUND"))
	}
}
