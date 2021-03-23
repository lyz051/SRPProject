package db_m9

import (
	"dbPractise/models"
	"dbPractise/models/db_m2"
)

type ACBranch struct {
	M9
	R           float32 `json:"r"`
	RPU         float32 `json:"rpu"`
	RateA       float32 `json:"rate_a"`
	X           float32 `json:"x"`
	X0          float32 `json:"x0"`
	XPU         float32 `json:"xpu"`
	IRated      float32 `json:"i_rated"`
	ParallelNum int     `json:"parallel_num"`
	G2          float32 `gorm:"column:G_2"`
	B2          float32 `gorm:"column:B_2"`
	Length      float32 `json:"length"`
}

func (this *ACBranch) TableName() string {
	return "ac_branch"
}

func (b ACBranch) Add() error {
	err := models.M9.Create(&b).Error
	if err != nil {
		return err
	}
	return nil
}

func (b ACBranch) GetDyanmicLine() (*db_m2.DynamicLine, error) {
	//err:=models.M9.Model(&ACBranch{}).Association("line_id").Error
	//if err!=nil{
	//	return nil,err
	//}

	var dynamicLine *db_m2.DynamicLine
	err := models.M9.Model(&ACBranch{}).Where("id = ?", b.ID).Association("line_id").
		Find(dynamicLine)
	if err != nil {
		return nil, err
	}

	return dynamicLine, nil
}
