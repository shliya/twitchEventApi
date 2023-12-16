package routes

import (
	"myapp/app/controller"

	"github.com/labstack/echo/v4"
)

func TwitchApiRoute(e *echo.Echo) {
	e.GET("/twitch/eventsub/get", controller.TwitchApiGetSubscriptionsList)
	e.POST("/twitch/eventsub/callback", controller.TwitchApiGetredirectUrl)
	e.POST("/twitch/eventsub/post", controller.TwitchApiCreateSubscriptionsList)
	e.DELETE("/twitch/eventsub/delete", controller.TwitchApiDeleteSubscriptionsList)
}
