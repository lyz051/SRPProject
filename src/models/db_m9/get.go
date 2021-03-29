package db_m9

import (
	"dbPractise/models"
	"gorm.io/gorm"
	"reflect"
)

//func (b *ACBranch) getDyanmicLine() (*db_m2.DynamicLine, error) {
//	var dynamicLine *db_m2.DynamicLine
//	err := models.M9.Model(&ACBranch{}).Where("id = ?", b.ID).Find(b).Error
//	if err != nil {
//		return nil, err
//	}
//	err = models.M9.Model(&ACBranch{}).Preload("id = ?", b.ID).Error
//
//	return dynamicLine, nil
//}

func ExistByID(table string, id int) (bool, error) {
	//var tableStruct interface{}
	//switch table {
	//case "ac_branch":
	//	tableStruct=tableStruct.(ACBranch)
	//case "transformer":
	//	tableStruct=tableStruct.(Transformer)
	//case "dc_branch":
	//	tableStruct=tableStruct.(DCBranch)
	//case "convertor":
	//	tableStruct=tableStruct.(Convertor)
	//}
	var res map[string]interface{}
	// FIXME: 无法指定表单
	//err := models.M9.Table(table).Where("id = ?", id).Take(&res).Error
	err := models.M9.Table(table).First(&res, "id = ?", id).Error
	//err:=models.M9.Raw("SELECT id, type FROM "+table+" WHERE id = ?", id).Scan(&res).Error
	if err != nil || err != gorm.ErrRecordNotFound {
		return false, err
	}

	if reflect.TypeOf(res) != reflect.TypeOf(reflect.Map) {
		return true, nil
	}

	return false, nil
}

func GetByID(table string, id int) (interface{}, error) {
	var tableStruct interface{}
	switch table {
	case "ac_branch":
		tableStruct = tableStruct.(ACBranch)
	case "transformer":
		tableStruct = tableStruct.(Transformer)
	case "dc_branch":
		tableStruct = tableStruct.(DCBranch)
	case "convertor":
		tableStruct = tableStruct.(Convertor)
	}

	err := models.M8.Where("id = ? and delete != 1", id).First(&tableStruct).Error
	if err != nil {
		return nil, err
	}
	return tableStruct, nil
}

func ExistACBranchByInstance(ins ACBranch) (bool, error) {
	var acbID []int
	err := models.M9.Model(&ACBranch{}).Select("id").Where("`delete` != ?", 1).
		Where(map[string]interface{}{
			"r":              ins.R,
			"rpu":            ins.RPU,
			"rate_a":         ins.RateA,
			"x":              ins.X,
			"x0":             ins.X0,
			"xpu":            ins.XPU,
			"i_rated":        ins.IRated,
			"parallel_num":   ins.ParallelNum,
			"G_2":            ins.G2,
			"B_2":            ins.B2,
			"length":         ins.Length,
			"parameter_type": ins.ParameterType,
		}).Pluck("id", &acbID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(acbID) > 0 {
		return true, nil
	}

	return false, nil
}

func ExistTransformerByInstance(ins Transformer) (bool, error) {
	var transID []int
	err := models.M9.Model(&Transformer{}).Select("id").Where("`delete` != ?", 1).
		Where(map[string]interface{}{
			"cdbz":   ins.CDBZ,
			"fdcldh": ins.FDCLDH,
			"s":      ins.S,
			"blbsm":  ins.BLBSM,
			"r_t":    ins.RT,
			"x":      ins.X,
			"tg":     ins.Tg,
			"jg":     ins.Jg,
			"fjt1":   ins.FJT1,
			"fjt2":   ins.FJT2,
		}).Pluck("id", &transID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(transID) > 0 {
		return true, nil
	}

	return false, nil
}

func ExistDCBranchByInstance(ins DCBranch) (bool, error) {
	var dcbID []int
	err := models.M9.Model(&DCBranch{}).Select("id").Where("`delete` != ?", 1).
		Where(map[string]interface{}{
			"cdbz":           ins.CDBZ,
			"i_base":         ins.IBase,
			"r":              ins.R,
			"x":              ins.X,
			"c":              ins.C,
			"glkzd":          ins.GLKZD,
			"w":              ins.W,
			"zlvolt":         ins.ZLVOLT,
			"Zlczccfj":       ins.ZLCZCCFJ,
			"nbczcgdj":       ins.NBCZCGDJ,
			"length":         ins.Length,
			"parameter_type": ins.ParameterType,
		}).Pluck("id", &dcbID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(dcbID) > 0 {
		return true, nil
	}

	return false, nil
}

func ExistConvertorhByInstance(ins Convertor) (bool, error) {
	var cvtID []int
	err := models.M9.Model(&Convertor{}).Select("id").Where("`delete` != ?", 1).
		Where(map[string]interface{}{
			"fqm":            ins.FQM,
			"zlxqs":          ins.ZLXQS,
			"pbdkdg":         ins.PBDKDG,
			"zxcfj":          ins.ZXCFJ,
			"zdcfj":          ins.ZDCFJ,
			"mgqfdyj":        ins.MGQFDYJ,
			"qdledz":         ins.QDLEDZ,
			"parameter_type": ins.ParameterType,
		}).Pluck("id", &cvtID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(cvtID) > 0 {
		return true, nil
	}

	return false, nil
}

func ExistACBranchByID(id int) (bool, error) {
	var acbID []int
	err := models.M9.Model(&ACBranch{}).Where("id = ?", id).Pluck("id", &acbID).Error
	//err:=models.M9.Raw("SELECT id, type FROM "+table+" WHERE id = ?", id).Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(acbID) > 0 {
		return true, nil
	}

	return false, nil
}

func ExistTransformerByID(id int) (bool, error) {
	var transID []int
	err := models.M9.Model(&Transformer{}).Where("id = ?", id).Pluck("id", &transID).Error
	//err:=models.M9.Raw("SELECT id, type FROM "+table+" WHERE id = ?", id).Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(transID) > 0 {
		return true, nil
	}

	return false, nil
}

func ExistDCBranchByID(id int) (bool, error) {
	var dcbID []int
	err := models.M9.Model(&DCBranch{}).Where("id = ?", id).Pluck("id", &dcbID).Error
	//err:=models.M9.Raw("SELECT id, type FROM "+table+" WHERE id = ?", id).Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(dcbID) > 0 {
		return true, nil
	}

	return false, nil
}

func ExistConvertorByID(id int) (bool, error) {
	var cvtID []int
	err := models.M9.Model(&Convertor{}).Where("id = ?", id).Pluck("id", &cvtID).Error
	//err:=models.M9.Raw("SELECT id, type FROM "+table+" WHERE id = ?", id).Scan(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(cvtID) > 0 {
		return true, nil
	}

	return false, nil
}

