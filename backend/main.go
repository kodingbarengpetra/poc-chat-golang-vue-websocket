package main

import (
	api "github.com/kodingbarengpetra/poc-chat-golang-vue-websocket/backend/generated"
	"github.com/labstack/echo/v4"
)

func main() {
	var handler api.ServerInterface = &Handler{}

	e := echo.New()

	api.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(":1323"))
}
