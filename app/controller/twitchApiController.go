package controller

import (
	"io"
	"myapp/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TwitchApiGetredirectUrl(c echo.Context) error {
	twitchBody, _ := io.ReadAll(c.Request().Body)
	response := model.TwitchEventSubModel(twitchBody)
	if response.Challenge != "" {
		return c.String(http.StatusAccepted, response.Challenge)
	} else {
		return c.String(http.StatusAccepted, response.Event.Reward.Title)
	}
}

func TwitchApiGetSubscriptionsList(c echo.Context) error {
	response := model.SubscriptionsList()
	return c.JSON(http.StatusAccepted, response)
}

func TwitchApiCreateSubscriptionsList(c echo.Context) error {
	response := model.SetSubscriptions()
	return c.JSON(http.StatusAccepted, response)
}

func TwitchApiDeleteSubscriptionsList(c echo.Context) error {
	response := model.DeleteUnUseSubscriptions()
	return c.JSON(http.StatusAccepted, response)
}
