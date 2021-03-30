package main

import (
	"dbPractise/models"
	"dbPractise/models/db_m9"
	"dbPractise/setting"
	"log"
	"testing"
)

func TestM9(t *testing.T) {
	setting.Setup()
	models.Setup()

	// Power Plant test
	////timenow := time.Now().Format("2006-01-02 15:04:05")
	//pp := map[string]interface{}{
	//	"name":            "changzhan1",
	//	"c_name":          "厂站1",
	//	"type":            "不知道是啥",
	//	"belong_province": "GuangDong",
	//	"belong_city":     "Guangzhou",
	//	"voltage":         float32(220),
	//	"add_user":        "lyz",
	//	//"add_time":			timenow,
	//	//"delete": 0,
	//}
	//
	//powerplant := db_m8.PowerPlant{
	//	Name:           pp["name"].(string),
	//	CName:          pp["c_name"].(string),
	//	Type:           pp["type"].(string),
	//	BelongProvince: pp["belong_province"].(string),
	//	BelongCity:     pp["belong_city"].(string),
	//	Voltage:        pp["voltage"].(float32),
	//	AddUser:        pp["add_user"].(string),
	//	//AddTime: 		pp["add_time"].(string),
	//	//Delete: pp["delete"].(int),
	//}
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
	////acb.Add()
	//// 2. Association
	//dynamicLine, _ := acb.GetDyanmicLine()
	//log.Println(dynamicLine)

	// M9 反射测试
	var acb db_m9.ACBranch
	var trans db_m9.Transformer
	var dcb db_m9.DCBranch
	var cvt db_m9.Convertor
	acb.LineName = "testLine"
	trans.TransName = "testTrans"
	trans.AddUser = "lzt"
	dcb.LineName = "testLine"
	cvt.AddUser = "lzt"
	cvt.QDLEDZ = 9
	cvt.TransName = "testTrans"

	// Add
	//err := db_m9.Add(cvt)
	//if err != nil {
	//	log.Println(err)
	//}

	// Update
	//acb.ID=6
	//trans.ID=2
	//data1:=map[string]interface{}{
	//	"ParallelNum":   9,
	//}
	//data2:=map[string]interface{}{
	//	"Tg":            1.,
	//}
	//db_m9.Update(acb,data1)
	//db_m9.Update(trans,data2)

	//Get
	//id := 1
	////exist, err := db_m9.ExistACBranchByID(id)
	////exist, err := db_m9.ExistDCBranchByID(id)
	////exist, err := db_m9.ExistConvertorByID(id)
	//exist, err := db_m9.ExistTransformerByID(id)
	//if exist && err == nil {
	//	//acb, err := db_m9.GetACBranchByID(id)
	//	//ins, err := db_m9.GetDCBranchByID(id)
	//	//ins, err := db_m9.GetConvertorByID(id)
	//	ins, err := db_m9.GetTransformerByID(id)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	log.Printf("%#v", ins)
	//}

	//Delete
	//acb.ID = 10
	////db_m9.Delete(acb)
	//db_m9.DeleteACBranchByID(acb.ID)

	// Give
	//id:=2
	//ins:=db_m9.GiveByID(id,"ac_branch")
	name := "testTrans"
	ins := db_m9.GiveByParentName(name, "convertor")
	log.Println(ins)
}
