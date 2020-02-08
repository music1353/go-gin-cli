package v1

import (
    _"log"
    "net/http"

    "github.com/gin-gonic/gin"

    "go-gin-cli/models"
)

/*  @desc 取得用戶資料
    @method GET
    @router /api/v1/users/detail
    @result {id, name, phone}
*/
func GetDetailUsersApi(c *gin.Context) {
    id := c.MustGet("id")

    users := models.Users{ID: id.(int)}
    users, err := users.GetUsersDetail()
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{
            "result": "",
            "msg": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "result": users,
        "msg": "",
    })
}