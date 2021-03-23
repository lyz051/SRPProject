package db_m2

type M2 struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
