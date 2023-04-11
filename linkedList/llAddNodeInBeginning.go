// Desc: Add a linked list node in the beginning of the linked list

package main

import ( 
    "fmt"
)

type Node struct{
    next *Node
    data int
}

type linkedList struct{
    head *Node
}

func (list *linkedList) addNode(data int){
    node:=&Node{data:data}
    curr:=list.head
    if curr == nil {
        list.head=node
        return
    }else{
        for curr.next != nil {
            curr=curr.next
        }
        curr.next=node
    }

}

func (list *linkedList) print(){
    curr:= list.head
    for curr.next != nil {
        fmt.Println(curr.data)
        curr=curr.next
    }
}

func main(){
    fmt.Println("hello")
    list:=&linkedList{}
    list.addNode(1)
    list.addNode(2)
    list.addNode(3)
    list.addNode(4)
    list.addNode(5)
    list.print()

    // take an integer input from the user
    fmt.Println("Enter a number to add in the beginning of the linked list")
    var num int
    fmt.Scanln(&num)
    node:=&Node{data:num}
    node.next=list.head
    list.head=node
    list.print()
}
