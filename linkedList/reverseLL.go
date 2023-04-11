// Desc: Reverse a linked list

package main

import (
    "fmt"
)


type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func Constructor() LinkedList {
    return LinkedList{}
}


func (l *LinkedList) insert(data int) {
    node := &Node{data: data}
    node.next = l.head
    l.head = node
}

func (l *LinkedList) print() {
    for node := l.head; node != nil; node = node.next {
        fmt.Printf("%d ", node.data)
    }
    fmt.Println()
}

func (l *LinkedList) reverse() {
    var prev *Node
    curr := l.head
    for curr != nil {
        next := curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    l.head = prev
}

func main() {
    l := &LinkedList{}
    l.insert(1)
    l.insert(2)
    l.insert(3)
    l.insert(4)
    l.insert(5)
    l.print()
    l.reverse()
    l.print()

}


