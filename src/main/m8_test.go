package main

import (
	"dbPractise/models"
	"dbPractise/models/db_m8"
	"dbPractise/setting"
	"fmt"
	"testing"
)

func TestPowerPlant(t *testing.T) {
	setting.Setup()
	models.Setup()

	// 增加
	//pp := map[string]interface{}{
	//	"name":            "changzhan4",
	//	"c_name":          "厂站4",
	//	"type":            "不知道是啥",
	//	"belong_province": "GuangDong",
	//	"belong_city":     "Guangzhou",
	//	"voltage":         float32(220),
	//	"add_user":        "lyz",
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
	//}
	//
	//powerplant.AddData()
	//if err := db_m8.AddPowerPlant(powerplant); err != nil {
	//	log.Println(err.Error())
	//}

	//Delete 测试
	//name := "changzhan4"
	//db_m8.DeletePowerPlant(name)

	//查询测试
	//查询:delete：已删除：1，未删除：0,全部：2
	//b := db_m8.GiveByID(1, 0,"power_plant")
	//fmt.Println(b)
	//a := db_m8.GiveByName("changzhan4",0, "power_plant")
	//fmt.Println(a)

	c := db_m8.GivePlowerPlantByVoltage(220, 0)
	fmt.Println(c)
}

func TestBus(t *testing.T) {
	setting.Setup()
	models.Setup()

	//增加
	//pp := map[string]interface{}{
	//	"name":            		"bus3",
	//	"c_name":          		"母线3",
	//	"type":            		"不知道是啥",
	//	"belong_substation":	"changzhan1",
	//	"voltage":         		float32(220),
	//	"add_user":        		"lyz",
	//}
	//
	//bus := db_m8.Bus{
	//	Name:           		pp["name"].(string),
	//	CName:          		pp["c_name"].(string),
	//	Type:           		pp["type"].(string),
	//	BelongSubstationName: 	pp["belong_substation"].(string),
	//	Voltage:        		pp["voltage"].(float32),
	//	AddUser:        		pp["add_user"].(string),
	//}
	//bus.AddData()

	//删除
	//name := "bus3"
	//db_m8.DeleteBusPlant(name)

	//查询
	a := db_m8.GiveByID(1, 0, "bus")
	fmt.Println(a)

	b := db_m8.GiveByName("bus1", 0, "bus")
	fmt.Println(b)

	c := db_m8.GiveBusByVoltage(220, 0)
	fmt.Println(c)
}

func TestACLine(t *testing.T) {
	setting.Setup()
	models.Setup()

	//增加
	//pp := map[string]interface{}{
	//	"name":            		"ACLine3",
	//	"c_name":          		"交流线路3",
	//	"type":            		"不知道是啥",
	//	"stid_a":				"changzhan1",
	//	"stid_b":				"changzhan2",
	//	"bus_a":				"bus1",
	//	"bus_b":				"bus2",
	//	"voltage":         		float32(220),
	//	"link_flag":			1,
	//	"parallel_flag":		1,
	//	"cascade_num":			1,
	//	"parameter_id":			1,
	//	"start_year":			2000,
	//	"start_month":			2,
	//	"add_user":        		"lyz",
	//}
	//
	//acline := db_m8.ACLine{
	//	Name:           		pp["name"].(string),
	//	CName:          		pp["c_name"].(string),
	//	Type:           		pp["type"].(string),
	//	STID_A:  				pp["stid_a"].(string),
	//	STID_B: 				pp["stid_b"].(string),
	//	Bus_A:  				pp["bus_a"].(string),
	//	Bus_B:      			pp["bus_b"].(string),
	//	Voltage:        		pp["voltage"].(float32),
	//	LinkFlag:  				pp["link_flag"].(int),
	//	ParallelFlag: 			pp["parallel_flag"].(int),
	//	CascadeNum: 			pp["cascade_num"].(int),
	//	ParameterID: 			pp["parameter_id"].(int),
	//	StartYear:  			pp["start_year"].(int),
	//	StartMonth:  			pp["start_month"].(int),
	//	AddUser:        		pp["add_user"].(string),
	//}
	//acline.AddData()

	//删除
	//name := "ACLine3"
	//db_m8.DeleteACLine(name)

	//查询
	//a := db_m8.GiveByID(1,0, "ac_line")
	//fmt.Println(a)
	//
	//b := db_m8.GiveByName("ACLine1", 0,"ac_line")
	//fmt.Println(b)
	//
	//c, d := db_m8.GiveACLineByBusA("bus1",0), db_m8.GiveACLineByBusB("bus2",0)
	//fmt.Println(c)
	//fmt.Println(d)
	//
	//e, f := db_m8.GiveACLineBySTIDA("changzhan1",0), db_m8.GiveACLineBySTIDB("changzhan2",0)
	//fmt.Println(e)
	//fmt.Println(f)

	//修改当前使用参数组ID
	new_id := 2
	db_m8.UpdateACLine("ACLine3", new_id)
}

