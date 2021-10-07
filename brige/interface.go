package main

import "fmt"

type Brand interface {
	Call()
}


type Apple struct {

}

func NewApple() *Apple {
	return &Apple{}
}

func (a Apple) Call() {
	fmt.Println("apple call")
}

type Huawei struct {

}

func NewHuawei() *Huawei {
	return &Huawei{}
}

func (h Huawei) Call() {
	fmt.Println("huawei call")
}



