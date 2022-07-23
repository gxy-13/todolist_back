package routers

import (
	"github.com/gin-gonic/gin"
	"tolist_back/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreatToDo)
		// 查看所有待办
		v1Group.GET("/todo", controller.GetToDoList)

		// 查看某一个
		v1Group.GET("/todo/:id", controller.GetToDo)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateToDo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteToDo)
	}
	return r
}
