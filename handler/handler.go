package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todomvc-app-template-golang/model"

	"net/http"
)

func Add(c *gin.Context) {
	var p model.Todo
	c.ShouldBindJSON(&p) //从前端拿到json格式的数据（item）,并放入一个model
	fmt.Println(p)
	model.AddAThing(&p) //将这个model存入数据库
	c.JSON(http.StatusOK, nil)
}

func Find(c *gin.Context) { //用户要么根据status查找，要么根据item查找
	var p model.Todo
	c.ShouldBindJSON(&p) //从前端拿到json格式的数据（item和status）,并放入一个model
	q := model.FindAThing(&p)
	c.JSON(http.StatusOK, &q) //讲找到的模型渲染回前端
}

func Update(c *gin.Context) {
	var p []model.Todo
	c.ShouldBindJSON(&p) //从前端拿到json格式的数据(可能不止一个),并放入一个model切片
	fmt.Println(p)
	for _, t := range p { //因为Model可能不止一个所以要遍历从前端拿到的数据
		model.UpdateAThing(&t) //遍历一个，更改一个
	}

	c.JSON(http.StatusOK, nil)
}

func Del(c *gin.Context) {
	var p model.Todo
	c.ShouldBindJSON(&p) //从前端拿到json格式的数据,并放入一个model
	fmt.Println(p)
	model.DelAThing(p.ID) //根据数据中的id在数据库中找到并删除
	c.JSON(http.StatusOK, nil)
}
