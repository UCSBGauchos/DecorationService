package routers

import (
	"wwwin-github.cisco.com/DevNet/DecorationService/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
