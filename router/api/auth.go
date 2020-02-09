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

type SignUp struct {
    Account  string `json:"account"`
    Password string `json:"password"`
    Role     string `json:"role"`
    Name     string `json:"name"`
    Phone    string `json:"phone"`
}

/*  @desc 註冊
    @method POST
    @router /api/signup
    @param account string
    @param password string
    @param name string
    @param phone string
*/
func SignUpAuthApi(c *gin.Context) {
    var signup SignUp
    c.ShouldBindJSON(&signup)

    // 檢查用戶是否存在
    var auth models.Auth
    isExist, err := auth.CheckAccount(signup.Account)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    // 若帳號存在，不給註冊
    if isExist {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": "此帳號已存在",
        })
        return
    }

    // insert auth
    auth = models.Auth{
        Account: signup.Account,
        Password: signup.Password,
        Role: signup.Role,
    }
    id, err := auth.InsertOne()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    // insert users
    users := models.Users{
        ID: id,
        Name: signup.Name,
        Phone: signup.Phone,
    }
    err = users.InsertOne()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": "",
        "msg": "註冊成功",
    })
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