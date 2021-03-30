package db_m10

import (
	"dbPractise/models"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type DynamicModel struct {
	ID          int       `json:"id" gorm:"primarykey"`           //动态元件ID
	Name        string    `json:"name"`                           //动态元件名称
	MName       string    `json:"m_name"`                         //动态元件英文名称
	Type        int       `json:"type"`                           //元件类型
	SubType     float32   `json:"sub_type"`                       //元件子类型
	ChgCode     float32   `json:"chg_code"`                       //Chg码
	BusName     string    `json:"bus_name"`                       //母线英文名
	KV          float32   `json:"kv"`                             //基准电压
	Mid         float32   `json:"mid"`                            //发电机识别码
	ParameterID int       `json:"parameter_id"`                   //当前使用的参数ID
	AddUser     string    `json:"add_user"`                       //添加数据的用户
	AddTime     time.Time `json:"add_time" gorm:"autoUpdateTime"` //添加时间
	Delete      int       `json:"delete"`                         //删除标志位
	Ready       int       `json:"ready"`                          //是否已填入数据
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
	err := models.M10.Select("m_name").Where("m_name = ?", mname).First(&dynamicmodel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if dynamicmodel.MName != "" {
		return true, nil
	}

	return false, nil
}

//根据ID检索动态元件
func GetDynamicModelByID(id int, delete int) (*DynamicModel, error) {
	var dynamicmodel DynamicModel
	var err error
	if delete == 1 || delete == 0 {
		err = models.M10.Where("id = ? AND `delete` = ? ", id, delete).First(&dynamicmodel).Error
	} else if delete == 2 {
		err = models.M10.Where("id = ?", id).First(&dynamicmodel).Error
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dynamicmodel, nil
}

//根据名称检索
func GetDynamicModelByName(name string, delete int) (*DynamicModel, error) {
	var dynamicmodel DynamicModel
	var err error
	if delete == 1 || delete == 0 {
		err = models.M10.Where("name = ? AND `delete` = ? ", name, delete).First(&dynamicmodel).Error
	} else if delete == 2 {
		err = models.M10.Where("name = ?", name).First(&dynamicmodel).Error
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dynamicmodel, nil
}

//根据英文名检索
func GetDynamicModelByMName(mname string, delete int) (*DynamicModel, error) {
	var dynamicmodel DynamicModel
	var err error
	if delete == 1 || delete == 0 {
		err = models.M10.Where("m_name = ? AND `delete` = ? ", mname, delete).First(&dynamicmodel).Error
	} else if delete == 2 {
		err = models.M10.Where("m_name = ?", mname).First(&dynamicmodel).Error
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dynamicmodel, nil
}

//根据基准电压检索
func GetDynamicModelByKV(kv float32, delete int) ([]DynamicModel, error) {
	var dynamicmodel []DynamicModel
	var err error
	if delete == 0 || delete == 1 {
		err = models.M10.Where("kv = ? AND `delete` = ? ", kv, delete).Find(&dynamicmodel).Error
	} else if delete == 2 {
		err = models.M10.Where("kv = ?", kv).Find(&dynamicmodel).Error
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	temp := make([]DynamicModel, len(dynamicmodel))
	for aIndex, a := range dynamicmodel {
		temp[aIndex] = a
	}
	return temp, nil
}

func GiveByID(id int, delete int) map[string]interface{} {
	b := make(map[string]interface{})
	a, _ := GetDynamicModelByID(id, delete)
	j, _ := json.Marshal(a)
	json.Unmarshal(j, &b)
	return b
}

func GiveByName(name string, delete int) map[string]interface{} {
	b := make(map[string]interface{})
	a, _ := GetDynamicModelByName(name, delete)
	j, _ := json.Marshal(a)
	json.Unmarshal(j, &b)
	return b
}

func GiveByMName(mname string, delete int) map[string]interface{} {
	b := make(map[string]interface{})
	a, _ := GetDynamicModelByMName(mname, delete)
	j, _ := json.Marshal(a)
	json.Unmarshal(j, &b)
	return b
}

func GiveByKV(kv float32, delete int) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetDynamicModelByKV(kv, delete)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
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
	if err := models.M10.Model(&DynamicModel{}).Where("m_name = ?", mname).Updates(data).Error; err != nil {
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
	if flag, _ := ExistDynamicModelByMName(mname); flag == true {
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
