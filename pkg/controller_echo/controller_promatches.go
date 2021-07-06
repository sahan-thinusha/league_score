package controller_echo

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/api_echo/video"
	"net/http"
)

func GetVideos(c echo.Context) error {
	result, err := video.GetVideos(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func APIControllerVideos(g *echo.Group) {
	g.GET("api/video", GetVideos)
}