package db_m2

import "dbPractise/models"

type DynamicLine struct {
	M2
	LineName string `gorm:"column:linename;index"`
	Type     int    `json:"type"`
	StIDA    int    `gorm:"column:stid_a;index"`
	StIDB    int    `gorm:"column:stid_b;index"`
	BusIDA   int    `gorm:"column:bus_id_a;index"`
	BusIDB   int    `gorm:"column:bus_id_b;index"`
}

func (dl *DynamicLine) Add() error {
	err := models.M2.Create(&dl).Error
	if err != nil {
		return err
	}
	return nil
}
