package postgresql

import (
    "database/sql"
    "fmt"
    "log"

    "go-gin-cli/pkg/setting"
)

var SqlDB *sql.DB

func Setup() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
                            setting.DatabaseSetting.Host,
                            setting.DatabaseSetting.Port,
                            setting.DatabaseSetting.User,
                            setting.DatabaseSetting.Password,
                            setting.DatabaseSetting.DBName)

    var err error
    SqlDB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err.Error())
        panic(err)
    }

    err = SqlDB.Ping()
    if err != nil {
        log.Fatal(err.Error())
        panic(err)
    }

    log.Println("Successful connected to database!")
}