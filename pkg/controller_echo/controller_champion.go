package controller_echo

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/api_echo/champion"
	"net/http"
)

func GetChampions(c echo.Context) error {
	result, err := champion.GetChampions(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}


func APIControllerChampion(g *echo.Group) {
	g.GET("api/champion", GetChampions)
}