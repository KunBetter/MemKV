// KVList
package MemKV

import (
	"fmt"
)

func BytesEqual(a, b []byte) bool {
	aLen := len(a)
	bLen := len(b)
	if aLen != bLen {
		return false
	}
	for i := 0; i < aLen; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type KVNode struct {
	Key   []byte
	Value interface{}
	Next  *KVNode
}

type KVList struct {
	Head   *KVNode
	Length uint32
}

func NewKVNode(key []byte, value interface{}) *KVNode {
	n := &KVNode{
		Key:   key,
		Value: value,
		Next:  nil,
	}
	return n
}

func NewKVList() *KVList {
	l := &KVList{
		Head:   NewKVNode([]byte("HEAD"), "HEAD"),
		Length: 0,
	}
	return l
}

func (l *KVList) Len() uint32 {
	return l.Length
}

func (l *KVList) Find(n *KVNode) (*KVNode, *KVNode, bool) {
	head := l.Head
	forward := l.Head
	for head.Next != nil {
		forward = head
		head = head.Next
		if BytesEqual(head.Key, n.Key) {
			return forward, head, true
		}
	}
	return forward, head, false
}

func (l *KVList) Print() {
	head := l.Head
	for head.Next != nil {
		head = head.Next
		fmt.Printf("<%s,%v>", string(head.Key), head.Value)
	}
	fmt.Println()
}

func (l *KVList) Add(n *KVNode) (add uint32) {
	_, cur, ok := l.Find(n)
	if ok {
		//update
		cur.Value = n.Value
		add = 0
	} else {
		cur.Next = n
		l.Length++
		add = 1
	}
	return
}

func (l *KVList) Del(n *KVNode) (del uint32) {
	del = 0
	if l.Length <= 0 {
		return
	}
	p, cur, ok := l.Find(n)
	if ok {
		p.Next = cur.Next
		l.Length--
		del = 1
	}
	return
}

func (l *KVList) Get(n *KVNode) *KVNode {
	if l.Length <= 1 {
		return l.Head.Next
	}
	_, cur, ok := l.Find(n)
	if !ok {
		return nil
	}
	return cur
}
