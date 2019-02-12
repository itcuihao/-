# 面试题13：在O(1)时间删除链表节点

## 题目

```
给定单向链表的头指针和一个节点指针，定义一个函数在O(1)时间删除该节点。
```

链表节点与函数定义：

```
type ListNode struct {
    Value int
    Node ListNode
}

func DeleteNode(pListHead, pToBeDeleted ListNode){}
```