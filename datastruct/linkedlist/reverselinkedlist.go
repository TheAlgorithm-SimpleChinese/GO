// -*- coding: utf-8 -*-
//author: liuyang
//software: GoLand
//file: reverselinkedlist.go
//time: 2022/5/19 1:04

package main

import (
	"fmt"
)

type Node struct {
	Data int
	next *Node
}

func Reverselink(head Node) (int, Node) {
	//判断链表是否为空
	if head.next == nil {
		return -1, Node{}
	}

	pre := head.next  //前驱节点
	cur := head.next  //当前节点
	next := head.next //后继节点
	// 把链表首节点变为尾节点
	cur = head.next
	next = cur.next
	cur.next = nil
	pre = cur
	cur = next
	// 使当前遍历到的节点cur指向其前驱节点
	for cur.next != nil {
		next = cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	// 链表最后一个节点指向倒数第二个节点
	cur.next = pre
	// 链表头结点指向原来链表的尾节点
	head.next = cur
	return 1, head
}

func main() {
	// 创建节点
	var (
		Head = Node{}
		P1   = Node{1, nil}
		P2   = Node{2, nil}
		P3   = Node{3, nil}
		P4   = Node{4, nil}
		P5   = Node{5, nil}
		P6   = Node{6, nil}
		P7   = Node{7, nil}
	)
	// 建立列表
	Head.next = &P7
	P7.next = &P6
	P6.next = &P5
	P5.next = &P4
	P4.next = &P3
	P3.next = &P2
	P2.next = &P1

	fmt.Println("逆置前链表为：")
	temp := Head.next
	for temp != nil {
		fmt.Println(temp.Data) // 输出当前链表数据
		temp = temp.next
	}

	fmt.Println("逆置后链表为：")
	temp1 := Node{}
	_, temp1 = Reverselink(Head)
	temp = temp1.next
	for temp != nil {
		fmt.Println(temp.Data) // 输出逆置后链表数据
		temp = temp.next
	}
}
