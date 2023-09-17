package t1

import (
	"fmt"
)

func Run(a []int, m int) {
	n := len(a)
	sum := 0

	for i := 0; i < m; i++ { // Считаем сумму первых m элементов
		sum += a[i]
	}

	fmt.Printf("%d ", sum)

	for i := m; i < n; i++ {
		sum -= a[i-m] // Вычитаем первый элемент суммы
		sum += a[i]   // Добавляем следующий
		fmt.Printf("%d ", sum)
	}

	fmt.Println()

}
