package main

func main() {

	a := Folder{Style{Brand: *NewApple()}}

	h := Folder{Style{Brand: NewHuawei()}}

	a1 := Upright{Style{Brand: NewApple()}}

	a.Call()
	h.Call()
	a1.Call()
}
