package gamer

import (
	"league_score/pkg/env"
	"league_score/pkg/model"
)

func GamerRegister(gamer *model.Gamer)  (*model.Gamer,error){
	db :=env.RDB
	err := db.Create(gamer).Error
	if err!=nil{
		return nil,err
	}else{
		return gamer,nil
	}
}
