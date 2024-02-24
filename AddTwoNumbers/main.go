package main

import (
	"math/big"
	"strconv"
)

func main() {

	head := &ListNode{Val: 1}

	for i := 0; i < 28; i++ {
		head.Append(0)
	}

	head.Append(1)

	head2 := &ListNode{Val: 5} // Primeiro nó
	head2.Append(6)
	head2.Append(4)

	addTwoNumbers(head, head2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sl1 := ""
	sl2 := ""

	for l1 != nil {
		sl1 = strconv.Itoa(l1.Val) + sl1
		l1 = l1.Next
	}

	for l2 != nil {
		sl2 = strconv.Itoa(l2.Val) + sl2
		l2 = l2.Next
	}

	n1, _ := new(big.Int).SetString(sl1, 10)
	n2, _ := new(big.Int).SetString(sl2, 10)

	result := []rune(new(big.Int).Add(n1, n2).String())

	firstValue, _ := strconv.Atoi(string(result[0]))
	r := &ListNode{Val: firstValue}

	for i := 1; i < len(result); i++ {
		vcon, _ := strconv.Atoi(string(result[i]))
		newNode := &ListNode{Val: vcon}
		newNode.Next = r
		r = newNode
	}

	return r
}

func (n *ListNode) Append(value int) {
	current := n
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Val: value}
}
