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

	/*
		// TODO作成
		user, _ := models.GetUser(1)
		fmt.Println(user)

		user.CreateTodo("First Todo")
	*/

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	/*
		// TODO一覧取得
		user, _ := models.GetUser(2)
		user.CreateTodo("Second Todo")

		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		// 該当ユーザのTODO一覧取得
		u := &models.User{}
		u.Name = "test2"
		u.Email = "test2@example.com"
		u.Password = "testtest2"
		fmt.Println(u)

		u.CreateUser()

		user, _ := models.GetUser(3)
		user.CreateTodo("Third Todo")

		user2, _ := models.GetUser(2)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		// TODO更新
		todo, _ := models.GetTodo(1)
		todo.Content = "update todo"
		todo.UpdateTodo()
	*/

	// TODO削除
	todo, _ := models.GetTodo(3)
	todo.DeleteTodo()
}
