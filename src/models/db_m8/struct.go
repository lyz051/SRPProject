package db_m8

import (
	"time"
)

type M8 struct {
	ID      int       `json:"id" gorm:"primarykey"`           //厂站ID
	Name    string    `json:"name"`                           //英文名
	CName   string    `json:"c_name"`                         //中文名
	AddTime time.Time `json:"add_time" gorm:"autoUpdateTime"` //添加时间
	Delete  int       `json:"delete"`                         //删除标志位
}

type PowerPlant struct {
	M8
	Type           string  `json:"type"`           //卡片类型
	BelongProvince string  `json:"blong_province"` //所属省网
	BelongCity     string  `json:"blong_city"`     //所属地市
	Voltage        float32 `json:"voltage"`        //电压等级
}

type Bus struct {
	M8
	Type                 string  `json:"type"`                   //卡片类型
	BelongSubstationName string  `json:"belong_substation_name"` //所属厂站英文名
	Voltage              float32 `json:"voltage"`                //电压等级
}

type ACLine struct {
	M8
	Type         string  `json:"type"`          //卡片类型
	STID_A       string  `json:"stid_a"`        //A端厂站
	STID_B       string  `json:"stid_b"`        //B端厂站
	Bus_A        string  `json:"bus_a"`         //A端母线名
	Bus_B        string  `json:"bus_b"`         //B端母线名
	Voltage      float32 `json:"voltage"`       //电压等级
	LinkFlag     int     `json:"link_flag"`     //区域联络线测点标志
	ParallelFlag int     `json:"parallel_flag"` //并联回路标识
	CascadeNum   int     `json:"cascade_num"`   //分段串联线路段号
	ParemeterID  int     `json:"paremeter_id"`  //当前使用的参数ID
	StartYear    int     `json:"start_year"`    //投运年
	StartMonth   int     `json:"start_month"`   //投运月
}

type DCLine struct {
	M8
	Type        string  `json:"type"`         //卡片类型
	ZLBus       string  `json:"zl_bus"`       //整流侧母线名
	ZLVolt      float32 `json:"zl_volt"`      //整流侧基电压
	NBBus       string  `json:"nb_bus"`       //逆变侧母线名
	NBVolt      float32 `json:"nb_volt"`      //逆变侧基电压
	ParameterID int     `json:"parameter_id"` //当前使用的参数ID
	DCSys       int     `json:"dc_sys"`       //所属直流系统
}

type Conventor struct {
	M8
	Type  string  `json:"type"`    //卡片类型
	Bus1  string  `json:"bus_1"`   //换流器母线名
	BusV1 float32 `json:"bus_v_1"` //换流器母线电压
	Bus2  string  `json:"bus_2"`   //换流变一次侧母线名
	BusV2 float32 `json:"bus_v_2"` //换流变一次侧基电压
	DCSys int     `json:"dc_sys"`  //所属直流系统
}

type DCSystem struct {
	M8
	Type    int    `json:"type"`     //直流系统极数
	Volt    int    `json:"volt"`     //额定直流电压
	Statue  int    `json:"statue"`   //状态
	TimeIn  string `json:"time_in"`  //投产日期
	TimeOut string `json:"time_out"` //退役日期
}

type Trans struct {
	M8
	Type        string  `json:"type"`         //卡片类型
	BelongSt    string  `json:"belong_st"`    //所属厂站英文名
	BusA        string  `json:"bus_a"`        //A端母线名
	BusB        string  `json:"bus_b"`        //B端母线名
	VoltageA    float32 `json:"voltage_a"`    //A端基准电压
	VoltageB    float32 `json:"voltage_b"`    //B端基准电压
	ParallelID  int     `json:"parallel_id"`  //并联回路标识
	ParameterID int     `json:"parameter_id"` //当前使用的参数ID
}
