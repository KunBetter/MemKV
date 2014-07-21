// BST
// binary search tree
package MemKV

import (
	"sync"
)

type BSTNode struct {
	Key         uint32
	Value       *KVList
	Left, Right *BSTNode
}

func NewBSTNode(key uint32, value *KVNode) *BSTNode {
	n := &BSTNode{
		Key:   key,
		Value: NewKVList(),
		Left:  nil,
		Right: nil,
	}
	n.Value.Add(value)
	return n
}

type BST struct {
	Root   *BSTNode
	NCount uint32 //node num
	RWLock *sync.RWMutex
}

func BSTer() *BST {
	return &BST{
		Root:   nil,
		NCount: 0,
		RWLock: new(sync.RWMutex),
	}
}

func (n *BSTNode) Find(parent *BSTNode, key uint32) (p, cur *BSTNode) {
	if key < n.Key {
		if n.Left == nil {
			return parent, n
		} else {
			return n.Left.Find(n, key)
		}
	} else if key > n.Key {
		if n.Right == nil {
			return parent, n
		} else {
			return n.Right.Find(n, key)
		}
	} else {
		return parent, n
	}
}

func (t *BST) Find(key uint32) (p, cur *BSTNode) {
	return t.Root.Find(t.Root, key)
}

func (t *BST) Add(key uint32, value *KVNode) {
	t.RWLock.Lock()
	defer t.RWLock.Unlock()
	var add uint32 = 1
	if t.Root == nil {
		t.Root = NewBSTNode(key, value)
		return
	}
	_, n := t.Find(key)
	if key < n.Key {
		n.Left = NewBSTNode(key, value)
	} else if key > n.Key {
		n.Right = NewBSTNode(key, value)
	} else {
		add = n.Value.Add(value)
	}
	t.NCount += add
}

func (n *BSTNode) LeftMax() (p, cur *BSTNode) {
	nLeft := n.Left
	p = nLeft
	for nLeft.Right != nil {
		p = nLeft
		nLeft = nLeft.Right
	}
	cur = nLeft
	return
}

func (t *BST) Get(key uint32, value *KVNode) interface{} {
	t.RWLock.RLock()
	defer t.RWLock.RUnlock()
	if t.Root == nil {
		return nil
	}
	_, n := t.Find(key)
	if key != n.Key {
		return nil
	} else {
		return n.Value.Get(value).Value
	}
	return nil
}

func (t *BST) Del(key uint32, value *KVNode) {
	t.RWLock.Lock()
	defer t.RWLock.Unlock()
	if t.Root == nil {
		return
	}
	np, n := t.Find(key)
	if n.Key == key {
		if n.Value.Length > 1 {
			t.NCount -= n.Value.Del(value)
			return
		}
		t.NCount -= 1
		if n.Left != nil && n.Right != nil {
			p, cur := n.LeftMax()
			n.Key = cur.Key
			n.Value = cur.Value
			if p.Key == cur.Key {
				n.Left = nil
			} else {
				p.Right = cur.Left
			}
		} else {
			if n.Key < np.Key {
				if n.Left == nil {
					np.Left = n.Right
				} else {
					np.Left = n.Left
				}
			} else if n.Key > np.Key {
				if n.Left == nil {
					np.Right = n.Right
				} else {
					np.Right = n.Left
				}
			} else {
				if n.Left == nil {
					t.Root = n.Right
				} else {
					t.Root = n.Left
				}
			}
		}
	}
}
