package main

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (h *Handler) PingGet(c echo.Context) error {
	return c.JSON(200, map[string]string{"message": "pong"})
}
