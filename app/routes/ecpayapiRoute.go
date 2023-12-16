package routes

import (
	"myapp/app/controller"

	"github.com/labstack/echo/v4"
)

func ECPayRoute(e *echo.Echo) {
	e.GET("/ecpayapi/:ecpayid", controller.ECPayApiController)
}
