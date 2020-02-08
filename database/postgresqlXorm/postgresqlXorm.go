package postgresqlXorm

import (
    "fmt"
    "log"

    "xorm.io/xorm"
    _ "github.com/lib/pq"

    "go-gin-cli/pkg/setting"
)

var Engine *xorm.Engine

func Setup() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
                            setting.DatabaseSetting.Host,
                            setting.DatabaseSetting.Port,
                            setting.DatabaseSetting.User,
                            setting.DatabaseSetting.Password,
                            setting.DatabaseSetting.Name)

    var err error
    Engine, err = xorm.NewEngine("postgres", psqlInfo)
    if err != nil {
        log.Print("here")
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