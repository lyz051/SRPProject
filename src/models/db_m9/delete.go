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
