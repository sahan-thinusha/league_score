package video

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/model"
	op "league_score/pkg/operation/promatches"
)

func GetVideos(c echo.Context)  ([]*model.ProMatches,error) {
	result,err := op.GetProMatches()
	return result, err
}
