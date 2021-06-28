package prediction

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"league_score/pkg/env"
	"league_score/pkg/model"
	_ "league_score/pkg/operation/gamer"
	_ "league_score/pkg/prediction/data"
	"league_score/pkg/prediction/recommender"
	"league_score/pkg/prediction/vector"
)

func PredictTeam(c echo.Context)  ([] model.Team,error) {
	ch1 := c.QueryParam("ch1")
	ch2 := c.QueryParam("ch2")
	ch3 := c.QueryParam("ch3")
	ch4 := c.QueryParam("ch4")

	pinput := model.PredictionInput{}
	champions := []string{}

	fmt.Println(ch1)
	if ch1!=""{
		champions = append(champions,ch1)
	}

	if ch2!=""{
		champions = append(champions,ch2)
	}

	if ch3!=""{
		champions = append(champions,ch3)
	}

	if ch4!=""{
		champions = append(champions,ch4)
	}

	pinput.Champions = champions

	input := make([]float64, len(env.ChampionToIndex))

	if len(pinput.Champions) == 0 || len(pinput.Champions) > 4 {
		log.Error("Invalid number of champions. Please provide 4 or less")
		return nil,errors.New("Invalid number of champions. Please provide 4 or less")
	}

	for _, val := range pinput.Champions {
		item, ok := env.ChampionToIndex[val]
		if !ok {
			log.Warningf("Unknown champion: %v", val)
			return nil,errors.New("Unknown champion")
		}

		input[item] = 1

	}


	recommendations, err := recommender.ENGINE.Recommend(input, vector.Pearson, pinput.Intercept, pinput.Shuffle, pinput.Serendipity)
	if err != nil {
		log.Errorf("Error predicting recommendation: %v", err)
		return nil,errors.New("Error predicting recommendation")

	}

	allRecommendations := [][]string{}
	for _, val := range recommendations {
		recommendedItem := []string{}
		for i, isRecommended := range val.GetRecommendation() {
			if isRecommended != 1 {
				continue
			}
			champ := env.IndexToChampion[int(i)]
			recommendedItem = append(recommendedItem, champ)
		}
		allRecommendations = append(allRecommendations, recommendedItem)
	}


	teams := []model.Team{}
	for _,rec := range allRecommendations{
		champions := []*model.Champion{}
		for _,data := range rec{
			champ := getChampionData(data)
			champions = append(champions,champ)
		}
		team := model.Team{}
		team.Champion = champions
		teams = append(teams,team)
	}


	return teams,err
}


func getChampionData(index string) *model.Champion{
	db :=env.RDB
	champion := model.Champion{}
	db.LogMode(true)
	err := db.Where("ch_id = ?",index).First(&champion).Error
	if err!=nil{
		return nil
	}else{
		return &champion
	}
}