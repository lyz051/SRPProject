package setting

import (
	"github.com/go-ini/ini"
	"log"
)

type Log struct {
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	RuntimeRootPath string
	TimeFormat      string
}

type DBNames struct {
	M2       string
	M3       string
	M4       string
	M5       string
	M6       string
	M8       string
	M9       string
	M10      string
	M11      string
	ScadaBpa string
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	TablePrefix string
	PageSize    int
}

var (
	LogSetting      = &Log{}
	DBNamesSetting  = &DBNames{}
	DatabaseSetting = &Database{}
)

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("./../config/app.ini")
	//cfg, err = ini.Load("./config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/app.ini': %v", err)
	}

	mapTo("log", LogSetting)
	mapTo("dbNames", DBNamesSetting)
	mapTo("database", DatabaseSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
