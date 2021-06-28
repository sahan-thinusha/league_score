package champion

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/model"
	op "league_score/pkg/operation/champion"
)

func GetChampions(c echo.Context)  ([]*model.Champion,error) {
	result,err := op.GetChampions()
	return result, err
}