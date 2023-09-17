package types

import (
	"fmt"
)

type List struct {
	Key  int
	Next *List
}

func InsertAfter(where *List, x *List) {
	x.Next = where.Next
	where.Next = x
}

func RemoveAfter(where List) {
	x := where.Next
	where.Next = where.Next.Next
	x.Next = nil // тут должно быть освобождение памяти, но в го это делает GC как я понял
}

func (list List) Print() {
	for list.Next != nil {
		fmt.Printf("%v ", list.Key)
		list = *list.Next
	}
	fmt.Printf("%v ", list.Key)
	fmt.Println()
}

func MakeList(array []int) List {
	var head List
	head.Key = array[0]
	cur := &head

	for i := 1; i < len(array); i++ {
		var x List
		x.Key = array[i]
		InsertAfter(cur, &x)
		cur = &x
	}

	return head
}

// --- Очередь ---

func Dequeue(Q []int) ([]int, int) {
	x := Q[0]
	Q = Q[1:]
	return Q, x
}

func Enqueue(Q []int, x int) []int {
	Q = append(Q, x)
	return Q
}

// --- Стек ---

func Pop(S []int) ([]int, int) {
	x := S[len(S)-1]
	S = S[:len(S)-1]
	return S, x
}

func Push(S []int, x int) []int {
	S = append(S, x)
	return S
}
