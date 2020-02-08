package postgresqlXorm

import (
    "fmt"
    "log"

    "xorm.io/xorm"
    _ "github.com/lib/pq"

    "go-gin-cli/pkg/setting"
)

var Engine *xorm.Engine

var DatabaseSetting = setting.Database{
    Host: "localhost",
    Port: 5432,
    User: "jensonsu",
    Password: "bug898beet201",
    DBName: "gincli",
}

func init() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
                            DatabaseSetting.Host,
                            DatabaseSetting.Port,
                            DatabaseSetting.User,
                            DatabaseSetting.Password,
                            DatabaseSetting.DBName)

    var err error
    Engine, err = xorm.NewEngine("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err.Error())
        panic(err)
    }

    err = Engine.Ping()
    if err != nil {
        log.Fatal(err)
        return
    }

    log.Println("Successful connected to postgresql!")
}