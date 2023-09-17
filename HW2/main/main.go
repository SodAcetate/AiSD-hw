package main

import (
	"fmt"
	types "hw/shared"
	"hw/t1"
	"hw/t2"
	"hw/t3"
	"hw/t4"
)

func main() {

	fmt.Println("== Task 1 ==")
	t1.Run([]int{1, 2, 3, 4, 5}, 2)

	fmt.Println("== Task 2 ==")
	t2.Run(")(()())(")

	fmt.Println("== Task 3 ==")
	list := types.MakeList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	list.Print()
	t3.Run(&list)
	list.Print()

	fmt.Println("== Task 4 ==")
	t4.Run([]int{1, 3, 3, 1, 3}, 2)

}
