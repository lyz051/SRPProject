package db_m9

import "encoding/json"

func GiveByID(id int, table string) map[string]interface{} {
	var b map[string]interface{}
	var a interface{}
	switch table {
	case "ac_branch":
		a, _ = GetACBranchByID(id)
	case "transformer":
		a, _ = GetTransformerByID(id)
	case "dc_branch":
		a, _ = GetDCBranchByID(id)
	case "convertor":
		a, _ = GetConvertorByID(id)
	default:
		return nil
	}
	j, _ := json.Marshal(a)
	json.Unmarshal(j, &b)
	return b
}

func GiveByParentName(name string, table string) map[int]map[string]interface{} {
	b := make(map[int]map[string]interface{})
	//var a interface{}
	switch table {
	case "ac_branch":
		a, _ := GetACBranchByParentName(name)
		for i := 0; i < len(a); i++ {
			var temp map[string]interface{}
			j, _ := json.Marshal(a[i])
			json.Unmarshal(j, &temp)
			b[i] = temp
		}
	case "transformer":
		a, _ := GetTransformerByParentName(name)
		for i := 0; i < len(a); i++ {
			var temp map[string]interface{}
			j, _ := json.Marshal(a[i])
			json.Unmarshal(j, &temp)
			b[i] = temp
		}
	case "dc_branch":
		a, _ := GetDCBranchByParentName(name)
		for i := 0; i < len(a); i++ {
			var temp map[string]interface{}
			j, _ := json.Marshal(a[i])
			json.Unmarshal(j, &temp)
			b[i] = temp
		}
	case "convertor":
		a, _ := GetConvertorByParentName(name)
		for i := 0; i < len(a); i++ {
			var temp map[string]interface{}
			j, _ := json.Marshal(a[i])
			json.Unmarshal(j, &temp)
			b[i] = temp
		}
	default:
		return nil
	}
	return b
}
