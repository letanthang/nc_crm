package db

import (
	"fmt"
	"github.com/letanthang/nc_crm/model"
)

func GetLevels() (*[]model.Level, error) {
	var levels []model.Level
	if dbc := GetDB().Where("id=?", 8).Find(&levels); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}
	return &levels, nil
}
