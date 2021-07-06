package match

import (
	"league_score/pkg/env"
	"league_score/pkg/model"
)

func GetMatches(id int64)  ([]*model.Match,error){
	db :=env.RDB
	matches := []*model.Match{}
	err := db.Model(model.Match{}).Where("uid = ?",id).Scan(&matches).Error
	if err!=nil{
		return nil,err
	}else{
		return matches,nil
	}
}
