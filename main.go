package main

import (
	"crypto/subtle"
	gorm "github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	logger "league_score/pkg/logger"
	model "league_score/pkg/model"
	"league_score/pkg/prediction/data"
	"league_score/pkg/prediction/recommender"
	"os"
	"strings"
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

	train, header, err := data.ReadData(*trainingData)
	if err != nil {
		log.Fatalf("Error reading training set: %v", err)
	}
	for i, val := range header {
		name := strings.ToLower(val)
		env.ChampionToIndex[name] = i
		env.IndexToChampion[i] = name
	}

	recommender.ENGINE = recommender.NewNeighborhoodBasedRecommender(train, 5)

	model.InitModels(database0)

	CreateDefaultUser()
	AddChampionData()


	RunProxy()
}

func AddChampionData() {

}

func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	run()
}



var (
	trainingData = pflag.String("trainingset", "static/winning_teams.csv", "path to training dataset")
)


func run() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	r := e.Group("")
	e.Static("/champion_images","champion_images")
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
	controller_echo.APIControllerPredict(u)

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

