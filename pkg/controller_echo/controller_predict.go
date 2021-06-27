package controller_echo

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/api_echo/prediction"
	"net/http"
)

func BuildTeam(c echo.Context) error {
	result, err := prediction.PredictTeam(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}



func APIControllerPredict(g *echo.Group) {
	g.GET("api/predict", BuildTeam)
}