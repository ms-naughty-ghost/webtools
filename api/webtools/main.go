package main

import (
	"fmt"
	"webtools/controller"
)

func main() {
	r := controller.NewRouter()
	r.Run(":80")
	fmt.Println("start")
}
