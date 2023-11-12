package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Address  string
	PostCode string
}

type UserProfile struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"LastName"`
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
		Age:       24,
		Height:    160,
	}
	fmt.Println(userProfile)
	userProfile.Address.PostCode = "11203"
	fmt.Println(userProfile.ToFullDesc())
	byteTxtJson, err := json.MarshalIndent(userProfile, "", "  ")
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println(string(byteTxtJson))
	}

	/* Json data define */
	dataJson :=
		`{
		"first_name": "sad",
		"LastName": "salty",
		"Age": 24,
		"Height": 160,
		"Address": {
		  "Address": "",
		  "PostCode": "11203"
		},
		"Bill": {
		  "BillAddress": ""
		}
	}`
	var myProfile UserProfile
	err = json.Unmarshal([]byte(dataJson), &myProfile)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println((myProfile))
	}
}
