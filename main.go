package main

import (
	"fmt"
	"nextBus/pkg/nextBus"
)

func main() {
	//	given route stop and direction
	// assume no validation is necessary
	routes := nextBus.GetRoutes()
	fmt.Println(len(routes))

}
