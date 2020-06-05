package main

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var l3 *ListNode

	if l1.Val >= l2.Val {
		l3 = l2
		l3.Next = mergeTwoLists(l1, l2.Next)
	} else {
		l3 = l1
		l3.Next = mergeTwoLists(l1.Next, l2)
	}

	return l3
}

func main() {

	// l1
	l10 := ListNode {Val:1,}
	l11 := ListNode {Val:3,}
	l12 := ListNode {Val:5,}
	l13 := ListNode {Val:7,}
	l14 := ListNode {Val:9,}
	
	l10.Next = &l11
	l11.Next = &l12
	l12.Next = &l13
	l13.Next = &l14

	// l2
	l20 := ListNode {Val:1,}
	l21 := ListNode {Val:3,}
	l22 := ListNode {Val:6,}
	l23 := ListNode {Val:8,}
	l24 := ListNode {Val:10,}
	
	l20.Next = &l21
	l21.Next = &l22
	l22.Next = &l23
	l23.Next = &l24

	l30 := mergeTwoLists(&l10, &l20)

	println(l30.Val)
	println(l30.Next.Val)
	println(l30.Next.Next.Val)
	println(l30.Next.Next.Next.Val)
	println(l30.Next.Next.Next.Next.Val)
	println(l30.Next.Next.Next.Next.Next.Val)
	println(l30.Next.Next.Next.Next.Next.Next.Val)
	println(l30.Next.Next.Next.Next.Next.Next.Next.Val)
	println(l30.Next.Next.Next.Next.Next.Next.Next.Next.Val)
	println(l30.Next.Next.Next.Next.Next.Next.Next.Next.Next.Val)

}