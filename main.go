package main

import (
    "go-gin-cli/router"
)

func main() {
    // defer db.SqlDB.Close() // for postgresql

    router := router.InitRouter()
    router.Run()
}