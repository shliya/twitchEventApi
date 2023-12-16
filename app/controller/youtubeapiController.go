package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func YoutubeApiController(c echo.Context) error {
	fmt.Println("test")
	return c.JSON(http.StatusAccepted, "test")
}
