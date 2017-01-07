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

func (t *TaskController) URLMapping() {
	t.Mapping("Get", t.GetTasks)
	t.Mapping("Post", t.PostTask)
	t.Mapping("Put", t.PutTask)
	t.Mapping("Delete", t.Delete)
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
	id := models.PostTask(task.Name)

	t.Data["json"] = H{
		"created": id,
	}
	t.ServeJSON()
}

func (t *TaskController) PutTask() {
	task := t.bind()

	id := models.PutTask(task)

	t.Data["json"] = H{
		"updated": id,
	}

	t.ServeJSON()
}

func (t *TaskController) DeleteTask() {
	id := t.Ctx.Input.Param(":id")
	models.DeleteTask(id)

	t.Data["json"] = H{
		"deleted": id,
	}
	t.ServeJSON()
}

func (t *TaskController) bind() (ta models.Task) {
	json.NewDecoder(t.Ctx.Request.Body).Decode(&ta)
	return
}
