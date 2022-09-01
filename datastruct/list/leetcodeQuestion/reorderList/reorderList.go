package main

import (
	. "holy-algorithm/datastruct/list/node"
)

/**
题目：https://leetcode.cn/problems/reorder-list/

重排链表
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：

输入：head = [1,2,3,4]
输出：[1,4,2,3]
示例 2：

输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]

提示：
链表的长度范围为 [1, 5 * 104]
1 <= node.val <= 1000

*/
func main() {
	list1 := MakeListNode([]int{1, 2, 3, 4, 5})
	reorderList(list1)
	PrintListNode(list1)
}

// reorderList 链表反转+链表合并+链表中点
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	mid := middleNode(head)
	// 保存两部分头结点
	l1 := head
	l2 := mid.Next
	// 断开mid与后半部分的联系
	mid.Next = nil
	// 反转 mid~tail部分的链表
	l2 = reverseList(l2)
	// 合并2部分链表
	mergeList(l1, l2)
}

// 寻找中间结点 O(n) O(1)
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

// mergeList 链表合并 O(n) O(1)
func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}
