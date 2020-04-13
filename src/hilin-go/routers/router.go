package routers

import (
	"hilin-go/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{})

	// beego.Router("/user", &controllers.TestController{},"*:QueryUserInfo")
}
