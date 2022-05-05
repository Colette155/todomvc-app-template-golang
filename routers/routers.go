package routers

import (
	"github.com/gin-gonic/gin"
	"todomvc-app-template-golang/handler"
)

func SetRouters() {
	r := gin.Default()

	group1 := r.Group("api")
	{
		group1.POST("add", handler.Add)
		group1.POST("del", handler.Del)
		group1.POST("update", handler.Update)
		group1.POST("find", handler.Find)
	}

	_ = r.Run(":9091")
}
