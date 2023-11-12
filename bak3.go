package main

import "fmt"

func bak3() {
	/* create array with fixed value */
	// arr := []int{2, 4, 4, 5}

	/* create [0 0 0 0] array */
	arr := make([]int, 4)
	arr[0] = 30
	fmt.Println(arr)

	/* slice string */
	txt := "today is sunday"
	fmt.Println(txt[0:5])
	fmt.Println(arr[0:1])

	/* check lenght */
	fmt.Println(len(txt))
}
