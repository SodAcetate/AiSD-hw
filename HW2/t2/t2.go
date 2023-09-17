package t2

import (
	"fmt"
	types "hw/shared"
)

func Run(input string) {
	var stack []rune // Стек символов (скобок)
	var x rune
	count := 0
	max := 0

	for _, r := range input {
		if len(stack) > 0 {
			stack, x = types.Pop(stack)
			if x == '(' && r == ')' {
				count += 2 // Две скобки вида () самоуничтожаются и инкрементируют счётчик
			} else {
				stack = types.Push(stack, x) // В противном случае просто пишем скобки в стек
				stack = types.Push(stack, r)

				if r == ')' { // Если последние две скобки вида )), то их уже не уничтожить
					if max < count { // В таком случае сбрасываем счётчик, т.к. мы точно не внутри корректной посл-ти
						max = count
					}
					count = 0
				}
			}
		} else {
			stack = types.Push(stack, r) // Первый элемент просто пишем в стек
		}
	}

	if max < count {
		max = count
	}

	fmt.Println(max)
}
