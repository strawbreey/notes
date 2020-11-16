---
title: "Redis Design and Implementation"
date: 2020-11-05T15:35:26+08:00
draft: false
---

### 链表和链表节点的实现

[redis 链表实现](https://github.com/redis/redis/blob/unstable/src/adlist.h)

```c++
typedef struct listNode {
  struct listNode * prev;
  struct listNode * next;
  void * value;
} listNode;
```

多个listNode可以通过prev和next指针组成双端链表

```c++
typedef struct list {
    // 表头节点
    listNode *head; 
    // 表尾节点
    listNode *tail;
    // 节点复制函数
    void *(*dup)(void *ptr);
    // 节点释放函数
    void (*free)(void *ptr);
    // 节点值对比函数
    int (*match)(void *ptr, void *key);
    // 链表所包含的节点数量
    unsigned long len;
} list;
```

list结构为链表提供了表头指针head、表尾指针tail，以及链表长度计数器len，而dup、free和match成员则是用于实现多态链表所需的类型特定函数：

链表被广泛用于实现Redis的各种功能，比如列表键、发布与订阅、慢查询、监视器等。

❑每个链表节点由一个listNode结构来表示，每个节点都有一个指向前置节点和后置节点的指针，所以Redis的链表实现是双端链表。

❑每个链表使用一个list结构来表示，这个结构带有表头节点指针、表尾节点指针，以及链表长度等信息。

❑因为链表表头节点的前置节点和表尾节点的后置节点都指向NULL，所以Redis的链表实现是无环链表。

❑通过为链表设置不同的类型特定函数，Redis的链表可以用于保存各种不同类型的值。


### 字典

字典，又称为符号表（symbol table）、关联数组（associative array）或映射（map），是一种用于保存键值对（key-value pair）的抽象数据结构。

在字典中，一个键（key）可以和一个值（value）进行关联（或者说将键映射为值），这些关联的键和值就称为键值对。字典中的每个键都是独一无二的，程序可以在字典中根据键查找与之关联的值，或者通过键来更新值，又或者根据键来删除整个键值对，等等。


```c++
typedef struct dictht {
  dictEnter **table;
  unsigned long size;
  unsigned long siemask;
  unsigned long used;
} dictht;
```

❑字典被广泛用于实现Redis的各种功能，其中包括数据库和哈希键。

❑Redis中的字典使用哈希表作为底层实现，每个字典带有两个哈希表，一个平时使用，另一个仅在进行rehash时使用。

❑当字典被用作数据库的底层实现，或者哈希键的底层实现时，Redis使用MurmurHash2算法来计算键的哈希值。

❑哈希表使用链地址法来解决键冲突，被分配到同一个索引上的多个键值对会连接成一个单向链表。

❑在对哈希表进行扩展或者收缩操作时，程序需要将现有哈希表包含的所有键值对rehash到新哈希表里面，并且这个rehash过程并不是一次性地完成的，而是渐进式地完成的