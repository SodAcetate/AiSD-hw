package t2

import (
	"fmt"
)

func Run(input string) {
	count := 0
	length := 0
	max := 0

	for i := 0; i < len(input); i++ { // Перебираем все возможные последовательности потому что я не придумал как это сделать за линейное время
		if input[i] == '(' { // Последовательность может начинаться только с (
			count = 0
			length = 0
			for j := i; j < len(input); j++ {
				if input[j] == '(' {
					count += 1
				}
				if input[j] == ')' {
					count -= 1
				}
				if count == 0 { // Если кол-во ( равно кол-ву ) значит получили верную последовательность -- обновляем макс длину
					length = j - i + 1
				}
				if count < 0 { // Если кол-во ) больше кол-ва ( значит верную последовательность уже не получим -- фиксируем результат
					break
				}
			}
			if length > max {
				max = length
			}
		}
	}

	fmt.Println(max)
}
