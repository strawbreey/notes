---
title: "Leetcode 0147 Insertion Sort List"
date: 2020-11-20T09:45:27+08:00
draft: false
tags: ['leetcode']
---

Sort a linked list using insertion sort.

![Insertion-sort-example](/images/Insertion-sort-example-300px.gif) 

A graphical example of insertion sort. The partial sorted list (black) initially contains only the first element in the list.
With each iteration one element (red) is removed from the input data and inserted in-place into the sorted list

Algorithm of Insertion Sort:

  - Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
  - At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
  - It repeats until no input elements remain.

对链表进行插入排序。


插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

![Insertion-sort-example](/images/Insertion-sort-example-300px.gif) 

插入排序算法：

- 插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
- 每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
- 重复直到所有输入数据插入完为止。


Example: 

```
输入: 4->2->1->3
输出: 1->2->3->4

输入: -1->5->3->4->0
输出: -1->0->3->4->5
```

Solution: 

```js
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var insertionSortList = function(head) {

    if (head === null) {
        return head
    }
    const newLinkedList = new ListNode(0)

    newLinkedList.next = head

    let last = head // 上一个节点
    let cur = head.next // 当前节点

    while (cur !== null) {
        if (last.val <= cur.val) {
            last = last.next
        } else {
            let prev = newLinkedList
            while ( prev.next.val <= cur.val) {
                prev = prev.next
            }
            last.next = cur.next
            cur.next = prev.next
            prev.next = cur
        }
        cur = last.next
    }
    return newLinkedList.next
};
```


### 参考链接 

- [对链表进行插入排序](https://leetcode-cn.com/problems/insertion-sort-list/)
- [Insertion Sort List](https://leetcode.com/problems/insertion-sort-list/)
