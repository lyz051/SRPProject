package db_m8

import "dbPractise/models"

//添加函数，可判断是否已在数据库表单中存在
func (a *PowerPlant) AddData() {
	flag, _ := ExistPowerPlantByName(a.Name)
	if flag == true {
		return
	} else {
		a.Add()
	}
}

func (a *Bus) AddData() {
	flag, _ := ExistBusByName(a.Name)
	if flag == true {
		return
	} else {
		a.Add()
	}
}

func (a *ACLine) AddData() {
	flag, _ := ExistACLineByName(a.Name)
	if flag == true {
		return
	} else {
		a.Add()
	}
}

func (a *DCLine) AddData() {
	flag, _ := ExistDCLineByName(a.Name)
	if flag == true {
		return
	} else {
		a.Add()
	}
}

func (a *DCSystem) AddData() {
	flag, _ := ExistDCSystemByName(a.Name)
	if flag == true {
		return
	} else {
		a.Add()
	}
}

func (a *Conventor) AddData() {
	flag, _ := ExistConventorByName(a.Name)
	if flag == true {
		return
	} else {
		a.Add()
	}
}

func (a *Trans) AddData() {
	flag, _ := ExistTransByName(a.Name)
	if flag == true {
		return
	} else {
		a.Add()
	}
}

func AddPowerPlant(powerplant PowerPlant) error {
	if err := models.M8.Create(&powerplant).Error; err != nil {
		return err
	}
	return nil
}

func (b *PowerPlant) Add() error {
	if err := models.M8.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func AddBus(bus Bus) error {
	if err := models.M8.Create(&bus).Error; err != nil {
		return err
	}
	return nil
}

func (b *Bus) Add() error {
	if err := models.M8.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func AddACLine(acline ACLine) error {
	if err := models.M8.Create(&acline).Error; err != nil {
		return err
	}
	return nil
}

func (b *ACLine) Add() error {
	if err := models.M8.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func AddDCLine(dcline DCLine) error {
	if err := models.M8.Create(&dcline).Error; err != nil {
		return err
	}
	return nil
}

func (b *DCLine) Add() error {
	if err := models.M8.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func AddConventor(conventor Conventor) error {
	if err := models.M8.Create(&conventor).Error; err != nil {
		return err
	}
	return nil
}

func (b *Conventor) Add() error {
	if err := models.M8.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func AddDCSystem(dcsystem DCSystem) error {
	if err := models.M8.Create(&dcsystem).Error; err != nil {
		return err
	}
	return nil
}

func (b *DCSystem) Add() error {
	if err := models.M8.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func AddTrans(trans Trans) error {
	if err := models.M8.Create(&trans).Error; err != nil {
		return err
	}
	return nil
}

func (b *Trans) Add() error {
	if err := models.M8.Create(&b).Error; err != nil {
		return err
	}
	return nil
}
