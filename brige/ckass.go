package main

import "fmt"

type Style struct {
	Brand Brand
}

type Folder struct {
	Style
}

func (f Folder) Show()  {
	fmt.Println("is folder")
}

func (f Folder) Call()  {
	f.Show()
	f.Brand.Call()
}

type Upright struct {
	Style
}

func (u Upright) Show()  {
	fmt.Println("is upright")
}

func (u Upright) Call() {
	u.Show()
	u.Brand.Call()
}







