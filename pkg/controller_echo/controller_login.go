package controller_echo

import (
	"league_score/pkg/api_echo/user"
)

import (
	"github.com/labstack/echo/v4"
)



func APIControllerUserBasic(g *echo.Group) {
	g.POST("api/login", user.Login)

}