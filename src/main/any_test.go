package main

import (
	"dbPractise/models"
	"dbPractise/models/db_m8"
	"dbPractise/setting"
	"testing"
)

func TestAny(t *testing.T) {
	setting.Setup()
	models.Setup()

	// Power Plant test
	////timenow := time.Now().Format("2006-01-02 15:04:05")
	pp := map[string]interface{}{
		"name":            "changzhan3",
		"c_name":          "厂站1",
		"type":            "不知道是啥",
		"belong_province": "GuangDong",
		"belong_city":     "Guangzhou",
		"voltage":         float32(220),
		"add_user":        "lyz",
		//"add_time":			timenow,
		//"delete": 0,
	}

	powerplant := db_m8.PowerPlant{
		Name:           pp["name"].(string),
		CName:          pp["c_name"].(string),
		Type:           pp["type"].(string),
		BelongProvince: pp["belong_province"].(string),
		BelongCity:     pp["belong_city"].(string),
		Voltage:        pp["voltage"].(float32),
		AddUser:        pp["add_user"].(string),
		//AddTime: 		pp["add_time"].(string),
		//Delete: pp["delete"].(int),
	}

	powerplant.AddData()
	//if err := db_m8.AddPowerPlant(powerplant); err != nil {
	//	log.Println(err.Error())
	//}

	// ac_branch 关联 line_id 测试
	// 0. Add dynamic_line
	//var dl db_m2.DynamicLine
	//dl.ID=1
	//dl.LineName="test"
	//dl.Add()

	// 1. Add ac_branch
	//db_m9.Setup()
	//var acb db_m9.ACBranch
	//acb.LineID = 1
	//acb.Add()
	// 2. Association
	//	dynamicLine, _ := acb.GetDyanmicLine()
	//	log.Println(dynamicLine)
}
