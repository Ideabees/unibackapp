package main

import (
	"fmt"

	"github.com/v1/uniapp/internal/router"
)

func main() {
	err := router.IntiateRoutes()
	if err != nil{
		fmt.Println("Error in starting the mudule", err)
	}
}
