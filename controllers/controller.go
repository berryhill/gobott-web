package controllers

import (
	"fmt"
	"net/http"
	"log"

	"github.com/labstack/echo"
	"github.com/boltdb/bolt"
)

func init() {
	db, err := bolt.Open("my.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
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

	return func(c echo.Context) error {
		err = nil

		var person struct {
			Name		string
		}
		person.Name = c.Param("name")

		return respond(c, err, person)
	}
}

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