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

//power_plant检索部分

//根据ID检索
func GetPowerPlantByID(id int) (*PowerPlant, error) {
	var powerplant PowerPlant
	err := models.M8.Where("id = ?", id).First(&powerplant).Error
	if err != nil {
		return nil, err
	}
	return &powerplant, nil
}

//根据英文名检索
func GetPowerPlantByName(name string) (*PowerPlant, error) {
	var powerplant PowerPlant
	err := models.M8.Where("name = ?", name).First(&powerplant).Error
	if err != nil {
		return nil, err
	}
	return &powerplant, nil
}

//根据基准电压检索
func GetAllPowerPlantByVoltage(voltage float32) ([]PowerPlant, error) {
	var powerplant []PowerPlant
	err := models.M8.Where("voltage = ?", voltage).Find(&powerplant).Error
	if err != nil {
		return nil, err
	}
	temp := make([]PowerPlant, len(powerplant))
	for aIndex, a := range powerplant {
		temp[aIndex] = a
	}
	return temp, nil
}

//Bus表单部分

//根据ID检索
func GetBusByID(id int) (*Bus, error) {
	var bus Bus
	err := models.M8.Where("id = ?", id).First(&bus).Error
	if err != nil {
		return nil, err
	}
	return &bus, nil
}

//根据英文名检索
func GetBusByName(name string) (*Bus, error) {
	var bus Bus
	err := models.M8.Where("name = ?", name).First(&bus).Error
	if err != nil {
		return nil, err
	}
	return &bus, nil
}

//根据基准电压检索全部
func GetAllBusByVoltage(voltage float32) ([]Bus, error) {
	var bus []Bus
	err := models.M8.Where("voltage = ?", voltage).Find(&bus).Error
	if err != nil {
		return nil, err
	}
	temp := make([]Bus, len(bus))
	for aIndex, a := range bus {
		temp[aIndex] = a
	}
	return temp, nil
}

//ACLine表单部分

func GetACLineByID(id int) (*ACLine, error) {
	var acline ACLine
	err := models.M8.Where("id = ?", id).First(&acline).Error
	if err != nil {
		return nil, err
	}
	return &acline, nil
}

func GetACLineByName(name string) (*ACLine, error) {
	var acline ACLine
	err := models.M8.Where("name = ?", name).First(&acline).Error
	if err != nil {
		return nil, err
	}
	return &acline, nil
}

//根据A端母线名检索
func GetAllACLineByBusA(bus_a string) ([]ACLine, error) {
	var acline []ACLine
	err := models.M8.Where("bus_a = ?", bus_a).Find(&acline).Error
	if err != nil {
		return nil, err
	}
	temp := make([]ACLine, len(acline))
	for aIndex, a := range acline {
		temp[aIndex] = a
	}
	return temp, nil
}

//根据B端母线名检索
func GetAllACLineByBusB(bus_b string) ([]ACLine, error) {
	var acline []ACLine
	err := models.M8.Where("bus_b = ?", bus_b).Find(&acline).Error
	if err != nil {
		return nil, err
	}
	temp := make([]ACLine, len(acline))
	for aIndex, a := range acline {
		temp[aIndex] = a
	}
	return temp, nil
}

//根据A端厂站名检索
func GetAllACLineSTIDA(stid_a string) ([]ACLine, error) {
	var acline []ACLine
	err := models.M8.Where("stid_a = ?", stid_a).Find(&acline).Error
	if err != nil {
		return nil, err
	}
	temp := make([]ACLine, len(acline))
	for aIndex, a := range acline {
		temp[aIndex] = a
	}
	return temp, nil
}

//根据B端厂站名检索
func GetAllACLineSTIDB(stid_b string) ([]ACLine, error) {
	var acline []ACLine
	err := models.M8.Where("stid_b = ?", stid_b).Find(&acline).Error
	if err != nil {
		return nil, err
	}
	temp := make([]ACLine, len(acline))
	for aIndex, a := range acline {
		temp[aIndex] = a
	}
	return temp, nil
}

//DCLine表单部分

func GetDCLineByID(id int) (*DCLine, error) {
	var dcline DCLine
	err := models.M8.Where("id = ?", id).First(&dcline).Error
	if err != nil {
		return nil, err
	}
	return &dcline, nil
}
func GetDCLineByName(name string) (*DCLine, error) {
	var dcline DCLine
	err := models.M8.Where("name = ?", name).First(&dcline).Error
	if err != nil {
		return nil, err
	}
	return &dcline, nil
}

//DCSystem表单部分

func GetDCSysByID(id int) (*DCSystem, error) {
	var dcsys DCSystem
	err := models.M8.Where("id = ?", id).First(&dcsys).Error
	if err != nil {
		return nil, err
	}
	return &dcsys, nil
}
func GetDCSysByName(name string) (*DCSystem, error) {
	var dcsys DCSystem
	err := models.M8.Where("name = ?", name).First(&dcsys).Error
	if err != nil {
		return nil, err
	}
	return &dcsys, nil
}

//Conventor表单部分

func GetConventorByID(id int) (*Conventor, error) {
	var conventor Conventor
	err := models.M8.Where("id = ?", id).First(&conventor).Error
	if err != nil {
		return nil, err
	}
	return &conventor, nil
}
func GetConventorByName(name string) (*Conventor, error) {
	var conventor Conventor
	err := models.M8.Where("name = ?", name).First(&conventor).Error
	if err != nil {
		return nil, err
	}
	return &conventor, nil
}

//Trans表单部分

func GetTransByID(id int) (*Trans, error) {
	var trans Trans
	err := models.M8.Where("id = ?", id).First(&trans).Error
	if err != nil {
		return nil, err
	}
	return &trans, nil
}
func GetTransByName(name string) (*Trans, error) {
	var trans Trans
	err := models.M8.Where("name = ?", name).First(&trans).Error
	if err != nil {
		return nil, err
	}
	return &trans, nil
}
