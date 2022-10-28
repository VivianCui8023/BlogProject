package routers

import (
	"BeeProject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//这个路由可以匹配/getImage/xxx/xx/xxx/xxx.txt类似路径
	beego.Router("/getImage/*.*", &controllers.MainController{})
	//这种声明匹配id是int类型
	beego.Router("/getUserInfo/:id:int", &controllers.MainController{})
	//自定义路由方法
	beego.Router("/register", &controllers.MainController{}, "Get:GetUserInfo")
}
