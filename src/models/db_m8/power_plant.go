package db_m8

import (
	"dbPractise/models"
	"gorm.io/gorm"
	"time"
)

type PowerPlant struct {
	ID             int     `json:"id" gorm:"primarykey"` //厂站ID
	Name           string  `json:"name"`                 //厂站英文名
	CName          string  `json:"c_name"`               //厂站中文名
	Type           string  `json:"type"`                 //卡片类型
	BelongProvince string  `json:"blong_province"`       //所属省网
	BelongCity     string  `json:"blong_city"`           //所属地市
	Voltage        float32 `json:"voltage"`              //电压等级
	AddUser        string  `json:"add_user"`             //添加数据的用户
	//AddTime			string		`json:"add_time"`
	AddTime time.Time `json:"add_time" gorm:"autoUpdateTime"` //添加时间
	Delete  int       `json:"delete"`                         //删除标志位
}

//根据ID在现有power_plant表单中检索是否存在
func ExistPowerPlantByID(id int) (bool, error) {
	var powerplant PowerPlant
	err := models.Db.Select("id").Where("id = ?", id).First(&powerplant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if powerplant.ID > 0 {
		return true, nil
	}

	return false, nil
}

//添加power_plant
func AddPowerPlant(powerplant PowerPlant) error {
	if err := models.Db.Create(&powerplant).Error; err != nil {
		return err
	}
	return nil
}

//根据厂站英文名在power_plant表单中修改元素
func EditPowerPlantByName(name string, data interface{}) error {
	if err := models.Db.Model(&PowerPlant{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

//根据ID检索该厂站信息
func GetPowerPlantByID(id int) (*PowerPlant, error) {
	var powerplant PowerPlant
	err := models.Db.Where("id = ?", id).First(&powerplant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &powerplant, nil
}

//根据英文名检索该厂站信息
func GetPowerPlantByName(name string) (*PowerPlant, error) {
	var powerplant PowerPlant
	err := models.Db.Where("name = ?", name).First(&powerplant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &powerplant, nil
}

//根据电压等级检索所有负荷条件厂站
func GetAllPowerPlantByVoltage(voltage float32) ([]*PowerPlant, error) {
	var powerplant []PowerPlant
	err := models.Db.Where("voltage = ?", voltage).Find(&powerplant).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	temp := make([]*PowerPlant, len(powerplant))
	for aIndex, a := range powerplant {
		temp[aIndex] = &a
	}
	return temp, nil
}
