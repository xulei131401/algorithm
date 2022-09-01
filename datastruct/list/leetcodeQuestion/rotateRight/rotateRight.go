package main

/**
题目：https://leetcode.cn/problems/rotate-list/
旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
示例 1：

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

提示：

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109


*/
func main() {
	head := MakeListNode([]int{1, 2, 3, 5, 7})
	head1 := MakeListNode([]int{1, 2, 3, 5, 7})
	k := 2
	PrintListNode(rotateRight(head, k))
	PrintListNode(rotateRight1(head1, k))
}

// rotateRight 迭代法 O(n), O(1)
// 思路：链表转成环，然后进行移动，断开
func rotateRight(head *ListNode, k int) *ListNode {
	// head为nil,head只有一个结点，k=0的时候都不需要移动
	if head == nil || k == 0 || head.Next == nil {
		return head
	}

	// 计算链表的长度,同时tmp最终是链表的尾巴元素
	length := 1
	tailNode := head
	for tailNode.Next != nil {
		tailNode = tailNode.Next
		length++
	}

	// 如果长度为k的倍数，相当于没有移动
	add := length - k%length
	if add == length {
		return head
	}

	// 链表转成环
	tailNode.Next = head

	// 寻找最终的尾部元素
	for add > 0 {
		tailNode = tailNode.Next
		add--
	}

	// 尾部元素的下一个结点就是头结点
	ret := tailNode.Next
	// 断开环
	tailNode.Next = nil
	return ret
}

// rotateRight1 迭代法 O(n) O(1)
// 思路：找到新的头结点位置，把两部分链表进行拼接即可
func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// k = 0 相当于没有移动
	if k == 0 {
		return head
	}

	// 计算链表的长度
	length := 0
	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		length++
	}

	// 如果长度为k的倍数，相当于没有移动
	remainder := k % length
	if remainder == 0 {
		return head
	}

	// 旋转之后链表的头结点
	var newHead *ListNode
	// 遍历寻找新的头结点，将链表分成两部分
	cur := head
	var preNode *ListNode
	step := 0

	leftHead := head
	for cur != nil {
		step++

		// 说明当前结点是新的头结点
		if step == (length - remainder + 1) {
			if preNode != nil {
				preNode.Next = nil
			}

			newHead = cur
			break
		}

		preNode = cur
		cur = cur.Next
	}

	curr := newHead
	for curr.Next != nil {
		curr = curr.Next
	}

	// 最后把两部分链表连接起来
	curr.Next = leftHead

	return newHead
}
