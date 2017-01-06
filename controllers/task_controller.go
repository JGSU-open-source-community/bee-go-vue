package controllers

import (
	"bee-go-vue/models"
	"fmt"

	"github.com/astaxie/beego"
)

type H map[string]interface{}

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
	name := t.GetString("name")

	fmt.Println("================test=================")
	fmt.Println(name)
}
