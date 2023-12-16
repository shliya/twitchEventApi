package controller

import (
	"fmt"
	"myapp/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ECPayApiController(c echo.Context) error {
	ecpayid := c.Param("ecpayid")
	response, err := model.ECPayApiModel(ecpayid)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusAccepted, response)
}
