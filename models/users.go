package models

import (
    "log"
    "errors"
    
    db "go-gin-cli/database/postgresqlXorm"
)

type Users struct {
    ID       int    `json:"id" xorm:"id"`
    Name     string `json:"name" xorm:"name"`
    Phone    string `json:"phone" xorm:"phone"`
}

/*  @desc 取得用戶資料
    @return user *Users
    @return err error
*/
func (u *Users) GetUsersDetail() (users Users, err error) {
    has, err := db.Engine.Where("id = ?", u.ID).Get(&users)
    if err != nil {
        log.Println(err)
        return
    }

    if !has {
        err = errors.New("the user is not exist")
        return
    }

    return
}

/*  @desc 新增一筆users
    @return err error
*/
func (u *Users) InsertOne() (err error) {
    _, err = db.Engine.InsertOne(u)
    if err != nil {
        return
    }

    return
}