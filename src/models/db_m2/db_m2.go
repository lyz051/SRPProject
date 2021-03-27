package db_m2

type DynamicPlant struct {
	ID   int    `json:"id" gorm:"primaryKey"` //厂站ID
	Name string `json:"name" gorm:"index"`    //厂站英文名
}

type DynamicBus struct {
	ID   int    `json:"id" gorm:"primaryKey"` //母线ID
	Name string `json:"name"`                 //母线英文名
}

type DynamicLine struct {
	ID        int    `json:"id" gorm:"primaryKey"` //线路ID
	Name      string `json:"name"`                 //线路英文名
	Linename  string `json:"linename"`             //线路名称
	Type      int    `json:"type"`                 //线路类型
	STID_A    int    `gorm:"column:stid_a"`        //A端厂站
	STID_B    int    `gorm:"column:stid_b"`        //B端厂站
	BusName_A string `gorm:"column:bus_name_a"`    //A端母线名称
	BusName_B string `gorm:"column:bus_name_b"`    //B端母线名称
}

type DynamicTrans struct {
	ID        int    `json:"id" gorm:"primaryKey"` //变压器ID
	Name      string `json:"name"`                 //变压器英文名
	Type      string `json:"type"`                 //卡片类型
	TransName string `json:"trans_name"`           //变压器名称
	BelongST  string `gorm:"column:belong_st"`     //所属厂站名称
	BusName_1 string `gorm:"column:bus_name_1"`    //1端母线名称
	BusName_2 string `gorm:"column:bus_name_2"`    //2端母线名称
}
