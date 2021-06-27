package gamer

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/model"
	op "league_score/pkg/operation/gamer"

)

func GetGamers(c echo.Context)  ([]*model.Gamer,error) {
	result,err := op.GetGamers()
	return result, err
}

