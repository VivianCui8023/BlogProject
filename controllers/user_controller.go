package controllers

import (
	"BeeProject/data"
	"BeeProject/models"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (uc *UserController) Post() {
	var userInf models.User
	err := uc.Ctx.BindJSON(&userInf)
	if err != nil {
		logs.Error(err.Error())
	}
	fmt.Printf("-------username------%s\n", userInf)

	db := data.StartDB()
	_, err = db.Insert(&userInf)
	if err != nil {
		logs.Error(err.Error())
	}

	uc.Ctx.WriteString("插入成功")
}
