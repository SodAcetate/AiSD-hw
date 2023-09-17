package t2

import (
	"fmt"
)

func Run(input string) {
	count := 0
	length := 0
	max := 0

	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			count += 1
		}
		if input[i] == ')' {
			count -= 1
		}
		length += 1
		if count == 0 {
			if length > max {
				max = length
			}
		}
		if count < 0 {
			count = 0
			length = 0
		}
	}

	count = 0
	length = 0

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == '(' {
			count -= 1
		}
		if input[i] == ')' {
			count += 1
		}
		length += 1
		if count == 0 {
			if length > max {
				max = length
			}
		}
		if count < 0 {
			count = 0
			length = 0
		}
	}

	fmt.Println(max)
}
