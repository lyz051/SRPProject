package db_m9

import (
	"dbPractise/models"
	"log"
	"reflect"
)

func Delete(ins interface{}) error {
	var table string
	switch reflect.TypeOf(ins) {
	case reflect.TypeOf(ACBranch{}):
		ins = ins.(ACBranch)
		table = "ac_branch"
	case reflect.TypeOf(Transformer{}):
		ins = ins.(Transformer)
		table = "transformer"
	case reflect.TypeOf(DCBranch{}):
		ins = ins.(DCBranch)
		table = "dc_branch"
	case reflect.TypeOf(Convertor{}):
		ins = ins.(Convertor)
		table = "convertor"
	}

	ty := reflect.TypeOf(ins)
	log.Println(ty)
	elem := reflect.ValueOf(&ins).Elem()
	val := elem.FieldByName("id")
	err := models.M9.Table(table).Select("delete").Where("id = ?", val.Interface()).Updates(
		map[string]interface{}{
			"delete": 1,
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteACBranchByID(id int) error {
	exist, err := ExistACBranchByID(id)
	if exist && err == nil {
		var acb ACBranch
		acb.ID = id
		if err = Update(acb, map[string]interface{}{
			"delete": 1,
		}); err != nil {
			return err
		}
		return nil
	}
	return err
}

func DeleteTransformerByID(id int) error {
	exist, err := ExistTransformerByID(id)
	if exist && err == nil {
		var trans Transformer
		trans.ID = id
		if err = Update(trans, map[string]interface{}{
			"delete": 1,
		}); err != nil {
			return err
		}
		return nil
	}
	return err
}

func DeleteDCBranchByID(id int) error {
	exist, err := ExistDCBranchByID(id)
	if exist && err == nil {
		var dcb DCBranch
		dcb.ID = id
		if err = Update(dcb, map[string]interface{}{
			"delete": 1,
		}); err != nil {
			return err
		}
		return nil
	}
	return err
}

func DeleteConvertorByID(id int) error {
	exist, err := ExistConvertorByID(id)
	if exist && err == nil {
		var cvt Convertor
		cvt.ID = id
		if err = Update(cvt, map[string]interface{}{
			"delete": 1,
		}); err != nil {
			return err
		}
		return nil
	}
	return err
}