func GetACBranchByID(id int) (*ACBranch, error) {
	var acb ACBranch
	err := models.M9.Model(&ACBranch{}).Where("id = ? and `delete` != 1", id).First(&acb).Error
	if err != nil {
		return nil, err
	}
	return &acb, nil
}

func GetTransformerByID(id int) (*Transformer, error) {
	var trans Transformer
	err := models.M9.Model(&Transformer{}).Where("id = ? and `delete` != 1", id).First(&trans).Error
	if err != nil {
		return nil, err
	}
	return &trans, nil
}

func GetDCBranchByID(id int) (*DCBranch, error) {
	var dcb DCBranch
	err := models.M9.Model(&DCBranch{}).Where("id = ? and `delete` != 1", id).First(&dcb).Error
	if err != nil {
		return nil, err
	}
	return &dcb, nil
}

func GetConvertorByID(id int) (*Convertor, error) {
	var cvt Convertor
	err := models.M9.Model(&Convertor{}).Where("id = ? and `delete` != 1", id).First(&cvt).Error
	if err != nil {
		return nil, err
	}
	return &cvt, nil
}

func GetACBranchByParentName(name string) ([]ACBranch, error) {
	var acb []ACBranch
	err := models.M9.Model(&ACBranch{}).Where("line_name = ? and `delete` != 1", name).Find(&acb).Error
	if err != nil {
		return nil, err
	}
	return acb, nil
}

func GetTransformerByParentName(name string) ([]Transformer, error) {
	var trans []Transformer
	err := models.M9.Model(&Transformer{}).Where("trans_name = ? and `delete` != 1", name).Find(&trans).Error
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func GetDCBranchByParentName(name string) ([]DCBranch, error) {
	var dcb []DCBranch
	err := models.M9.Model(&DCBranch{}).Where("line_name = ? and `delete` != 1", name).Find(&dcb).Error
	if err != nil {
		return nil, err
	}
	return dcb, nil
}

func GetConvertorByParentName(name string) ([]Convertor, error) {
	var cvt []Convertor
	err := models.M9.Model(&Convertor{}).Where("trans_name = ? and `delete` != 1", name).Find(&cvt).Error
	if err != nil {
		return nil, err
	}
	return cvt, nil
}
