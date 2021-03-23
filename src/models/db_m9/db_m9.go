package db_m9

import (
	"time"
)

type M9 struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	LineID        int       `json:"line_id" gorm:"index"`
	Type          string    `json:"type"`
	AddUser       string    `json:"add_user"`
	AddTime       time.Time `gorm:"autoUpdateTime"`
	Delete        int       `json:"delete"`
	Ready         int       `json:"ready"`
	ParameterType int       `json:"parameter_type"`
}

type ACBranch struct {
	M9
	//DynamicLine db_m2.DynamicLine `json:"dynamic_line"`
	R           float32 `json:"r"`
	RPU         float32 `json:"rpu"`
	RateA       float32 `json:"rate_a"`
	X           float32 `json:"x"`
	X0          float32 `json:"x0"`
	XPU         float32 `json:"xpu"`
	IRated      float32 `json:"i_rated"`
	ParallelNum int     `json:"parallel_num"`
	G2          float32 `gorm:"column:G_2"`
	B2          float32 `gorm:"column:B_2"`
	Length      float32 `json:"length"`
}

func (this *ACBranch) TableName() string {
	return "ac_branch"
}

// 建立关联
//func Setup() {
//err := models.M9.Table("ac_branch").Association("line_id").Error
//if err != nil {
//	log.Println(err)
//}
//}
