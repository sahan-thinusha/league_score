package match

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"league_score/pkg/env"
	"league_score/pkg/model"
	op "league_score/pkg/operation/match"
)

func GetMatchData(c echo.Context)  ([]*model.Match,error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*env.JwtCustomClaims)

	result,err := op.GetMatches(claims.Id)
	return result, err

}