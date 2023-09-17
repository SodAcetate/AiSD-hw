package t3

import (
	types "hw/shared"
)

func Run(head *types.List) {

	var prev *types.List = nil
	var cur *types.List = nil
	next := head.Next

	for next != nil {
		prev = cur
		cur = next
		next = cur.Next
		cur.Next = prev
	}

	head.Next = cur

}
