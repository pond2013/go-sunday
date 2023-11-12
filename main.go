package main

import "fmt"

type SaveLogInterface interface {
	SaveLog()
}

func SaveLog(st SaveLogInterface) {
	st.SaveLog()
}

type TransferBBL struct {
	name string
}

type TransferKTB struct {
	name string
}

func (tBBL TransferBBL) SaveLog() {
	fmt.Println("save to database", tBBL)
}

func (tKTB TransferKTB) SaveLog() {
	fmt.Println("save to database", tKTB)
}

func main() {
	transA := TransferBBL{name: "BBL"}
	transB := TransferKTB{name: "KTB"}
	SaveLog(&transA)
	SaveLog(&transB)
}
