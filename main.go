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
	e.Get("/add_person/:name", controllers.AddPerson("AddPerson/:Name"))
	e.Get("/get_person/:bucket/:key", controllers.GetPerson("GetPerson"))
	e.Get("/get_people", controllers.GetPeople("GetPeople"))

	e.Get("/get_all_reports", controllers.GetAllReports("GetAllReports"))

	mqtt.StartMqttClient()

	fmt.Println("Running a Server on localhost:1323")
	e.Run(standard.New(":1323"))
}

func init() {
	//bot.NewBot()
}

