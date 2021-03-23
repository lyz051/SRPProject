package db_m9

import "dbPractise/models"

func (b *ACBranch) Add() error {
	err := models.M9.Create(b).Error
	if err != nil {
		return err
	}
	return nil
}
