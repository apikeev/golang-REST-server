package main

import (
	"Baryon/cmd/baryon/api/task"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	taskserver := task.NewTaskServer()
	task.RegisterHandlers(e, taskserver)

	e.Logger.Fatal(e.Start("localhost:"+os.Getenv("SERVERPORT")))
}
