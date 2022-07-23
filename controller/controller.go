package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tolist_back/models"
)

func CreatToDo(c *gin.Context) {
	// 从请求中获取数据
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateToDo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code":    200,
		//	"message": "ok",
		//	"data":    todo,
		//})
	}

}

func GetToDoList(c *gin.Context) {
	// 定义一个接收所有todo的slice
	todos := make([]models.Todo, 0, 100)
	// 从数据库查找所有的todo
	err := models.GetAllToDo(&todos)
	// 放回给前端json格式的所有todo
	for _, v := range todos {
		fmt.Printf("%#v\n", v)
	}
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, todos)
	}

}

func GetToDo(c *gin.Context) {
	// 获取uri上的id
	var todo models.Todo
	id := c.Param("id")
	err := models.SelectToDoById(&todo, id)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateToDo(c *gin.Context) {
	// 获取uri上的id
	var todo models.Todo
	id := c.Param("id")
	err := models.UpdateToDo(&todo, id)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func DeleteToDo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	err := models.DeleteById(&todo, id)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
