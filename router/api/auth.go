package api

import (
    _"log"
    "net/http"

    "github.com/gin-gonic/gin"

    "go-gin-cli/models"
)

type Body struct {
    Account  string `json:"account" form:"account"`
    Password string `json:"password" form:"password"`
}

/*  @desc 登入
    @method POST
    @router /api/login
    @param account string
    @param password string
    @result token string
*/
func LoginAuthApi(c *gin.Context) {
    var body Body
    c.ShouldBindJSON(&body)

    // 檢查用戶是否存在
    var auth models.Auth
    isExist, auth, err := auth.CheckAuth(body.Account, body.Password)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    if isExist == false {
        c.JSON(http.StatusUnauthorized, gin.H{
            "result": "",
            "msg": "無此用戶",
        })
        return
    }

    // 用戶存在，取得用戶資料
    // auth, err := auth.GetAuthDetail()
    // if err != nil {
    //     c.JSON(http.StatusServiceUnavailable, gin.H{
    //         "result": "",
    //         "msg": err.Error(),
    //     })
    //     return
    // }

    // 發放token
    token, err := auth.IssueToken(auth.ID, auth.Account, auth.Role)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": gin.H{
            "token": token,
        },
        "msg": "",
    })
}