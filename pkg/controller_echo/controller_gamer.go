package controller_echo

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/api_echo/gamer"
	"net/http"
)

func GetGamer(c echo.Context) error {
	result, err := gamer.GetGamer(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetGamers(c echo.Context) error {
	result, err := gamer.GetGamers(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func CreateGamer(c echo.Context) error {
	result, err := gamer.CreateGamer(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}



func APIControllerGamer(g *echo.Group) {
	g.GET("api/gammer/:id", GetGamer)
	g.GET("api/gamers", GetGamers)

}

func APIControllerGamerOpen(g *echo.Group) {

	g.POST("api/gamer", CreateGamer)

}