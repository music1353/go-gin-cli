package models

import (
    "log"

    "go-gin-cli/middleware"
    db "go-gin-cli/database/postgresqlXorm"
)

type Auth struct {
    ID       int    `json:"id" xorm:"id"`
    Account  string `json:"account" form:"account" xorm:"account"`
    Password string `json:"password" form:"password" xorm:"password"`
    Role     string `json:"role" form:"role" xorm:"role"`
}

/*  @desc 是否存在此用戶, 並返回資料
    @param account string
    @param password string
    @return isExist bool
    @return auth Auth
    @return err error
*/
func (a *Auth) CheckAuth(account, password string) (isExist bool, auth Auth, err error) {
    isExist = false

    has, err := db.Engine.Where("account = ? AND password = ?", account, password).Get(&auth)
    if err != nil {
        return
    }

    if has {
        isExist = true
        err = nil
    }

    return
}

/*  @desc 發放jwt token
    @param account string
    @param password string
    @return token string
    @return err error
*/
func (a *Auth) IssueToken(id int, account string, role string) (token string, err error) {
    token, err = middleware.GenerateToken(id, account, role)
    if err != nil {
        log.Println(err)
        return
    }
    return
}