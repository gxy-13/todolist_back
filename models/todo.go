package models

import (
	"fmt"
	"tolist_back/dao"
)

// Todo Model
type Todo struct {
	Id     int    `json:"id"`
	Tittle string `json:"title"`
	Status bool   `json:"status"`
}

// Todo CRUD
func CreateToDo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	if err != nil {
		fmt.Printf("数据插入失败，%#v\n", err.Error())
		return err
	}
	fmt.Printf("todolist添加成功，%#v\n", todo)
	return nil
}

func GetAllToDo(todos *[]Todo) (err error) {
	err = dao.DB.Debug().Find(&todos).Error
	if err != nil {
		fmt.Println("获取数据失败", err.Error())
		return err
	} else {
		fmt.Println("获取列表成功")
		return nil
	}
}

func SelectToDoById(todo *Todo, id string) (err error) {
	err = dao.DB.Debug().Select("id = ?", id).Find(&todo).Error
	if err != nil {
		fmt.Printf("没有%s的数据", id)
		return err
	}
	return nil
}

func UpdateToDo(todo *Todo, id string) (err error) {
	dao.DB.Debug().Find(&todo, id)
	if todo.Status == true {
		err = dao.DB.Debug().Model(&Todo{}).Where("id = ?", id).Update("status", 0).Error
		//err = dao.DB.Debug().Update("status", 1).Find(&todo).Error
	} else {
		err = dao.DB.Debug().Model(&Todo{}).Where("id = ?", id).Update("status", 1).Error
		//err = dao.DB.Debug().Update("status", 0).Find(&todo).Error
	}
	if err != nil {
		fmt.Printf("数据%s修改失败", id)
		return err
	} else {
		fmt.Printf("数据%s修改成功", id)
		return nil
	}
}

func DeleteById(todo *Todo, id string) (err error) {
	err = dao.DB.Where("id = ", id).Delete(&todo).Error
	if err != nil {
		fmt.Println("删除失败")
		return err
	} else {
		fmt.Println("删除成功")
		return nil
	}
}
