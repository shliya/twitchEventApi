package routes

import (
	"myapp/app/controller"

	"github.com/labstack/echo/v4"
)

func YoutubeRoute(e *echo.Echo) {
	e.GET("/youtube/oauth", controller.YoutubeApiController)
}
