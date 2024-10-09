package main

import (
	"app/config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DBName)
	fmt.Println(config.Config.LogFile)
}
