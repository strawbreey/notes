---
title: "Leetcode 0002 Add Two Numbers"
date: 2020-11-26T00:10:42+08:00
draft: false
---

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

Solution
```js
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    let node = new ListNode(0)
    let step = 0
    let curNode = node

    while(step || l1 || l2) {
        let val1 = l1 ? l1.val : 0
        let val2 = l2 ? l2.val : 0
        
        node.next = new ListNode((val1 + val2 + step) % 10)
        node = node.next
        step = Math.floor((val1 + val2 + step) / 10)
        l1 = l1 ?  l1.next : null
        l2 = l2 ? l2.next : null
    }
    return curNode.next
};
```

### 参考资料

- [两数相加](https://leetcode-cn.com/problems/add-two-numbers/)