package db_m9

import (
	"dbPractise/models"
	"log"
	"time"
)

type M9 struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	LineID        int       `json:"line_id"`
	Type          string    `json:"type"`
	AddUser       string    `json:"add_user"`
	AddTime       time.Time `gorm:"autoUpdateTime"`
	Delete        int       `json:"delete"`
	Ready         int       `json:"ready"`
	ParameterType int       `json:"parameter_type"`
}

func Setup() {
	err := models.M9.Table("ac_branch").Association("line_id").Error
	if err != nil {
		log.Println(err)
	}
}
