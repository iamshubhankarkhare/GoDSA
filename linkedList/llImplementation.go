// Description: Linked List Implementation in Go
package main

import "fmt"
// create a node struct
type Node struct {
    data int
    next *Node
}

// create a linked list struct
type LinkedList struct {
    head *Node
}

// add a node to the linked list
func (list *LinkedList) add(data int) {
    // create a new node
    node := &Node{data: data}
    // if the head is nil, set the head to the new node
    if list.head == nil {
        list.head = node
    } else {
        // else, set the next node to the new node
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = node
    }
}

// print the linked list
func (list *LinkedList) print() {
    current := list.head
    for current != nil {
        fmt.Println(current.data)
        current = current.next
    }
}


func main() {
    fmt.Println("Hello, World!")
    // add numbers to the linked list
    list := LinkedList{}
    list.add(1)
    list.add(2)
    list.add(3)
    list.add(4)
    list.add(5)

    // print the linked list
    list.print()
}
