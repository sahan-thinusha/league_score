package gamer

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/model"
	op "league_score/pkg/operation/gamer"
)

func CreateGamer(c echo.Context) (*model.Gamer,error) {
	gamer := model.Gamer{}
	if err := c.Bind(&gamer); err != nil {
		return nil, err
	}
	result,err := op.GamerRegister(&gamer)
	return result, err
}
