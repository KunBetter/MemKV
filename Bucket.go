// Bucket
package MemKV

type KVNode struct {
	Key   []byte
	Value interface{}
	KHash uint32
	Next  *KVNode
}

type KVList struct {
	Head *KVNode
}

type Bucket struct {
	Nodes []*KVList
	Size  int
}

func NewKVNode(key []byte, value interface{}, kh uint32) *KVNode {
	n := &KVNode{
		Key:   key,
		Value: value,
		KHash: kh,
		Next:  nil,
	}
	return n
}

func NewKVList() *KVList {
	l := &KVList{
		Head: NewKVNode([]byte(""), "HEAD", uint32(1)),
	}
	return l
}

func NewBucket(size int) *Bucket {
	b := &Bucket{
		Size:  size,
		Nodes: make([]*KVList, size+1),
	}
	for i := 0; i < size+1; i++ {
		b.Nodes[i] = NewKVList()
	}
	return b
}

func (l *KVList) Add(n *KVNode) {
	head := l.Head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = n
}

func (l *KVList) Get(kh uint32) *KVNode {
	head := l.Head
	for head.Next != nil {
		head = head.Next
		if head.KHash == kh {
			return head
		}
	}
	return nil
}

func (b *Bucket) Put(key []byte, value interface{}, kh, bucketID uint32) {
	n := NewKVNode(key, value, kh)
	b.Nodes[bucketID].Add(n)
}

func (b *Bucket) Get(key []byte, kh, bucketID uint32) interface{} {
	n := b.Nodes[bucketID].Get(kh)
	if n == nil {
		return nil
	} else {
		return n.Value
	}
}
