package db_m9

import (
	"dbPractise/models"
	"gorm.io/gorm"
	"reflect"
)

//func (b *ACBranch) getDyanmicLine() (*db_m2.DynamicLine, error) {
//	var dynamicLine *db_m2.DynamicLine
//	err := models.M9.Model(&ACBranch{}).Where("id = ?", b.ID).Find(b).Error
//	if err != nil {
//		return nil, err
//	}
//	err = models.M9.Model(&ACBranch{}).Preload("id = ?", b.ID).Error
//
//	return dynamicLine, nil
//}

func ExistByID(table string, id int) (bool, error) {
	//var tableStruct interface{}
	//switch table {
	//case "ac_branch":
	//	tableStruct=tableStruct.(ACBranch)
	//case "transformer":
	//	tableStruct=tableStruct.(Transformer)
	//case "dc_branch":
	//	tableStruct=tableStruct.(DCBranch)
	//case "convertor":
	//	tableStruct=tableStruct.(Convertor)
	//}
	var res map[string]interface{}
	// FIXME: 无法指定表单
	//err := models.M9.Table(table).Where("id = ?", id).Take(&res).Error
	err := models.M9.Table(table).First(&res, "id = ?", id).Error
	//err:=models.M9.Raw("SELECT id, type FROM "+table+" WHERE id = ?", id).Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if reflect.TypeOf(res) != reflect.TypeOf(reflect.Map) {
		return true, nil
	}

	return false, nil
}

func GetACBranchByID(id int) (*ACBranch, error) {
	var acb ACBranch
	err := models.M8.Where("id = ? and delete != 1", id).First(&acb).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &acb, nil
}

func GetByID(table string, id int) (interface{}, error) {
	var tableStruct interface{}
	switch table {
	case "ac_branch":
		tableStruct = tableStruct.(ACBranch)
	case "transformer":
		tableStruct = tableStruct.(Transformer)
	case "dc_branch":
		tableStruct = tableStruct.(DCBranch)
	case "convertor":
		tableStruct = tableStruct.(Convertor)
	}

	err := models.M8.Where("id = ? and delete != 1", id).First(&tableStruct).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tableStruct, nil
}
