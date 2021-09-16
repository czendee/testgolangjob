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


// RemoveDups removes duplicates from linked list
// Uses additional map to track duplicates
func RemoveDups(n *Node) {
	var prev *Node
	dups := make(map[int]bool)

	for n != nil {
		d := n.data

		if _, ok := dups[d]; ok {
			// n.Remove
			prev.next = n.next
		} else {
			dups[d] = true
			prev = n
		}

		n = n.next
	}
}

// RemoveDupsRunner removes duplicates from linked list
// Uses runner technique, no need in additional data structures
func RemoveDupsRunner(n *Node) {
	for n != nil {
		r := n
		for r.next != nil {
			if r.next.data == n.data {
				r.next = r.next.next
			} else {
				r = r.next
			}
		}

		n = n.next
	}
}

func test01Sucess(){
    
	node := New(1)
	node.AppendToTail(1)
	node.AppendToTail(2)
	node.AppendToTail(2)
	node.AppendToTail(2)
	node.AppendToTail(3)
	node.AppendToTail(1)
        log.Print("test01Sucess step 1")
	RemoveDups(node)
	n := node
	for i, d := range []int{1, 2, 3} {
	        //log.Print("test01Sucess step 2")
		if n.data != d {
  		        log.Printf("%dth node expected to have data %v, got %v", i+1, d, n.data)

		}else{
	            log.Printf("test01Sucess step 2 , el data es %v, el d es: %v",n.data,d )
		}		
		n = n.next
	}

}

func test02Sucess(){
	node := New(1)
	node.AppendToTail(1)
	node.AppendToTail(2)
	node.AppendToTail(2)
	node.AppendToTail(2)
	node.AppendToTail(3)
	node.AppendToTail(1)
	log.Print("test02Sucess step 1")        
	RemoveDupsRunner(node)
	n := node
	for i, d := range []int{1, 2, 3} {
//	        log.Printf("test02Sucess step 2 , el data es %v, el d es: %v",n.data,d )
		if n.data != d {
  		        log.Printf("%dth node expected to have data %v, got %v", i+1, d, n.data)
		}else{
	            log.Printf("test02Sucess step 2 , el data es %v, el d es: %v",n.data,d )
		}
		n = n.next
	}

}



func main() {
     test01Sucess();
     test02Sucess()
}
