package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func getStatus(c echo.Context) error {
	return c.JSON(http.StatusCreated, &httpMsg{Message: "The server is up and running."})
}
