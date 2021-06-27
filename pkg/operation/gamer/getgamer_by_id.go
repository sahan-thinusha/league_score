package gamer

import (
	"league_score/pkg/env"
	"league_score/pkg/model"
)

func GetGamer(id int)  (*model.Gamer,error){
	db :=env.RDB
	gamer := model.Gamer{}
	err := db.First(&gamer,id).Error
	if err!=nil{
		return nil,err
	}else{
		return &gamer,nil
	}
}

