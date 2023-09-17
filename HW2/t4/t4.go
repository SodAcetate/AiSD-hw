package t4

import "fmt"

func Run(ticket []int, k int) {
	time := 0

	for i := 0; i < k; i++ { // Те, кто перед K в очереди, пройдут не более ticket[k] раз
		if ticket[i] > ticket[k] {
			time += ticket[k]
		} else {
			time += ticket[i]
		}
	}

	time += ticket[k] // К должен пройти очередь ticket[k] раз

	for i := k + 1; i < len(ticket); i++ { // Те, кто за K в очереди, проходят не более ticket[k]-1 раз
		if ticket[i] > ticket[k]-1 {
			time += ticket[k] - 1
		} else {
			time += ticket[i]
		}
	}

	fmt.Printf("time: %v s\n", time)
}
