package router

import (
    "github.com/gin-gonic/gin"

    "go-gin-cli/router/api"
    "go-gin-cli/router/api/v1"
    "go-gin-cli/middleware"
)

func InitRouter() *gin.Engine {
    r := gin.Default()

    apiBase := r.Group("/api")
    {
        apiBase.POST("/login", api.LoginAuthApi)
    }

    apiv1 := r.Group("/api/v1", middleware.AuthJwtMiddle())
    {
        users := apiv1.Group("/users")
        {
            users.GET("/detail", v1.GetDetailUsersApi)
        }
    }
    
    return r
}