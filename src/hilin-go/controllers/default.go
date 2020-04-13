package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "hilin.com"
	c.Data["Email"] = "hilin@qq.com"
	c.TplName = "index.tpl"
}
