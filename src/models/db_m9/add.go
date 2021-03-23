package db_m9

import (
	"dbPractise/models"
	"reflect"
)

func (a ACBranch) Add() error {
	err := models.M9.Create(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Transformer) Add() error {
	err := models.M9.Create(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (d DCBranch) Add() error {
	err := models.M9.Create(&d).Error
	if err != nil {
		return err
	}
	return nil
}

func (c Convertor) Add() error {
	err := models.M9.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func Add(ins interface{}) error {
	val := reflect.ValueOf(ins)

	var params []reflect.Value
	retVal := val.MethodByName("Add").Call(params)
	err, ok := retVal[0].Interface().(error)
	if !ok {
		return err
	}
	return nil
}
