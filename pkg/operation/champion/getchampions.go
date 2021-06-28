package champion


import (
	"league_score/pkg/env"
	"league_score/pkg/model"
)

func GetChampions()  ([]*model.Champion,error){
	db :=env.RDB
	champions := []*model.Champion{}
	err := db.Find(&champions).Error
	if err!=nil{
		return nil,err
	}else{
		return champions,nil
	}
}

