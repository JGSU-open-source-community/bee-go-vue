package routers

import (
	"bee-go-vue/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TaskController{}, "Get:ShowIndex")
	beego.Router("/tasks", &controllers.TaskController{}, "Get:GetTasks")
}
