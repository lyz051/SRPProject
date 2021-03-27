package models

import (
	"dbPractise/pkg/logging"
	"dbPractise/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"reflect"
	"time"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	M2   *gorm.DB
	M8   *gorm.DB
	M9   *gorm.DB
	M10  *gorm.DB
	M11  *gorm.DB
	STOB *gorm.DB
)
var pageSize int

// Setup initializes the database instance
func Setup() {
	var err error

	pageSize = setting.DatabaseSetting.PageSize

	M2, err = open(setting.DBNamesSetting.M2)
	M8, err = open(setting.DBNamesSetting.M8)
	M9, err = open(setting.DBNamesSetting.M9)
	M10, err = open(setting.DBNamesSetting.M10)
	M11, err = open(setting.DBNamesSetting.M11)
	STOB, err = open(setting.DBNamesSetting.ScadaBpa)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}

// driver: "gorm.io/driver/mysql"
func open(name string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host, name)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.DatabaseSetting.TablePrefix, // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                                // use singular table name, table for `User` would be `user` with this option enabled
			//NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		Logger: logger.New(
			log.New(logging.F, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: 500 * time.Millisecond,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
		//NowFunc: func() time.Time {
		//	return time.Now()
		//},
	})
	return db, err
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		switch db.Statement.ReflectValue.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				rv := reflect.Indirect(db.Statement.ReflectValue.Index(i))
				field1 := db.Statement.Schema.FieldsByDBName["updated_at"]
				field1.Set(rv, currentTime)
				field := db.Statement.Schema.FieldsByDBName["created_at"]
				field.Set(rv, currentTime)
			}
		case reflect.Struct:
			db.Statement.SetColumn("created_at", currentTime)
		}
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	// if _, ok := db.Statement.Settings.Load("gorm:update_time_stamp"); ok {
	if db.Statement.Schema != nil {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		db.Statement.SetColumn("UpdatedAt", currentTime)
		db.Statement.AddClause(clause.Where{
			Exprs: []clause.Expression{clause.Eq{Column: "deleted_at"}},
		})
	}
}
