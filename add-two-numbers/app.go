package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode, inc ...int) *ListNode {

	var l3 *ListNode
	
	if l1 == nil && l2 == nil {
		if len(inc) > 0 && inc[0] == 1 {
			return &ListNode{Val:1,}
		}
		return nil
	}

	if l1 == nil {
		l1 = &ListNode {Val:0,}
	}
	if l2 == nil {
		l2 = &ListNode {Val:0,}
	}


	intTmp := l1.Val + l2.Val
	if len(inc) > 0 && inc[0] == 1 {
		intTmp += 1 
	}

	incTmp := 0
	if intTmp >= 10 {
		intTmp = intTmp - 10
		incTmp = 1
	} 
	l3 = &ListNode{Val:intTmp,}
	l3.Next = addTwoNumbers(l1.Next, l2.Next, incTmp)

	return l3
}

func main() {
// l1
	l10 := ListNode {Val:2,}
	l11 := ListNode {Val:4,}
	l12 := ListNode {Val:3,}
	// l13 := ListNode {Val:7,}
	// l14 := ListNode {Val:9,}
	
	l10.Next = &l11
	l11.Next = &l12
	// l12.Next = &l13
	// l13.Next = &l14

	// l2
	l20 := ListNode {Val:5,}
	l21 := ListNode {Val:6,}
	l22 := ListNode {Val:4,}
	// l23 := ListNode {Val:8,}
	// l24 := ListNode {Val:1,}
	
	l20.Next = &l21
	l21.Next = &l22
	// l22.Next = &l23
	// l23.Next = &l24

	l30 := addTwoNumbers(&l10, &l20)

	println(l30.Val)
	println(l30.Next.Val)
	println(l30.Next.Next.Val)
	// println(l30.Next.Next.Next.Val)
	// println(l30.Next.Next.Next.Next.Val)
	// println(l30.Next.Next.Next.Next.Next.Val)
}