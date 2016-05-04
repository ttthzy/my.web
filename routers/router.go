package routers

import (
	"github.com/astaxie/beego"
	//"github.com/mikespook/gorbac"
	"my.web/controllers"
)

func init() {



	beego.Router("/", &controllers.MainController{})
}
