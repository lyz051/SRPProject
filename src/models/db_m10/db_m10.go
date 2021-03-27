package db_m10

import (
	"dbPractise/models"
	"gorm.io/gorm"
	"time"
)

type DynamicModel struct {
	ID          int       `json:"id"`           //动态元件ID
	Name        string    `json:"name"`         //动态元件名称
	MName       string    `json:"mname"`        //动态元件英文名称
	Type        int       `json:"type"`         //元件类型
	SubType     float32   `json:"sub_type"`     //元件子类型
	ChgCode     float32   `json:"chg_code"`     //Chg码
	BusName     string    `json:"bus_name"`     //母线英文名
	KV          float32   `json:"kv"`           //基准电压
	Mid         float32   `json:"mid"`          //发电机识别码
	ParameterID int       `json:"parameter_id"` //当前使用的参数ID
	AddUser     string    `json:"add_user"`     //添加数据的用户
	AddTime     time.Time `json:"add_time"`     //添加时间
	Delete      int       `json:"delete"`       //删除标志位
	Ready       int       `json:"ready"`        //是否已填入数据
}

//增加动态设备
func (a *DynamicModel) AddData() {
	flag, _ := ExistDynamicModelByName(a.Name)
	if flag == true {
		return
	} else {
		a.Add()
	}
}

func AddDynamicModel(dynamicmodel DynamicModel) error {
	if err := models.M10.Create(&dynamicmodel).Error; err != nil {
		return err
	}
	return nil
}
func (b *DynamicModel) Add() error {
	if err := models.M10.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

//检索部分
//检索当前是否已存在该设备
func ExistDynamicModelByName(name string) (bool, error) {
	var dynamicmodel DynamicModel
	err := models.M10.Select("name").Where("name = ?", name).First(&dynamicmodel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if dynamicmodel.Name != "" {
		return true, nil
	}

	return false, nil
}
func ExistDynamicModelByMName(mname string) (bool, error) {
	var dynamicmodel DynamicModel
	err := models.M10.Select("mname").Where("mname = ?", mname).First(&dynamicmodel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if dynamicmodel.Name != "" {
		return true, nil
	}

	return false, nil
}

//根据ID检索动态元件
func GetDynamicModelByID(id int) (*DynamicModel, error) {
	var dynamicmodel DynamicModel
	err := models.M10.Where("id = ?", id).First(&dynamicmodel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &dynamicmodel, nil
}

//根据名称检索
func GetDynamicModelByName(name string) (*DynamicModel, error) {
	var dynamicmodel DynamicModel
	err := models.M10.Where("name = ?", name).First(&dynamicmodel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &dynamicmodel, nil
}

//根据英文名检索
func GetDynamicModelByMName(mname string) (*DynamicModel, error) {
	var dynamicmodel DynamicModel
	err := models.M10.Where("mname = ?", mname).First(&dynamicmodel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &dynamicmodel, nil
}

//根据基准电压检索
func GetDynamicModelByKV(kv float32) ([]*DynamicModel, error) {
	var dynamicmodel []DynamicModel
	err := models.M10.Where("kv = ?", kv).Find(&dynamicmodel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	temp := make([]*DynamicModel, len(dynamicmodel))
	for aIndex, a := range dynamicmodel {
		temp[aIndex] = &a
	}
	return temp, nil
}

//修改参数部分
//根据动态元件名称修改元素
func EditDynamicModelByName(name string, data interface{}) error {
	if err := models.M10.Model(&DynamicModel{}).Where("name = ?", name).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

//根据动态元件英文名修改元素
func EditDynamicModelByMName(mname string, data interface{}) error {
	if err := models.M10.Model(&DynamicModel{}).Where("mname = ?", mname).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

//删除元件部分
//根据动态元件名称删除设备
func DeleteDynamicModelByName(name string) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistDynamicModelByName(name); flag == true {
		EditDynamicModelByName(name, a)
	} else {
		return
	}
}

//根据动态元件英文名删除设备
func DeleteDynamicModelByMName(mname string) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistDynamicModelByName(mname); flag == true {
		EditDynamicModelByMName(mname, a)
	} else {
		return
	}
}

//修改当前使用参数部分
//根据名称修改
func UpdateDynamicModelByName(name string, id int) {
	a := map[string]interface{}{
		"parameter_id": id,
	}
	if flag, _ := ExistDynamicModelByName(name); flag == true {
		EditDynamicModelByName(name, a)
	} else {
		return
	}
}

//根据引英文名修改
func UpdateDynamicModelByMName(mname string, id int) {
	a := map[string]interface{}{
		"parameter_id": id,
	}
	if flag, _ := ExistDynamicModelByMName(mname); flag == true {
		EditDynamicModelByMName(mname, a)
	} else {
		return
	}
}
