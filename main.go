package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//example input string
	/*
		fmt.Println("enter your name:")

		// var txtInput string
		txtInput := ""

		fmt.Scan(&txtInput)
		fmt.Println("Hello " + txtInput)
	*/

	// check variable type
	age := 2
	fmt.Println(reflect.TypeOf(age))

	// convert type
	height := "100"
	h, err := strconv.Atoi(height)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(h)
	}

	water := "12.40"
	fWater, err := strconv.ParseFloat(water, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fWater)
	}
}
