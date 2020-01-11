package hello

import (
	"github.com/labstack/echo"
	"net/http"
)

func Get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello world!")
}
