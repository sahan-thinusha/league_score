package gamer

import (
	"league_score/pkg/env"
	"league_score/pkg/model"
)

func GetGamers()  ([]*model.Gamer,error){
	db :=env.RDB
	gamers := []*model.Gamer{}
	err := db.Find(&gamers).Error
	if err!=nil{
		return nil,err
	}else{
		return gamers,nil
	}
}

