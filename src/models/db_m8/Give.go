package db_m8

import (
	"encoding/json"
)

//根据ID
func GiveByID(id int, delete int, table string) map[string]interface{} {
	b := make(map[string]interface{})
	var a interface{}
	switch table {
	case "power_plant":
		a, _ = GetPowerPlantByID(id, delete)
	case "bus":
		a, _ = GetBusByID(id, delete)
	case "ac_line":
		a, _ = GetACLineByID(id, delete)
	case "dc_line":
		a, _ = GetDCLineByID(id, delete)
	case "dc_system":
		a, _ = GetDCSysByID(id, delete)
	case "conventor":
		a, _ = GetConventorByID(id, delete)
	case "trans":
		a, _ = GetTransByID(id, delete)
	default:
		return nil
	}
	//a, _ := GetDCLineByID(id)
	j, _ := json.Marshal(a)
	json.Unmarshal(j, &b)
	return b
}

//根据英文名
func GiveByName(name string, delete int, table string) map[string]interface{} {
	b := make(map[string]interface{})
	var a interface{}
	switch table {
	case "power_plant":
		a, _ = GetPowerPlantByName(name, delete)
	case "bus":
		a, _ = GetBusByName(name, delete)
	case "ac_line":
		a, _ = GetACLineByName(name, delete)
	case "dc_line":
		a, _ = GetDCLineByName(name, delete)
	case "dc_system":
		a, _ = GetDCSysByName(name, delete)
	case "conventor":
		a, _ = GetConventorByName(name, delete)
	case "trans":
		a, _ = GetTransByName(name, delete)
	default:
		return nil
	}
	j, _ := json.Marshal(a)
	json.Unmarshal(j, &b)
	return b
}

//powerplant
func GivePlowerPlantByVoltage(voltage float32, delete int, limit int) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	//c2 := make(map[string]interface{})
	d, _ := GetAllPowerPlantByVoltage(voltage, delete, limit)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		//fmt.Println(c2)
		//fmt.Println(c2["id"])
		c1[i] = c2
	}
	return c1
}

//bus
func GiveBusByVoltage(voltage float32, delete int, limit int) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllBusByVoltage(voltage, delete, limit)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}

//acline
func GiveACLineBySTIDA(stid_a string, delete int, limit int) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllACLineSTIDA(stid_a, delete, limit)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}

func GiveACLineBySTIDB(stid_b string, delete int, limit int) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllACLineSTIDB(stid_b, delete, limit)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}

func GiveACLineByBusA(busa string, delete int, limit int) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllACLineByBusA(busa, delete, limit)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}

func GiveACLineByBusB(busb string, delete int, limit int) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllACLineByBusB(busb, delete, limit)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}
