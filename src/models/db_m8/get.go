package db_m8

import (
	"dbPractise/models"
	"gorm.io/gorm"
)

//查找是否存在

//根据name在现有power_plant表单中检索是否存在当前name
func ExistPowerPlantByName(name string) (bool, error) {
	var powerplant PowerPlant
	err := models.M8.Select("name").Where("name = ?", name).First(&powerplant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if powerplant.Name != "" {
		return true, nil
	}

	return false, nil
}

func ExistBusByName(name string) (bool, error) {
	var bus Bus
	err := models.M8.Select("name").Where("name = ?", name).First(&bus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if bus.Name != "" {
		return true, nil
	}

	return false, nil
}

func ExistACLineByName(name string) (bool, error) {
	var acline ACLine
	err := models.M8.Select("name").Where("name = ?", name).First(&acline).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if acline.Name != "" {
		return true, nil
	}

	return false, nil
}

func ExistDCLineByName(name string) (bool, error) {
	var dcline DCLine
	err := models.M8.Select("name").Where("name = ?", name).First(&dcline).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if dcline.Name != "" {
		return true, nil
	}

	return false, nil
}

func ExistDCSystemByName(name string) (bool, error) {
	var dcsystem DCSystem
	err := models.M8.Select("name").Where("name = ?", name).First(&dcsystem).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if dcsystem.Name != "" {
		return true, nil
	}

	return false, nil
}

func ExistConventorByName(name string) (bool, error) {
	var conventor Conventor
	err := models.M8.Select("name").Where("name = ?", name).First(&conventor).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if conventor.Name != "" {
		return true, nil
	}

	return false, nil
}

func ExistTransByName(name string) (bool, error) {
	var trans Trans
	err := models.M8.Select("name").Where("name = ?", name).First(&trans).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if trans.Name != "" {
		return true, nil
	}

	return false, nil
}

//根据ID检索
func GetPowerPlantByID(id int) (*PowerPlant, error) {
	var powerplant PowerPlant
	err := models.M8.Where("id = ?", id).First(&powerplant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &powerplant, nil
}

//根据英文名检索
func GetPowerPlantByName(name string) (*PowerPlant, error) {
	var powerplant PowerPlant
	err := models.M8.Where("name = ?", name).First(&powerplant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &powerplant, nil
}

//根据基准电压检索
func GetAllPowerPlantByVoltage(voltage float32) ([]*PowerPlant, error) {
	var powerplant []PowerPlant
	err := models.M8.Where("voltage = ?", voltage).Find(&powerplant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	temp := make([]*PowerPlant, len(powerplant))
	for aIndex, a := range powerplant {
		temp[aIndex] = &a
	}
	return temp, nil
}
