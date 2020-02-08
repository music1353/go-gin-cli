package main

import (
    "github.com/gin-gonic/gin"
    "go-gin-cli/router"

    "go-gin-cli/pkg/setting"
    db "go-gin-cli/database/postgresqlXorm"
)

func init() {
    setting.Setup()
    db.Setup()
}

func main() {
    // defer db.SqlDB.Close() // for postgresql
    gin.SetMode(setting.ServerSetting.RunMode)

    router := router.InitRouter()
    router.Run()
}