package main

import (
	"app/models"
	"fmt"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DBName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	/*
		// ユーザ作成
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@example.com"
		u.Password = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	/*
		// ユーザ更新
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		// ユーザ削除
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/
}