func TestDCLine(t *testing.T) {
	setting.Setup()
	models.Setup()

	//增加
	//pp := map[string]interface{}{
	//	"name":            		"DCLine3",
	//	"c_name":          		"直流线路3",
	//	"type":            		"不知道是啥",
	//	"zl_bus":				"bus1",
	//	"zl_volt":				float32(220),
	//	"nb_bus":				"bus2",
	//	"nb_volt":				float32(220),
	//	"paremeter_id":			1,
	//	"dc_sys":				"dcsys1",
	//	"add_user":        		"lyz",
	//}
	//
	//dcline := db_m8.DCLine{
	//	Name:           		pp["name"].(string),
	//	CName:          		pp["c_name"].(string),
	//	Type:           		pp["type"].(string),
	//	ZLBus:  				pp["zl_bus"].(string),
	//	ZLVolt:  				pp["zl_volt"].(float32),
	//	NBBus:  				pp["nb_bus"].(string),
	//	NBVolt:  				pp["nb_volt"].(float32),
	//	ParameterID:  			pp["paremeter_id"].(int),
	//	DCSys:  				pp["dc_sys"].(string),
	//	AddUser:        		pp["add_user"].(string),
	//}
	//dcline.AddData()

	//删除
	//name := "DCLine3"
	//db_m8.DeleteDCLine(name)

	//查询
	a := db_m8.GiveByID(1, 0, "dc_line")
	fmt.Println(a)

	b := db_m8.GiveByName("DCLine1", 0, "dc_line")
	fmt.Println(b)

	//修改参数组ID
	//new_id := 2
	//db_m8.UpdateDCLine("DCLine3", new_id)
}

func TestDCSys(t *testing.T) {
	setting.Setup()
	models.Setup()

	//增加
	//pp := map[string]interface{}{
	//	"name":            		"dcsys3",
	//	"c_name":          		"直流系统3",
	//	"type":            		3,
	//	"volt":					220,
	//	"statue":				1,
	//	"time_in":				"2000-01-01",
	//	"time_out":				"2020-12-31",
	//	"add_user":        		"lyz",
	//}
	//
	//dcsys := db_m8.DCSystem{
	//	Name:           		pp["name"].(string),
	//	CName:          		pp["c_name"].(string),
	//	Type:           		pp["type"].(int),
	//	Volt:  					pp["volt"].(int),
	//	Statue:  				pp["statue"].(int),
	//	TimeIn:  				pp["time_in"].(string),
	//	TimeOut:  				pp["time_out"].(string),
	//	AddUser:        		pp["add_user"].(string),
	//}
	//dcsys.AddData()

	//删除
	//name := "dcsys3"
	//db_m8.DeleteDCSystem(name)

	//查询
	a, b := db_m8.GiveByID(1, 0, "dc_system"), db_m8.GiveByName("dcsys1", 0, "dc_system")
	fmt.Println(a)
	fmt.Println(b)
}

func TestConventor(t *testing.T) {
	setting.Setup()
	models.Setup()

	//增加
	//pp := map[string]interface{}{
	//	"name":            		"conventor3",
	//	"c_name":          		"换流器3",
	//	"type":            		"i dont know",
	//	"bus_1":				"bus1",
	//	"bus_2":				"bus2",
	//	"bus_v_1":				float32(220),
	//	"bus_v_2":				float32(220),
	//	"dc_sys":				"dcsys1",
	//	"add_user":        		"lyz",
	//}
	//
	//conventor := db_m8.Conventor{
	//	Name:           		pp["name"].(string),
	//	CName:          		pp["c_name"].(string),
	//	Type:           		pp["type"].(string),
	//	Bus1:  					pp["bus_1"].(string),
	//	Bus2:  					pp["bus_2"].(string),
	//	BusV1:  				pp["bus_v_1"].(float32),
	//	BusV2:  				pp["bus_v_2"].(float32),
	//	DCSys:  				pp["dc_sys"].(string),
	//	AddUser:        		pp["add_user"].(string),
	//}
	//conventor.AddData()

	//删除
	//name := "conventor3"
	//db_m8.DeleteConventor(name)

	//查询
	a, b := db_m8.GiveByID(1, 0, "conventor"), db_m8.GiveByName("conventor1", 0, "conventor")
	fmt.Println(a)
	fmt.Println(b)
}

func TestTrans(t *testing.T) {
	setting.Setup()
	models.Setup()

	//增加
	//pp := map[string]interface{}{
	//	"name":            		"trans3",
	//	"c_name":          		"变压器3",
	//	"type":            		"i dont know",
	//	"belong_st":			"changzhan1",
	//	"bus_a":				"bus1",
	//	"bus_b":				"bus2",
	//	"voltage_a":			float32(220),
	//	"voltage_b":			float32(220),
	//	"parallel_flag":		1,
	//	"parameter_id":			1,
	//	"add_user":        		"lyz",
	//}
	//
	//trans := db_m8.Trans{
	//	Name:           		pp["name"].(string),
	//	CName:          		pp["c_name"].(string),
	//	Type:           		pp["type"].(string),
	//	BelongSt:  				pp["belong_st"].(string),
	//	BusA:  					pp["bus_a"].(string),
	//	BusB:  					pp["bus_b"].(string),
	//	VoltageA:  				pp["voltage_a"].(float32),
	//	VoltageB:  				pp["voltage_b"].(float32),
	//	ParameterID:  			pp["parameter_id"].(int),
	//	ParallelFlag:  			pp["parallel_flag"].(int),
	//	AddUser:        		pp["add_user"].(string),
	//}
	//trans.AddData()

	//删除
	//name := "trans3"
	//db_m8.DeleteTrans(name)

	//查询
	a, b := db_m8.GiveByID(1, 0, "trans"), db_m8.GiveByName("trans1", 0, "trans")
	fmt.Println(a)
	fmt.Println(b)

	//修改参数组ID
	new_id := 2
	db_m8.UpdateTrans("trans3", new_id)
}
