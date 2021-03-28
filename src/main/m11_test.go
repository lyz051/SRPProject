package main

import (
	"dbPractise/models"
	"dbPractise/models/db_m11"
	"dbPractise/setting"
	"fmt"
	"testing"
)

func TestGeneratorParameter(t *testing.T) {
	setting.Setup()
	models.Setup()

	//增加
	//pp := map[string]interface{}{
	//	"type":            		"miea",
	//	"g_name":				"dynamicmodel3",
	//	"kv":					float32(220),
	//	"mid":					float32(232),
	//	"emws":					float32(23),
	//	"p":					float32(2.1),
	//	"q":					float32(3.2),
	//	"mva_base":				float32(220),
	//	"ra":					float32(21),
	//	"xdp":					float32(1.1),
	//	"xqp":					float32(1.1),
	//	"xd":					float32(1.2),
	//	"xq":					float32(1.2),
	//	"td0p":					float32(1.3),
	//	"tq0p":					float32(1.3),
	//	"n":					float32(2.2),
	//	"a":					float32(1.1),
	//	"b":					float32(1.01),
	//	"d":					float32(1.02),
	//	"add_user":        		"lyz",
	//}
	//
	//gp := db_m11.GeneratorParemeter{
	//	Type:           		pp["type"].(string),
	//	GName:  			    pp["g_name"].(string),
	//	KV:  					pp["kv"].(float32),
	//	Mid:  					pp["mid"].(float32),
	//	Emws:  					pp["emws"].(float32),
	//	P:  					pp["p"].(float32),
	//	Q:  					pp["q"].(float32),
	//	MVABase:  				pp["mva_base"].(float32),
	//	Ra:  					pp["ra"].(float32),
	//	Xdp:  					pp["xdp"].(float32),
	//	Xqp:  					pp["xqp"].(float32),
	//	Xd:  					pp["xd"].(float32),
	//	Xq:  					pp["xq"].(float32),
	//	Td0p:  					pp["td0p"].(float32),
	//	Tq0p:  					pp["tq0p"].(float32),
	//	N:  					pp["n"].(float32),
	//	A:  					pp["a"].(float32),
	//	B:  					pp["b"].(float32),
	//	D:  					pp["d"].(float32),
	//	AddUser:        		pp["add_user"].(string),
	//}
	//gp.AddData()

	//删除
	//id := 3
	//db_m11.DeleteGeneratorParemeterByID(id)

	//查找
	a := db_m11.GiveByGName("dynamicmodel1")
	fmt.Println(a)
}
