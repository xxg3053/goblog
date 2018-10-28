package main

import (
	"fmt"
	"github.com/xxg3053/goblog/dashboard/service"
)

var appName = "dashboard"

func main()  {
	fmt.Printf("Starting %v\n", appName)
	service.StartwebServer("6768")
}
