package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"hilin-go/models"
)

var user models.User

type TestController struct {
	beego.Controller
}

func (c *TestController) QueryUserInfo()  {
	username := c.GetString("username")
	password := c.GetString("password")

	user, err := user.QueryUserOneByNameAndPwd(username, password)
	fmt.Println(user, err)
}