package db_m8

//用于更新表单中所使用线路或设备的参数ID

func UpdateACLine(name string, id int) {
	a := map[string]interface{}{
		"parameter_id": id,
	}
	if flag, _ := ExistACLineByName(name); flag == true {
		EditACLineByName(name, a)
	} else {
		return
	}
}

func UpdateDCLine(name string, id int) {
	a := map[string]interface{}{
		"parameter_id": id,
	}
	if flag, _ := ExistDCLineByName(name); flag == true {
		EditDCLineByName(name, a)
	} else {
		return
	}
}

func UpdateTrans(name string, id int) {
	a := map[string]interface{}{
		"parameter_id": id,
	}
	if flag, _ := ExistTransByName(name); flag == true {
		EditTransByName(name, a)
	} else {
		return
	}
}
