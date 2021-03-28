package main

import (
	"dbPractise/models"
	"dbPractise/setting"
	"testing"
)

func TestDynamicModel(t *testing.T) {
	setting.Setup()
	models.Setup()

	//增加
	//pp := map[string]interface{}{
	//	"name":            		"动态模型3",
	//	"m_name":          		"dynamicmodel3",
	//	"type":            		1,
	//	"sub_type":				float32(2.5),
	//	"chg_code":				float32(23.34),
	//	"bus_name":				"bus1",
	//	"kv":					float32(220),
	//	"mid":					float32(23),
	//	"parameter_id":			1,
	//	"add_user":        		"lyz",
	//}
	//
	//dm := db_m10.DynamicModel{
	//	Name:           		pp["name"].(string),
	//	MName:          		pp["m_name"].(string),
	//	Type:           		pp["type"].(int),
	//	SubType:  				pp["sub_type"].(float32),
	//	ChgCode:  				pp["chg_code"].(float32),
	//	BusName:  				pp["bus_name"].(string),
	//	KV:  					pp["kv"].(float32),
	//	Mid:  					pp["mid"].(float32),
	//	ParameterID:  			pp["parameter_id"].(int),
	//	AddUser:        		pp["add_user"].(string),
	//}
	//dm.AddData()

	//删除
	//m_name := "dynamicmodel3"
	//db_m10.DeleteDynamicModelByMName(m_name)

	//查询
	//a, b, c := db_m10.GiveByID(1), db_m10.GiveByMName("dynamicmodel2"), db_m10.GiveByName("动态模型3")
	//fmt.Println(a,"\n", b, "\n", c)
	//
	//d := db_m10.GiveByKV(220)
	//fmt.Println(d)

	//修改参数组ID
	//new_id := 2
	//db_m10.UpdateDynamicModelByMName("dynamicmodel3", new_id)
}
