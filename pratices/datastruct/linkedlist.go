package datastruct

import "fmt"

// MyLinkedList struct
type MyLinkedList struct {
	Val  int
	Prev *MyLinkedList
	Next *MyLinkedList
}

// Constructor method
/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (list *MyLinkedList) Get(index int) int {
	res := -1
	// ignore head node
	p := list.Next
	for p != nil && index > 0 {
		p = p.Next
		index--
	}
	// find node
	if index == 0 && p != nil {
		res = p.Val
	}
	return res
}

// AddAtHead add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (list *MyLinkedList) AddAtHead(val int) {
	node := &MyLinkedList{val, nil, nil}
	p := list.Next
	if p != nil {
		p.Prev = node
	}
	node.Next = p
	list.Next = node
}

// AddAtTail Append a node of value val to the last element of the linked list.
func (list *MyLinkedList) AddAtTail(val int) {
	node := &MyLinkedList{val, nil, nil}
	p := list
	for p.Next != nil {
		p = p.Next
	}
	p.Next = node
	node.Prev = p
}

// AddAtIndex Add a node of value val before the index-th node in the linked list.
// If index equals to the length of linked list, the node will be appended to the end of linked list.
// If index is greater than the length, the node will not be inserted.
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	// index == 0
	if index == 0 {
		list.AddAtHead(val)
		return
	}
	p := list.Next
	for p.Next != nil && index > 1 {
		p = p.Next
		index--
	}
	if index == 1 {
		node := &MyLinkedList{val, nil, nil}
		q := p.Next
		if q != nil {
			q.Prev = node
		}
		node.Next = q
		node.Prev = p
		p.Next = node
	}
}

// DeleteAtIndex Delete the index-th node in the linked list, if the index is valid.
func (list *MyLinkedList) DeleteAtIndex(index int) {
	p := list.Next
	if index == 0 {
		if p != nil {
			list.Next = p.Next
			p.Next.Prev = nil
		}
		return
	}
	for p.Next != nil && index > 1 {
		p = p.Next
		index--
	}
	// if index > 1, then index is invalid
	if index == 1 {
		// index equals length，do nothing
		if p.Next == nil {
			return
		}
		// index < length, delete p.Next
		if p.Next.Next != nil {
			p.Next.Next.Prev = p
		}
		p.Next = p.Next.Next
	}
}

// func (this *MyLinkedList) Length() {
// 	p := this.Next
// 	for p != nil {

// 	}
// }

// Output foreach display list
func (list *MyLinkedList) Output() {
	// 忽略头结点
	p := list.Next
	for p != nil {
		fmt.Printf("%d \t", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}

// Link use for test list
func Link() {
	obj := Constructor()
	obj.AddAtHead(7)
	fmt.Printf("indx: %d val: %d \n", 2, obj.Get(2))
	obj.Output()
	obj.AddAtHead(2)
	fmt.Printf("indx: %d val: %d \n", 2, obj.Get(2))
	obj.Output()
	obj.AddAtHead(1)
	fmt.Printf("indx: %d val: %d \n", 2, obj.Get(2))
	fmt.Printf("indx: %d val: %d \n", 1, obj.Get(1))
	obj.AddAtIndex(1, 3)
	obj.AddAtIndex(4, 4)
	obj.DeleteAtIndex(5)
	obj.Output()
	obj.DeleteAtIndex(4)
	obj.Output()
	obj.DeleteAtIndex(1)
	obj.Output()
	obj.DeleteAtIndex(0)
	obj.Output()
	obj.AddAtIndex(0, 0)
	obj.Output()
	fmt.Printf("indx: %d val: %d \n", 0, obj.Get(0))
	// obj.AddAtIndex(3, 0)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
[[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]

["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
[[],            [7],         [2],        [1],       [3,0],       [2],             [6],       [4],[4],[4],[5,0],[6]]
*/
