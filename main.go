package main

import (
	"fmt"

	"github.com/TDDD96-G11-BLHub/dbman/lib"
)

// DB PR
func main() {
	fmt.Println("Hello from BLHub database manager!")

	//client := lib.SetupConnection()

	lib.SetupConnection()
	lib.ConnectHello()
	lib.FetchHello()
	lib.UpdateHello()

}
