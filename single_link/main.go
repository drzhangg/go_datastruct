package main

import "fmt"

type Node struct {
	id   int
	name string
	next *Node
}

//向链表末尾插入一个节点
func InsertNode(head *Node, newNode *Node) {
	/*
		1.先找到链表的最后一个节点
		2.创建一个辅助节点
		3.将newNode加入链表的最后
	*/
	temp := head
	//循环遍历链表
	for {
		//如果链表的next为nil，说明是最后一个节点
		if temp.next == nil {
			break
		}
		//一直循环链表，直到最后一个节点
		temp = temp.next
	}
	//将newNode加入节点的最后
	temp.next = newNode
}

//获取链表列表
func ListNode(head *Node) {
	/*
		1.先判断这个链表是否为空
		2.遍历链表，打印每个节点的值
	*/
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}

	for {
		fmt.Printf("[%d, %s] -->", temp.next.id, temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	node := &Node{}
	node1 := &Node{
		id:   1,
		name: "jerry",
	}
	node2 := &Node{
		id:   2,
		name: "tom",
	}

	InsertNode(node, node2)
	InsertNode(node, node1)


	ListNode(node)
}
