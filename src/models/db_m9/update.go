package db_m9

import (
	"dbPractise/models"
	"reflect"
)

func (a ACBranch) Update(data map[string]interface{}) error {
	if err := models.M9.Model(&ACBranch{}).Where("id = ?", a.ID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (t Transformer) Update(data map[string]interface{}) error {
	if err := models.M9.Model(&Transformer{}).Where("id = ?", t.ID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (d DCBranch) Update(data map[string]interface{}) error {
	if err := models.M9.Model(&DCBranch{}).Where("id = ?", d.ID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (c Convertor) Update(data map[string]interface{}) error {
	if err := models.M9.Model(&Convertor{}).Where("id = ?", c.ID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func Update(ins interface{}, data map[string]interface{}) error {
	val := reflect.ValueOf(ins)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(data))
	retVal := val.MethodByName("Update").Call(params)
	err, ok := retVal[0].Interface().(error)
	if !ok {
		return err
	}
	return nil
}
