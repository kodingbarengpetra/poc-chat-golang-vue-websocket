package main

import (
	"github.com/kodingbarengpetra/poc-chat-golang-vue-websocket/backend/generated"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	var handler generated.ServerInterface = &Handler{}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	generated.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(":1323"))
}
