package db_m9

import (
	"dbPractise/models"
	"dbPractise/models/db_m2"
)

func (b *ACBranch) getDyanmicLine() (*db_m2.DynamicLine, error) {
	var dynamicLine *db_m2.DynamicLine
	err := models.M9.Model(&ACBranch{}).Where("id = ?", b.ID).Find(b).Error
	if err != nil {
		return nil, err
	}
	err = models.M9.Model(&ACBranch{}).Preload("id = ?", b.ID).Error

	return dynamicLine, nil
}
