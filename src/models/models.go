package models

import (
	"dbPractise/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error

	// driver: "gorm.io/driver/mysql"
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.DatabaseSetting.TablePrefix, // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                                // use singular table name, table for `User` would be `user` with this option enabled
			//NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		NowFunc: func() time.Time {
			return time.Now()
		},
	})

	// driver "github.com/jinzhu/gorm/dialects/mysql"
	//db, err = gorm.Open(mysql.New(mysql.Config{DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//	setting.DatabaseSetting.User,
	//	setting.DatabaseSetting.Password,
	//	setting.DatabaseSetting.Host,
	//	setting.DatabaseSetting.Name)}), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		TablePrefix:   setting.DatabaseSetting.TablePrefix, // table name prefix, table for `User` would be `t_users`
	//		SingularTable: true,                                // use singular table name, table for `User` would be `user` with this option enabled
	//		//NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
	//	},
	//	NowFunc: func() time.Time {
	//		return time.Now()
	//	},
	//})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}
