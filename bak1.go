package main

import (
	"fmt"
)

func bak1() {
	// Example input string
	/*
		fmt.Println("enter your name:")

		// var txtInput string = ""
		txtInput := ""

		fmt.Scan(&txtInput)
		fmt.Println("Hello " + txtInput)
	*/

	// check variable type
	/*
		age := 2
		fmt.Println(reflect.TypeOf(age))
	*/

	// convert type
	/*
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

		txtShow := ""
		txtShow = fmt.Sprintf("water = %.2f, height = %d", fWater, h)
		fmt.Println(txtShow)
	*/

	// Iteration for
	/*
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	*/

	// Switch case
	score := 80
	switch score {
	case 80:
		{
			fmt.Println("A")
		}
	case 70:
		{
			fmt.Println("B")
		}
	}

	if score >= 80 {
		fmt.Println("A")
	} else if score < 80 && score > 70 {
		fmt.Println("B")
	} else {
		fmt.Println("unknown")
	}
}
