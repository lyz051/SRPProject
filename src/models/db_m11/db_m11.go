package db_m11

import (
	"dbPractise/models"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type GeneratorParemeter struct {
	ID        int       `json:"id" gorm:"primarykey"`           //参数组ID
	Type      string    `json:"type"`                           //卡片类型
	GName     string    `json:"g_name"`                         //动态元件名称
	KV        float32   `json:"kv"`                             //发电机基准电压
	Mid       float32   `json:"mid"`                            //发电机识别码
	Emws      float32   `json:"emws"`                           //发电机动能
	P         float32   `json:"p"`                              //发电机初始有功占母线潮流功率比例
	Q         float32   `json:"q"`                              //发电机初始无功占母线潮流功率比例
	MVABase   float32   `json:"mva_base"`                       //发电机标幺基准容量
	Ra        float32   `json:"ra"`                             //定子电阻
	Xdp       float32   `json:"xdp"`                            //直轴暂态电抗
	Xqp       float32   `json:"xqp"`                            //交轴暂态电抗
	Xd        float32   `json:"xd"`                             //直轴不饱和同步电抗
	Xq        float32   `json:"xq"`                             //交轴不饱和同步电抗
	Td0p      float32   `gorm:"column:td0p"`                    //饱和相关的指数
	Tq0p      float32   `gorm:"column:tq0p"`                    //饱和相关的系数
	N         float32   `json:"n"`                              //饱和相关的系数
	A         float32   `json:"a"`                              //电机阻尼转矩系数
	B         float32   `json:"b"`                              //添加数据的用户
	D         float32   `json:"d"`                              //添加时间
	AddUser   string    `json:"add_user"`                       //删除标志位
	AddTime   time.Time `json:"add_time" gorm:"autoUpdateTime"` //添加时间
	Delete    int       `json:"delete"`                         //删除标志位
	Ready     int       `json:"ready"`                          //是否已填入数据
	Paremeter int       `json:"paremeter"`                      //参数的类型
}

//增加动态元件参数

func (a GeneratorParemeter) AddData() error {
	exist, err := ExistGeneratorParemeterByInstance(a)
	if !exist && err == nil {
		a.Add()
	}
	if !exist && err != nil {
		return err
	}
	return nil
}

func (b *GeneratorParemeter) Add() error {
	if err := models.M11.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

//检索部分
//根据参数组ID查找该组参数在表单中是否存在
func ExistGeneratorParemeterByID(id int) (bool, error) {
	var generatorparemeter GeneratorParemeter
	err := models.M11.Select("id").Where("id = ?", id).First(&generatorparemeter).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if generatorparemeter.ID > 0 {
		return true, nil
	}
	return false, nil
}

func ExistGeneratorParemeterByInstance(ins GeneratorParemeter) (bool, error) {
	var GPID []int
	err := models.M11.Model(&GeneratorParemeter{}).Select("id").Where("`delete` != ?", 1).
		Where(map[string]interface{}{
			"g_name":    ins.GName,
			"kv":        ins.KV,
			"mid":       ins.Mid,
			"emws":      ins.Emws,
			"p":         ins.P,
			"q":         ins.Q,
			"xq":        ins.Xq,
			"xd":        ins.Xd,
			"mva_base":  ins.MVABase,
			"ra":        ins.Ra,
			"xdp":       ins.Xdp,
			"xqp":       ins.Xqp,
			"td0p":      ins.Td0p,
			"tq0p":      ins.Tq0p,
			"n":         ins.N,
			"a":         ins.A,
			"b":         ins.B,
			"d":         ins.D,
			"paremeter": ins.Paremeter,
		}).Pluck("id", &GPID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(GPID) > 0 {
		return true, nil
	}

	return false, nil
}

//根据动态元件名称检索全部该动态元件的参数
func GetGeneratorParemeterByGName(gname string, delete int) ([]GeneratorParemeter, error) {
	var generator []GeneratorParemeter
	var err error
	if delete == 0 || delete == 1 {
		err = models.M11.Where("g_name = ? AND `delete` = ? ", gname, delete).Find(&generator).Error
	} else if delete == 2 {
		err = models.M11.Where("g_name = ? ", gname).Find(&generator).Error
	} else {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	temp := make([]GeneratorParemeter, len(generator))
	for aIndex, a := range generator {
		temp[aIndex] = a
	}
	return temp, nil
}

func GiveByGName(gname string, delete int) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetGeneratorParemeterByGName(gname, delete)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}

//删除部分，根据参数组ID删除该参数组
//修改参数
func EditGeneratorParemeterByID(id int, data interface{}) error {
	if err := models.M11.Model(&GeneratorParemeter{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

//修改删除标志位
func DeleteGeneratorParemeterByID(id int) {
	a := map[string]interface{}{
		"delete": 1,
	}
	if flag, _ := ExistGeneratorParemeterByID(id); flag == true {
		EditGeneratorParemeterByID(id, a)
	} else {
		return
	}
}
