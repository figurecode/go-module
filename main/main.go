package main

import (
	"fmt"

	"ypmodule/calc"

	"workspace/user/repo/somepackage"
)

func main() {
	fmt.Println(calc.AddInts(1, 2))

	somepackage.Func()
	fmt.Println("Ok")
}
