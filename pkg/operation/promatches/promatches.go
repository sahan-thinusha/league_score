package promatches

import (
	"league_score/pkg/env"
	"league_score/pkg/model"
)

func GetProMatches()  ([]*model.ProMatches,error){
	db :=env.RDB
	matches := []*model.ProMatches{}
	err := db.Find(&matches).Error
	if err!=nil{
		return nil,err
	}else{
		return matches,nil
	}
}

