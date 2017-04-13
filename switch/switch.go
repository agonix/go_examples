package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend yey")
	default:
		fmt.Println("Its a work day")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Para dite")
	case t.Hour() > 12:
		fmt.Println("Mas dite")
	}
}
