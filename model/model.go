package model

import (
	"fmt"
	"todomvc-app-template-golang/db"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Item   string `json:"item"`
	Status int    `json:"status"`
}

//model的增删改查

func AddAThing(todo *Todo) {
	var m = Todo{Item: todo.Item, Status: 0} //拿到一个模型
	fmt.Println(m)
	db.DB.Model(&Todo{}).Create(&m) //将这个模型存入数据库
}

func DelAThing(id uint) { //拿到一个id
	db.DB.Model(&Todo{}).Where("ID", id).Delete(&Todo{}) //根据这个id在数据库中找到并删除相应的内容
}

func UpdateAThing(todo *Todo) { //select通过where找到的数据的Item,status和update_time,然后更改为传入的数据
	db.DB.Model(&Todo{}).Where("ID", todo.ID).Select("Item", "Status", "update_at").Updates(Todo{
		Item:   todo.Item,
		Status: todo.Status,
	})

}

func FindAThing(todo *Todo) (k *[]Todo) {
	var p []Todo
	var t = db.DB
	if todo.Status != -1 {
		t = db.DB.Model(&Todo{}).Where("Status", todo.Status) //这一句必执行，即查找出所有数据
	}
	if todo.Item != "" { //这一句类似于筛选条件，即范围缩小到用户输入的，也就是从前端拿到的item
		t = db.DB.Model(&Todo{}).Where("item LIKE ?", "%"+todo.Item+"%")
	}
	t.Find(&p) //将查找的结果(可能不止一个)，放入model切片
	return &p  //返回model切片用于前端渲染
}
