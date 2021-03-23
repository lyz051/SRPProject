package db_m9

import (
	"time"
)

type M9 struct {
	ID            int `gorm:"primaryKey"`
	Type          string
	AddUser       string
	AddTime       time.Time `gorm:"autoUpdateTime"`
	Delete        int
	Ready         int
	ParameterType int
}

type ACBranch struct {
	M9

	LineID int `gorm:"index"`
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

type Transformer struct {
	M9

	TransID int     `gorm:"index"`
	CDBZ    float32 `gorm:"column:cdbz"`   // 区域联络线测点标志
	FDCLDH  float32 `gorm:"column:fdcldh"` // 分段串联短号
	S       float32 // 额定容量
	BLBSM   int     `gorm:"column:blbsm"` // 并联变数目
	RT      float32 `gorm:"column:r_t"`   // 铜损等效电阻
	X       float32 // 漏抗
	Tg      float32 `gorm:"column:tg"`   // 铁损等效电导
	Jg      float32 `gorm:"column:jg"`   // 激磁导纳
	FJT1    float32 `gorm:"column:fjt1"` // 节点1固定分接头
	FJT2    float32 `gorm:"column:fjt2"` // 节点2固定分接头
}

type DCBranch struct {
	M9

	LineID   int     `gorm:"index"`
	CDBZ     float32 `gorm:"column:cdbz"` // 区域联络线测点标志
	IBase    float32 // 直流线路额定电流
	R        float32 // 直流线路电阻
	X        float32 // 直流电路电感
	C        float32 // 直流线路电容
	GLKZD    string  `gorm:"column:glkzd"` // 安排直流功率控制点
	W        float32 // 安排直流功率
	ZLVOLT   float32 `gorm:"column:zlvolt"`   // 给定整流侧直流电压
	zlczccfj float32 `gorm:"column:zlczccfj"` // 整流侧政策触发角
	NBCZCGDJ float32 `gorm:"column:nbczcgdj"` // 逆变侧正常关断角
	Length   float32 `json:"length"`
}

func (this *DCBranch) TableName() string {
	return "dc_branch"
}

type Convertor struct {
	M9

	TransID string  `gorm:"index"`
	FQM     float32 `gorm:"column:fqm"`    // 分区名
	ZLXQS   float32 `gorm:"column:zlxqs"`  // 整流线桥数
	PBDKDG  float32 `gorm:"column:pbdkdg"` // 平波电抗电感
	ZXCFJ   float32 `gorm:"column:zxcfj"`  // 最小触发角
	ZDCFJ   float32 `gorm:"column:zdcfj"`  // 最大触发角
	MGQFDYJ float32 `gorm:"mgqfdyj"`       // 每个桥阀电压降
	QDLEDZ  float32 `gorm:"qdledz"`        // 桥电流额定值
}
