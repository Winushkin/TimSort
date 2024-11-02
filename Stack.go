package main

//      prev <- Node -> next

type Node struct {
	Value []int
	prev  *Node
}

func newNode(value []int) *Node {
	return &Node{value, nil}
}

// node1 <- node2 <- node3 <- root

type Stack struct {
	root   *Node
	length int
}

func (stack *Stack) Len() int {
	return stack.length
}

func (stack *Stack) Stack() {
	stack.length = 0
	stack.root = nil
}

func (stack *Stack) PushBack(value []int) {
	var cur *Node = newNode(value)
	if stack.length == 0 {
		stack.root = cur
	} else {
		cur.prev = stack.root
		stack.root = cur
	}
	stack.length++
}

func (stack *Stack) Peek() []int {
	if stack.length != 0 {
		return stack.root.Value
	}
	return nil
}

func (stack *Stack) Front() []int {
	cur := stack.root

	for cur.prev != nil {
		cur = cur.prev
	}
	return cur.Value
}

func (stack *Stack) PopBack() []int {
	if stack.length != 0 {
		cur := stack.root
		stack.root = cur.prev
		stack.length--
		return cur.Value
	}
	return nil
}

func (stack *Stack) PopFront() []int {
	if stack.length != 0 {
		cur := stack.root
		var rtn []int
		if stack.length == 1 {
			rtn = cur.Value
			stack.root = nil
		} else {
			for cur.prev.prev != nil {
				cur = cur.prev
			}
			rtn = cur.prev.Value
			cur.prev = nil
		}
		stack.length--
		return rtn
	}
	return nil
}
