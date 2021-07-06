package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	gorm "github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"io/ioutil"
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
	AddVideoData()
	CreateDefaultData()

	RunProxy()
}

func AddChampionData() {
	db := env.RDB
	champions := []*model.Champion{}
	db.Find(&champions)
	if len(champions) == 0 {
		file, err := ioutil.ReadFile("static/champions_data.json")
		if err != nil {
			fmt.Println(err)
		}
		data := []*model.Champion{}
		_ = json.Unmarshal(file, &data)
		for _, champ := range data {
			db.Create(champ)
		}
	}
}
func AddVideoData() {
	db := env.RDB
	pro := []*model.ProMatches{}
	db.Find(&pro)

	if len(pro) == 0 {

		file, err := ioutil.ReadFile("static/promatches.json")
		fmt.Println(string(file))
		if err != nil {
			fmt.Println(err)
		}
		data := []*model.ProMatches{}
		err = json.Unmarshal(file, &data)
		if err != nil {
			fmt.Println(err)
		}
		for _, pro := range data {
			fmt.Println(pro)
			db.Create(pro)
		}
	}
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
	e.Static("/champion_images", "champion_images")
	e.Static("/match", "matches")


	r := e.Group("")

	jwtConfig := middleware.JWTConfig{
		Claims:     &env.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(jwtConfig))
	controller_echo.APIControllerPredict(r)
	controller_echo.APIControllerChampion(r)
	controller_echo.APIControllerGamer(r)
	controller_echo.APIControllerVideos(r)
	controller_echo.APIControllerMatch(r)

	u := e.Group("/")
	u.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("league_score")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("league_score@123")) == 1 {
			return true, nil
		}
		return false, nil
	}))
	controller_echo.APIControllerUserBasic(u)
	controller_echo.APIControllerGamerOpen(u)

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

func CreateDefaultData() {
	db := env.RDB
	pro := []*model.Match{}
	db.Find(&pro)

	if len(pro) == 0 {

		file, err := ioutil.ReadFile("static/match.json")
		fmt.Println(string(file))
		if err != nil {
			fmt.Println(err)
		}
		data := []*model.Match{}
		err = json.Unmarshal(file, &data)
		if err != nil {
			fmt.Println(err)
		}
		for _, pro := range data {
			user := model.User{}
			user.ID = 1
			db.Create(pro)
		}
	}
}

