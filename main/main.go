package main

import (
	"arrint"
	"fmt"
	"github.com/figurecode/mymath"

	"ypmodule/calc"

	"workspace/user/repo/somepackage"
)

func main() {
	fmt.Println(calc.AddInts(1, 2))

	fmt.Println(mymath.Add(1, 3))

	arrA := arrint.ArrInt{1, 3, 4, 5}
	arrB := arrint.ArrInt{2, 4, 5, 6}
	arrAds := arrint.Add(arrA, arrB)

	fmt.Println(arrAds.String())

	somepackage.Func()
	fmt.Println("Ok")
}
