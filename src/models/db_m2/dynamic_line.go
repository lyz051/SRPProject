package db_m2

import "dbPractise/models"

type DynamicLine struct {
	M2
	LineName string `gorm:"column:linename"`
	Type     int    `json:"type"`
	StIDA    int    `gorm:"column:stid_a"`
	StIDB    int    `gorm:"column:stid_b"`
	BusIDA   int    `gorm:"column:bus_id_a"`
	BusIDB   int    `gorm:"column:bus_id_b"`
}

func (dl *DynamicLine) Add() error {
	err := models.M2.Create(&dl).Error
	if err != nil {
		return err
	}
	return nil
}
