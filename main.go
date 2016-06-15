package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	"github.com/gobott-web/controllers"
	"github.com/gobott-web/mqtt"
)

func main() {
	e := echo.New()

	e.Get("/ping", controllers.Ping("Ping"))

	e.Get("/add_user/:name", controllers.AddUser("AddUser/:Name"))
	e.Get("/get_user/:bucket/:key", controllers.GetUser("GetUser"))
	e.Get("/get_user", controllers.GetUsers("GetUsers"))

	e.Get("/get_all_reports", controllers.GetAllReports("GetAllReports"))

	mqtt.StartMqttClient()

	fmt.Println("Running a Server on localhost:1323")
	e.Run(standard.New(":1323"))
}

func init() {
	//bot.NewBot()
}