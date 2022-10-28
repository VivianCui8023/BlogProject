package routers

import (
	"BeeProject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.Router("/user", &controllers.UserController{})
}
