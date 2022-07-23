package main

import (
	"fmt"
	"tolist_back/dao"
	"tolist_back/routers"
)

func main() {

	err := dao.InitMySQL()
	if err != nil {
		fmt.Println("数据库链接失败")
	} else {
		fmt.Println("数据库链接成功")
	}
	r := routers.SetupRouter()
	r.Run(":9000")
}
