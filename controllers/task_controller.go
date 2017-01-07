package controllers

import (
	"encoding/json"

	"bee-go-vue/models"

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

	task := t.bind()
	models.PostTask(task.Name)
	t.Redirect("/", 302)
	return
}

func (t *TaskController) PutTask() {
	task := t.bind()

	models.PutTask(task)
	t.Redirect("/", 302)
	return

}

func (t *TaskController) DeleteTask() {
	id := t.Ctx.Input.Param(":id")
	models.DeleteTask(id)

	t.Redirect("/", 302)
	return
}

func (t *TaskController) bind() (ta models.Task) {
	json.NewDecoder(t.Ctx.Request.Body).Decode(&ta)
	return
}
