package router

import (
	"github.com/astaxie/beego"
	"goweb-minishop/controller"
)

func init() {
	beego.Router("/test", &controller.ShopController{}, "*:TestBeego")
	beego.Router("/getProduct", &controller.ShopController{}, "*:GetProduct")
}
