package db_m8

import (
	"dbPractise/models"
	"fmt"
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

//检索部分，查找全部元素
//power_plant检索部分

//根据ID检索
func GetPowerPlantByID(id int, delete int) (*PowerPlant, error) {
	var powerplant PowerPlant
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("id = ? and `delete` = ? ", id, delete).First(&powerplant).Error
	} else if delete == 2 {
		err = models.M8.Where("id = ?", id).First(&powerplant).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &powerplant, nil
}

//根据英文名检索
func GetPowerPlantByName(name string, delete int) (*PowerPlant, error) {
	var powerplant PowerPlant
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("name = ? and `delete` = ? ", name, delete).First(&powerplant).Error
	} else if delete == 2 {
		err = models.M8.Where("name = ?", name).First(&powerplant).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &powerplant, nil
}

//根据基准电压检索
func GetAllPowerPlantByVoltage(voltage float32, delete int) ([]PowerPlant, error) {
	var powerplant []PowerPlant
	var err error
	if delete == 1 || delete == 0 {
		err = models.M8.Where("voltage = ? AND `delete` = ? ", voltage, delete).Find(&powerplant).Error
	} else if delete == 2 {
		err = models.M8.Where("voltage = ?", voltage).Find(&powerplant).Error
	} else {
		return nil, nil
	}
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
func GetBusByID(id int, delete int) (*Bus, error) {
	var bus Bus
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("id = ? and `delete` = ? ", id, delete).First(&bus).Error
	} else if delete == 2 {
		err = models.M8.Where("id = ?", id).First(&bus).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &bus, nil
}

//根据英文名检索
func GetBusByName(name string, delete int) (*Bus, error) {
	var bus Bus
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("name = ? and `delete` = ? ", name, delete).First(&bus).Error
	} else if delete == 2 {
		err = models.M8.Where("name = ?", name).First(&bus).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &bus, nil
}

//根据基准电压检索全部
func GetAllBusByVoltage(voltage float32, delete int) ([]Bus, error) {
	var bus []Bus
	var err error
	if delete == 1 || delete == 0 {
		err = models.M8.Where("voltage = ? AND `delete` = ? ", voltage, delete).Find(&bus).Error
	} else if delete == 2 {
		err = models.M8.Where("voltage = ?", voltage).Find(&bus).Error
	} else {
		return nil, nil
	}
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

func GetACLineByID(id int, delete int) (*ACLine, error) {
	var acline ACLine
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("id = ? and `delete` = ? ", id, delete).First(&acline).Error
	} else if delete == 2 {
		err = models.M8.Where("id = ?", id).First(&acline).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &acline, nil
}

func GetACLineByName(name string, delete int) (*ACLine, error) {
	var acline ACLine
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("name = ? and `delete` = ? ", name, delete).First(&acline).Error
	} else if delete == 2 {
		err = models.M8.Where("name = ?", name).First(&acline).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &acline, nil
}

//根据A端母线名检索
func GetAllACLineByBusA(bus_a string, delete int) ([]ACLine, error) {
	var acline []ACLine
	var err error
	if delete == 1 || delete == 0 {
		err = models.M8.Where("bus_a = ? AND `delete` = ? ", bus_a, delete).Find(&acline).Error
	} else if delete == 2 {
		err = models.M8.Where("bus_a = ?", bus_a).Find(&acline).Error
	} else {
		return nil, nil
	}
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
func GetAllACLineByBusB(bus_b string, delete int) ([]ACLine, error) {
	var acline []ACLine
	var err error
	if delete == 1 || delete == 0 {
		err = models.M8.Where("bus_a = ? AND `delete` = ? ", bus_b, delete).Find(&acline).Error
	} else if delete == 2 {
		err = models.M8.Where("bus_a = ?", bus_b).Find(&acline).Error
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
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
func GetAllACLineSTIDA(stid_a string, delete int) ([]ACLine, error) {
	var acline []ACLine
	var err error
	if delete == 1 || delete == 0 {
		err = models.M8.Where("stid_a = ? AND `delete` = ? ", stid_a, delete).Find(&acline).Error
	} else if delete == 2 {
		err = models.M8.Where("stid_a = ?", stid_a).Find(&acline).Error
	} else {
		return nil, nil
	}
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
func GetAllACLineSTIDB(stid_b string, delete int) ([]ACLine, error) {
	var acline []ACLine
	var err error
	if delete == 1 || delete == 0 {
		err = models.M8.Where("stid_a = ? AND `delete` = ? ", stid_b, delete).Find(&acline).Error
	} else if delete == 2 {
		err = models.M8.Where("stid_a = ?", stid_b).Find(&acline).Error
	} else {
		return nil, nil
	}
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

func GetDCLineByID(id int, delete int) (*DCLine, error) {
	var dcline DCLine
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("id = ? and `delete` = ? ", id, delete).First(&dcline).Error
	} else if delete == 2 {
		err = models.M8.Where("id = ?", id).First(&dcline).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dcline, nil
}
func GetDCLineByName(name string, delete int) (*DCLine, error) {
	var dcline DCLine
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("name = ? and `delete` = ? ", name, delete).First(&dcline).Error
	} else if delete == 2 {
		err = models.M8.Where("name = ?", name).First(&dcline).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dcline, nil
}

//DCSystem表单部分

func GetDCSysByID(id int, delete int) (*DCSystem, error) {
	var dcsys DCSystem
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("id = ? and `delete` = ? ", id, delete).First(&dcsys).Error
	} else if delete == 2 {
		err = models.M8.Where("id = ?", id).First(&dcsys).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dcsys, nil
}
func GetDCSysByName(name string, delete int) (*DCSystem, error) {
	var dcsys DCSystem
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("name = ? and `delete` = ? ", name, delete).First(&dcsys).Error
	} else if delete == 2 {
		err = models.M8.Where("name = ?", name).First(&dcsys).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dcsys, nil
}

//Conventor表单部分

func GetConventorByID(id int, delete int) (*Conventor, error) {
	var conventor Conventor
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("id = ? and `delete` = ? ", id, delete).First(&conventor).Error
	} else if delete == 2 {
		err = models.M8.Where("id = ?", id).First(&conventor).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &conventor, nil
}
func GetConventorByName(name string, delete int) (*Conventor, error) {
	var conventor Conventor
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("name = ? and `delete` = ? ", name, delete).First(&conventor).Error
	} else if delete == 2 {
		err = models.M8.Where("name = ?", name).First(&conventor).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &conventor, nil
}

//Trans表单部分

func GetTransByID(id int, delete int) (*Trans, error) {
	var trans Trans
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("id = ? and `delete` = ? ", id, delete).First(trans).Error
	} else if delete == 2 {
		err = models.M8.Where("id = ?", id).First(trans).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &trans, nil
}
func GetTransByName(name string, delete int) (*Trans, error) {
	var trans Trans
	var err error
	if delete == 0 || delete == 1 {
		err = models.M8.Where("name = ? and `delete` = ? ", name, delete).First(&trans).Error
	} else if delete == 2 {
		err = models.M8.Where("name = ?", name).First(&trans).Error
	} else {
		fmt.Println("删除标志位输入无效")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &trans, nil
}
