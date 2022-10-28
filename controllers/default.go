package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	path := c.Ctx.Input.Param(":path")
	fmt.Println(path)
	//不能忘记加:，否则不能提取出参数
	ext := c.Ctx.Input.Param(":ext")
	fmt.Println(ext)
	c.Ctx.ResponseWriter.Write([]byte("filepath:" + path + "---ext:" + ext))
}
