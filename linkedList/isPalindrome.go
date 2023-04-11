// desc: find is linked list is palindrome

package main

import (
    "fmt"
)

type Node struct {
    data int
    next *Node
}

func findMiddle(head *Node) *Node {
    slow := head
    fast := head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
    }
    return slow
}
func reverse(head *Node) *Node {
    var prev *Node
    for head != nil {
        next := head.next
        head.next = prev
        prev = head
        head = next
    }
    return prev
}

func isPalindrome(head *Node) bool {
    if head == nil || head.next == nil {
        return true
    }
    middle := findMiddle(head)
    head2 := reverse(middle)
    for head != nil && head2 != nil {
        if head.data != head2.data {
            return false
        }
        head = head.next
        head2 = head2.next
    }
    return true
}

func main() {
    head := &Node{data: 1}
    head.next = &Node{data: 2}
    head.next.next = &Node{data: 3}
    head.next.next.next = &Node{data: 2}
    head.next.next.next.next = &Node{data: 3}
    fmt.Println(isPalindrome(head))
}
