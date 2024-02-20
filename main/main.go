package main

import (
	"fmt"
	"github.com/figurecode/mymath"

	"ypmodule/calc"

	"workspace/user/repo/somepackage"
)

func main() {
	fmt.Println(calc.AddInts(1, 2))

	fmt.Println(mymath.Add(1, 3))

	somepackage.Func()
	fmt.Println("Ok")
}
