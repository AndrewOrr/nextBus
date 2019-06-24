package main

import (
	"flag"
	"fmt"
	"nextBus/pkg/nextBus"
	"os"
)

func main() {
	useCommandLineArgs := false
	route := flag.String("route", func() string { useCommandLineArgs = true; return "default" }(), "Name of the bus route")
	stop := flag.String("stop", func() string { useCommandLineArgs = true; return "default" }(), "Name of the stop")
	direction := flag.String("direction", func() string { useCommandLineArgs = true; return "default" }(), "Cardinal Direction")
	if useCommandLineArgs {
		args := os.Args[1:]
		if len(args) < 3 {
			fmt.Println("run -h for options and command line suggestions")
			return
		}
		route = &args[0]
		stop = &args[1]
		direction = &args[2]
	}
	nextTime, err := getNextTime(*route, *stop, *direction)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(nextTime)

}

func getNextTime(inputRoute, inputStop, inputDirection string) (string, error) {
	route, err := nextBus.GetRoute(inputRoute)
	if err != nil {
		return "", err
	}
	direction, err := nextBus.GetDirection(route, inputDirection)
	if err != nil {
		return "", err
	}
	stop, err := nextBus.GetStop(route, direction, inputStop)
	if err != nil {
		return "", err
	}
	return nextBus.GetNextTrip(route, direction, stop)
}
