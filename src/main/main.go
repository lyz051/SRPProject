package main

import (
	"dbPractise/models"
	"dbPractise/setting"
)

func main() {
	setting.Setup()
	models.Setup()
	//logging.Setup()

	// Add
	//article := map[string]interface{}{
	//	"tag_id":  1,
	//	"title":   "test",
	//	"desc":    "add for testing",
	//	"content": "whatever",
	//	"state":   1,
	//}
	//util.AddArticle(article)

	//Edit
	//for i:=1;i<30;i++{
	//	article:=map[string]interface{}{
	//		"content": "modified by Edit",
	//	}
	//	util.EditArticle(i,article)
	//}

	// Delete
	//util.DeleteArticle(26)

	//Get
	//article := util.GetArticle(20)
	//log.Println("\t[GET]:\t", article.ID, "\t", article.CreatedAt)

	//Get all
	//articles := util.GetArticles(1)
	//log.Println("\t[GET ALL]")
	//for _, article := range articles {
	//	log.Println("\t", article.ID, "\t", article.CreatedAt)
	//}

	// Clean
	//util.CleanArticles()

	// Exist
	//exist := util.ExistArticleByID(25)
	//log.Println(exist)

	//b := db_m8.GiveByID(1, 0,"power_plant")
	//fmt.Println(b)

	//num := 1
	//var m int
	//var powerplant []db_m8.PowerPlant
	//for {
	//	_, err := fmt.Scan(&m)
	//	if err == io.EOF{
	//		break
	//	}else{
	//		c := db_m8.GivePlowerPlantByVoltage(220, 2, num)
	//		powerplant, _ = db_m8.GetAllPowerPlantByVoltage(220, 2, num)
	//		fmt.Println(c)
	//	}
	//	num++
	//	if num > len(powerplant){
	//		break
	//	}
	//}

	//num := 1
	//var powerplant []db_m8.PowerPlant
	//for {
	//	var m int
	//	_, err := fmt.Scanln(&m)
	//	if err == io.EOF{
	//		break
	//	}else{
	//		c := db_m8.GivePlowerPlantByVoltage(220, 2, num)
	//		powerplant, _ = db_m8.GetAllPowerPlantByVoltage(220, 2, num)
	//		fmt.Println(c)
	//	}
	//	num++
	//	if num > len(powerplant){
	//		break
	//	}
	//}

}
