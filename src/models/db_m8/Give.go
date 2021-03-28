package db_m8

import "encoding/json"

//根据ID
func GiveByID(id int, table string) map[string]interface{} {
	b := make(map[string]interface{})
	var a interface{}
	switch table {
	case "power_plant":
		a, _ = GetPowerPlantByID(id)
	case "bus":
		a, _ = GetBusByID(id)
	case "ac_line":
		a, _ = GetACLineByID(id)
	case "dc_line":
		a, _ = GetDCLineByID(id)
	case "dc_system":
		a, _ = GetDCSysByID(id)
	case "conventor":
		a, _ = GetConventorByID(id)
	case "trans":
		a, _ = GetTransByID(id)
	default:
		return nil
	}
	//a, _ := GetDCLineByID(id)
	j, _ := json.Marshal(a)
	json.Unmarshal(j, &b)
	return b
}

//根据英文名
func GiveByName(name string, table string) map[string]interface{} {
	b := make(map[string]interface{})
	var a interface{}
	switch table {
	case "power_plant":
		a, _ = GetPowerPlantByName(name)
	case "bus":
		a, _ = GetBusByName(name)
	case "ac_line":
		a, _ = GetACLineByName(name)
	case "dc_line":
		a, _ = GetDCLineByName(name)
	case "dc_system":
		a, _ = GetDCSysByName(name)
	case "conventor":
		a, _ = GetConventorByName(name)
	case "trans":
		a, _ = GetTransByName(name)
	default:
		return nil
	}
	j, _ := json.Marshal(a)
	json.Unmarshal(j, &b)
	return b
}

//powerplant
func GivePlowerPlantByVoltage(voltage float32) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	//c2 := make(map[string]interface{})
	d, _ := GetAllPowerPlantByVoltage(voltage)
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
func GiveBusByVoltage(voltage float32) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllBusByVoltage(voltage)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}

//acline
func GiveACLineBySTIDA(stid_a string) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllACLineSTIDA(stid_a)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}

func GiveACLineBySTIDB(stid_b string) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllACLineSTIDB(stid_b)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}

func GiveACLineByBusA(busa string) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllACLineByBusA(busa)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}

func GiveACLineByBusB(busb string) map[int]map[string]interface{} {
	c1 := make(map[int]map[string]interface{})
	d, _ := GetAllACLineByBusB(busb)
	for i := range d {
		var c2 map[string]interface{}
		k, _ := json.Marshal(d[i])
		json.Unmarshal(k, &c2)
		c1[i] = c2
	}
	return c1
}
