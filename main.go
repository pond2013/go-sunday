package main

import "fmt"

type Address struct {
	Address  string
	PostCode string
}

type UserProfile struct {
	Firstname string
	Lastname  string
	Age       int
	Height    float32

	Address Address

	Bill struct {
		BillAddress string
	}
}

func (u UserProfile) ToFullDesc() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

func main() {
	fmt.Println("xx")

	/* example map */
	user := map[string]string{}
	user["username"] = "pondhub"
	user["password"] = "xxxxxx"
	fmt.Println(user)
	fmt.Println(user["username"])

	/* example struct */
	userProfile := UserProfile{
		Firstname: "pondhub",
		Lastname:  "salty",
	}
	userProfile.Address.PostCode = "11203"
	fmt.Println(userProfile)

	fmt.Println(userProfile.ToFullDesc())
}
