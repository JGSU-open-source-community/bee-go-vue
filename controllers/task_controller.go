package controllers

import (
	"bee-go-vue/models"

	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

func (t *TaskController) ShowIndex() {
	t.TplName = "index.tpl"
}

func (t *TaskController) GetTasks() {
	datas := models.GetTasks()
	t.Data["json"] = datas
	t.ServeJSON()
}

func (t *TaskController) PostTask() {

}
