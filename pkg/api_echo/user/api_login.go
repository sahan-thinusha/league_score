package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"league_score/pkg/env"
	"league_score/pkg/model"
	"net/http"
	"time"
)

type Payload struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
}


func Login(c echo.Context) error {

	loginRequest := new(Payload)
	er1 := c.Bind(loginRequest)

	if er1 != nil {
		return er1
	}
	db := env.RDB
	users := []*model.User{}
	db.Find(&users)
	var targetUser *model.User
	if len(users) == 0 {
		targetUser = &model.User{}
		targetUser.Role = env.ADMIN
		targetUser.Email = "admin"
		targetUser.Password = "admin"
		targetUser.Name = "Admin"
		db.Create(&targetUser)
	} else {
		user := model.User{}
		b := db.Where("email = ? AND password = ?", loginRequest.Email, loginRequest.Password).First(&user).RecordNotFound()
		if !b {
			targetUser = &user
		}
	}

	if targetUser == nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	gamer := model.Gamer{}
	db.Where("user_id = ?",targetUser.ID).First(&gamer)

	claims := env.JwtCustomClaims{Sub: targetUser.Email, Auth: targetUser.Role, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		IssuedAt:  time.Now().Unix(),
	},Id: targetUser.ID,SummonerName: gamer.RiotId,SummonerRegion: gamer.Region}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	c.Response().Header().Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", t))
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Origin, X-Requested-With, Content-Type, Accept")

	if err != nil {
		return err
	}
	targetUser.Token = t
	return c.JSON(http.StatusOK, targetUser)
}
