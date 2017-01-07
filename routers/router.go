package routers

import (
	"bee-go-vue/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TaskController{}, "Get:ShowIndex")
	beego.Router("/tasks", &controllers.TaskController{}, "Get:GetTasks")
	beego.Router("/task", &controllers.TaskController{}, "Post:PostTask")
	beego.Router("/task", &controllers.TaskController{}, "Put:PutTask")
	beego.Router("/task/:id", &controllers.TaskController{}, "Delete:DeleteTask")

}
