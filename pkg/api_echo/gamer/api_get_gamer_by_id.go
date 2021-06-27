package gamer

import (
	"github.com/labstack/echo/v4"
	"league_score/pkg/model"
	op "league_score/pkg/operation/gamer"
	"strconv"
)

func GetGamer(c echo.Context)  (*model.Gamer,error) {
	id, _ := strconv.Atoi(c.Param("id"))
	result,err := op.GetGamer(id)
	return result, err

}

