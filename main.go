package main

import (
	"github.com/astaxie/beego"
	"github.com/mikespook/gorbac"
	_ "my.web/routers"
)

func main() {

	// 初始化角色权限
	rbac := gorbac.New()
	r1 := gorbac.NewStdRole("admin") //管理员角色

	pis := make(gorbac.Permissions)
	pis["All"] = gorbac.NewStdPermission("All")           //拥有所有权限
	pis["Create"] = gorbac.NewStdPermission("Create")     //拥有c权限
	pis["Retrieve"] = gorbac.NewStdPermission("Retrieve") //拥有r权限
	pis["Update"] = gorbac.NewStdPermission("Update")     //拥有u权限
	pis["Delete"] = gorbac.NewStdPermission("Delete")     //拥有d权限

	for _, v := range pis {
		r1.Assign(v)
	}

	rbac.Add(r1)
	
	// 启动beego
	beego.Run()
}
