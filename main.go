package main

import (
	"fmt"
	"go-todos/config"
)

func main() {
	conf := config.GetConfig()
	fmt.Println(conf)
}
