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

//按照id顺序进行插入
func InsertNodeDesc(head *Node, newHead *Node) {
	temp := head
	flag := true

	for {
		if temp.next == nil { //说明最后一个节点
			break
		} else if temp.next.id > newHead.id { //说明newNode应该在head前面
			break
		} else if temp.next.id == newHead.id { //该节点已经插过
			flag = false
			break
		}
	}

	if !flag {
		fmt.Printf("对不起，该id已经存在:%d", newHead.id)
		return
	} else {
		newHead.next = temp.next
		temp.next = newHead
	}
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

//根据id删除节点
func DelNode(head *Node, id int) {
	temp := head
	flag := false

	for {
		if temp.next == nil {
			break
		} else if temp.next.id == id {
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		//eg：删除节点2，将节点1的next指向节点3
		temp.next = temp.next.next
	} else {
		fmt.Println("该id不存在")
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
	node3 := &Node{
		id:   3,
		name: "bob",
	}

	InsertNode(node, node2)
	InsertNode(node, node1)
	InsertNode(node, node3)

	ListNode(node)

	DelNode(node, 2)
	ListNode(node)
}

//TODO:从队列头部取，从尾部取
