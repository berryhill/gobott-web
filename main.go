package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/gobott-web/controllers"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()

	e.Get("/ping", controllers.Ping("Hello"))
	e.Get("/add_person/:name", controllers.AddPerson("Hello"))
	e.Get("/get_person/:bucket/:key", controllers.GetPerson("Hello"))

	fmt.Println("Running a Server on localhost:1323")
	e.Run(standard.New(":1323"))
}

func init() {
	//bot.NewBot()
}
