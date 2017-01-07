package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type Task struct {
	Id   int    `orm:"pk" json:"id"`
	Name string `orm:"size(255)" json:"name"`
	Done bool   `json:"done"`
}

type TaskCollection struct {
	Tasks []Task `json:"iterms"`
}

func init() {
	orm.RegisterModel(new(Task))
}

func GetTasks() TaskCollection {
	sql := "select *from task"

	o := orm.NewOrm()

	var maps []orm.Params
	if _, err := o.Raw(sql).Values(&maps); err != nil {
		log.Fatalf("query task failed that error was: %s", err.Error())
	}

	result := TaskCollection{}
	for _, task := range maps {

		id, _ := strconv.Atoi(task["id"].(string))
		done, _ := strconv.ParseBool(task["done"].(string))
		data := Task{
			Id:   id,
			Name: task["name"].(string),
			Done: done,
		}
		result.Tasks = append(result.Tasks, data)
	}
	return result
}

func PostTask(name string) int64 {
	sql := `insert into task(name, done) values('` + name + `','` + "0" + `')`

	o := orm.NewOrm()

	result, err := o.Raw(sql).Exec()

	if err != nil {
		log.Fatalf("insert task into table failed that error was: %s", err.Error())
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	return id
}

func PutTask(task Task) int64 {
	var done = "0"

	if task.Done {
		done = "1"
	}

	sql := `update task set name="` + task.Name + `", done=` + done + ` where id = ` + strconv.FormatInt(int64(task.Id), 10) + ``

	o := orm.NewOrm()

	result, err := o.Raw(sql).Exec()
	if err != nil {
		log.Fatalf("update task table failed that error was: %s", err.Error())
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Fatalln(err)
	}
	return id
}

func DeleteTask(id string) {
	sql := `delete from task where id=` + id + ``

	o := orm.NewOrm()

	if _, err := o.Raw(sql).Exec(); err != nil {
		log.Fatalf("delete task from table failed that error was: %s", err.Error())
	}
}
