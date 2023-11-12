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

func (tBBL *TransferBBL) SaveLog() {
	fmt.Println("save to database", tBBL.name)
}

func (tKTB *TransferKTB) SaveLog() {
	fmt.Println("save to database", tKTB.name)
}

func main() {
	transA := TransferBBL{name: "BBL test"}
	transB := TransferKTB{name: "KTB test"}
	SaveLog(&transA)
	SaveLog(&transB)
}
