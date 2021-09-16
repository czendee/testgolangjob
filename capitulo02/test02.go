package main

import "log"

// Node struct for linked list
type Node struct {
	data int
	next *Node
}

// New returns a new node
func New(d int) *Node {
	n := new(Node)
	n.data = d
	return n
}

// AppendToTail appends a new node with data d at the tail of list l
func (node *Node) AppendToTail(d int) {
	end := New(d)
	for node.next != nil {
		node = node.next
	}
	node.next = end
}


// KthToLast returns kth to the last element
func KthToLast(node *Node, k int) *Node {
	n := node
	l := 0
	for ; n != nil; l++ {
		n = n.next
	}

	n = node
	for i := 1; i <= l-k-1; i++ {
		n = n.next
	}

	return n
}

// KthToLastRecursive returns kth to the last element
func KthToLastRecursive(node *Node, k int) (*Node, int) {
	if node.next == nil {
		return node, 0
	}

	n, i := KthToLastRecursive(node.next, k)
	i++

	if i == k {
		return node, i
	}

	return n, i
}

// KthToLastIterative returns kth to the last element
func KthToLastIterative(node *Node, k int) *Node {
	iterator := node
	runner := node

	for i := 0; i < k; i++ {
		iterator = iterator.next
	}

	for iterator.next != nil {
		iterator = iterator.next
		runner = runner.next
	}

	return runner
}


func test01Sucess(){
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)
	node.AppendToTail(4)
	node.AppendToTail(5)

	k := 0
	actual := KthToLast(node, k).data
	expected := 5
	if actual != expected {
		log.Printf("%dth element expected to be %v, got %v", k, expected, actual)
	}else{
	   log.Printf("test01Sucess ok 1");
	}    

	k = 2
	actual = KthToLast(node, k).data
	expected = 3
	if actual != expected {
		log.Printf("%dth element expected to be %v, got %v", k, expected, actual)
	}else{
	   log.Printf("test01Sucess ok 2");
	}    

	k = 4
	actual = KthToLast(node, k).data
	expected = 1
	if actual != expected {
		log.Printf("%dth element expected to be %v, got %v", k, expected, actual)
	}else{
	   log.Printf("test01Sucess ok 3");
	}    
}

func test02Sucess(){
	node := New(1)
	node.AppendToTail(2)
	node.AppendToTail(3)
	node.AppendToTail(4)
	node.AppendToTail(5)

	k := 0
	actual, _ := KthToLastRecursive(node, k)
	expected := 5
	if actual.data != expected {
		log.Printf("%dth element expected to be %v, got %v", k, expected, actual.data)
	}else{
	   log.Printf("test02Sucess ok 1");
	}

	k = 2
	actual, _ = KthToLastRecursive(node, k)
	expected = 3
	if actual.data != expected {
		log.Printf("%dth element expected to be %v, got %v", k, expected, actual.data)
	}else{
	   log.Printf("test02Sucess ok 2");
	}

	k = 4
	actual, _ = KthToLastRecursive(node, k)
	expected = 1
	if actual.data != expected {
		log.Printf("%dth element expected to be %v, got %v", k, expected, actual.data)
	}else{
	   log.Printf("test02Sucess ok 3");
	}
}



func main() {
     test01Sucess();
     test02Sucess()
}
