package main

import (
	"crypto/subtle"
	gorm "github.com/jinzhu/gorm"
	logger "league_score/pkg/logger"
	model "league_score/pkg/model"
	"os"
)

import (
	"league_score/pkg/env"
)
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
import (
	"flag"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
import (
	"league_score/pkg/controller_echo"
)

func readEnvs() {

	if val := os.Getenv(env.REST_PORT); val != "" {
		env.RestPort = val
	}

	if val := os.Getenv(env.E3_URL); val != "" {
		env.E3url = val
	}
	if val := os.Getenv(env.E3_DIALET); val != "" {
		env.E3DIALET = val
	}
}

func main() {

	readEnvs()
	env.LoadEnvs()

	database0, err := gorm.Open("mysql", env.E3user+":"+env.E3pwd+"@tcp("+env.E3host+":"+env.E3port+")/"+env.E3db+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logger.Log.Error(err)
	}
	env.RDB = database0

	model.InitModels(database0)

	CreateDefaultUser()

	RunProxy()
}

func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	run()
}

var (
	endpoint = flag.String("endpoint", "localhost:50051", "Your Description")
)

func run() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	r := e.Group("")
	e.Static("/newsFeed_images","newsFeed_images")
	e.Static("/doctor_images","doctor_images")
//	r := e.Group("/")
	jwtConfig := middleware.JWTConfig{
		Claims:     &env.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}


	r.Use(middleware.JWTWithConfig(jwtConfig))


	u := e.Group("/")
	u.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("league_score")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("league_score@123")) == 1 {
			return true, nil
		}
		return false, nil
	}))
	controller_echo.APIControllerUserBasic(u)

	e.Logger.Fatal(e.Start(":" + env.RestPort))
}


func CreateDefaultUser(){
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
	}
}

