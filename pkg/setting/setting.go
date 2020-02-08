package setting

import ()

type Database struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}
var DatabaseSetting = &Database{}


type Mongo struct {
    Host string
}
var MongoSetting = &Mongo{}