package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mikespook/gorbac"
)

type MainController struct {
	beego.Controller
	rbac *gorbac.RBAC
}

func (c *MainController) Get() {

	// role, permission, err := c.rbac.Get("admin")

	// gorbac.StdRole.Permissions["All"]
	// p := &gorbac.StdRole.Permissions
	// c.Data["Website"] = p
	// if c.rbac.IsGranted("admin", p3, nil) {
	// 	c.Data["Website"] = p3.ID
	// }

	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
