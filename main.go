package main

import (
	"fmt"
	//"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	"github.com/gobott-web/controllers"
	"github.com/gobott-web/mqtt"
)

func main() {
	e := echo.New()

	e.File("/index", "views/index.html")

	e.Get("/ping", controllers.Ping("PING"))

	e.Get("/add_user/:name", controllers.AddUser("AddUser/:Name"))
	e.Get("/get_user/:bucket/:key", controllers.GetUser("GetUser"))
	e.Get("/get_user", controllers.GetUsers("GetUsers"))

	//e.Get("/get_all_reports", controllers.GetAllReports("GetAllReports"))
	e.Get("/reports", controllers.GetReports("GetReports"))

	e.Get("/resume_report", controllers.ResumeReport("Resumer Report"))
	e.Get("/halt_report", controllers.HaltReport("Halt Report"))
	e.Get("/timer/:seconds", controllers.SetTimer("Set Timer"))

	mqtt.StartMqttClient()

	fmt.Println("Running a Server on localhost:1323")
	e.Run(standard.New(":1323"))
}

func init() {
	//bot.NewBot()
}