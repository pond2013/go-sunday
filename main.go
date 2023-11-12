package main

import (
	"fmt"
	"time"
)

func Ping1S(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("ping : ", i)
		time.Sleep(1 * time.Second)
	}
	c <- 10
}

func SendNoti5S(c chan string) {
	fmt.Println("Send noti")
	time.Sleep(5 * time.Second)
	fmt.Println("Send noti")
	c <- "success"
}

func main() {
	c1 := make(chan int)
	c2 := make(chan string)
	go Ping1S(c1)
	go SendNoti5S(c2)

	pingVal, notMess := <-c1, <-c2
	fmt.Println(pingVal, notMess)
	fmt.Println("completed")
}
