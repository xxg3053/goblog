package main

import (
	"fmt"
	"github.com/xxg3053/goblog/accountservice/service"
)

var appName = "accountService"

func main()  {
	fmt.Printf("Starting %v\n", appName)
	service.StartwebServer("6767")
}
