package setting

import (
    "log"
    "time"

    "github.com/go-ini/ini"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}


type Database struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
}
var DatabaseSetting = &Database{}


type Mongo struct {
    Host string
}
var MongoSetting = &Mongo{}


var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
    var err error
    cfg, err = ini.Load("conf/app.ini")
    if err != nil {
        log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
    }

    mapTo("server", ServerSetting)
    mapTo("database", DatabaseSetting)
    mapTo("mongo", MongoSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
    err := cfg.Section(section).MapTo(v)
    if err != nil {
        log.Fatalf("Cfg.MapTo %s err: %v", section, err)
    }
}