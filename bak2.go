package main

import (
	"fmt"
	"time"
)

func IsInSystem(username string) bool {
	return true
}

func GetUserDetail(username string) (int, string) {
	return 201, "manager"
}

func GetDeparture(username string, departure *string) {
	if username != "" {
		*departure = "home"
	}
}

func CheckLogin(username string, password string) {
	if IsInSystem(username) {
		fmt.Println("found user in system")
		GetUserDetail(username)
		departure := ""
		GetDeparture(username, &departure)
	}
}

func GetMember() {
	fmt.Println("please wait...")
	time.Sleep(3 * time.Second)
}

func LogEnd() {
	time.Now()
	fmt.Println("completed program ")
	fmt.Println(time.Now())
}

func CheckServerResponse() {
	fmt.Println("check server time")
	time.Sleep(3 * time.Second)
	panic("server error")
}

func bak2() {
	/*
		defer first -> excute last
		for example LogEnd() will excute before anon()
	*/
	// defer func() {
	// 	fmt.Println("example anonymouse function")
	// }()
	// defer LogEnd()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover ")
			fmt.Println(r)
		}
	}()

	GetMember()

	CheckServerResponse()

}
