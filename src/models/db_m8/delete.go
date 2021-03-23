package db_m8

import (
	"dbPractise/models"
)

//修改现有拓扑删除标志位为1
func DeletePowerPlant(name string) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistPowerPlantByName(name); flag == true {
		EditPowerPlantByName(name, a)
	} else {
		return
	}
}

func DeleteBusPlant(name string) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistBusByName(name); flag == true {
		EditBusByName(name, a)
	} else {
		return
	}
}

func DeleteACLinePlant(name string) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistACLineByName(name); flag == true {
		EditACLineByName(name, a)
	} else {
		return
	}
}

func DeleteDCLinePlant(name string) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistDCLineByName(name); flag == true {
		EditDCLineByName(name, a)
	} else {
		return
	}
}

func DeleteDCSystemPlant(name string) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistDCSystemByName(name); flag == true {
		EditDCSystemByName(name, a)
	} else {
		return
	}
}

func DeleteConventorPlant(name string) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistConventorByName(name); flag == true {
		EditConventorByName(name, a)
	} else {
		return
	}
}

func DeleteTransPlant(name string) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistTransByName(name); flag == true {
		EditTransByName(name, a)
	} else {
		return
	}
}

//修改元素部分
//根据厂站英文名在表单中修改元素
func EditPowerPlantByName(name string, data interface{}) error {
	if err := models.M8.Model(&PowerPlant{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func EditBusByName(name string, data interface{}) error {
	if err := models.M8.Model(&Bus{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func EditACLineByName(name string, data interface{}) error {
	if err := models.M8.Model(&ACLine{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func EditDCLineByName(name string, data interface{}) error {
	if err := models.M8.Model(&DCLine{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func EditDCSystemByName(name string, data interface{}) error {
	if err := models.M8.Model(&DCSystem{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func EditConventorByName(name string, data interface{}) error {
	if err := models.M8.Model(&Conventor{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func EditTransByName(name string, data interface{}) error {
	if err := models.M8.Model(&Trans{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
