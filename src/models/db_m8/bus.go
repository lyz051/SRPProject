package db_m8

import (
	"dbPractise/models"
)

//type Bus struct {
//	ID                   int       `json:"id" gorm:"primarykey"`           //母线ID
//	Name                 string    `json:"name"`                           //母线英文名
//	CName                string    `json:"c_name"`                         //母线中文名
//	Type                 string    `json:"type"`                           //卡片类型
//	BelongSubstationName string    `json:"belong_substation_name"`         //所属厂站英文名
//	Voltage              float32   `json:"voltage"`                        //电压等级
//	AddUser              string    `json:"add_user"`                       //添加数据的用户
//	AddTime              time.Time `json:"add_time" gorm:"autoUpdateTime"` //添加时间
//	Delete               int       `json:"delete"`                         //删除标志位
//}

//func AddBus(bus Bus) error {
//	if err := models.M8.Create(&bus).Error; err != nil {
//		return err
//	}
//	return nil
//}

func EditBusByName(name string, data interface{}) error {
	if err := models.M8.Model(&Bus{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
