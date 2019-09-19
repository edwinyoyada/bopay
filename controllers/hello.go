package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
