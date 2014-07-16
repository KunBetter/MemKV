// Bucket
package MemKV

type Bucket struct {
	Elems []*BST
	Size  int
}

func NewBucket(size int) *Bucket {
	b := &Bucket{
		Size:  size,
		Elems: make([]*BST, size),
	}
	for i := 0; i < size; i++ {
		b.Elems[i] = BSTer()
	}
	return b
}

func (b *Bucket) Put(key []byte, value interface{}, keyHash, bucketID uint32) {
	b.Elems[bucketID].Add(keyHash, &KVNode{Key: key, Value: value, Next: nil})
}

func (b *Bucket) Get(key []byte, keyHash, bucketID uint32) interface{} {
	return b.Elems[bucketID].Get(keyHash, &KVNode{Key: key, Value: nil, Next: nil})
}

func (b *Bucket) Del(key []byte, keyHash, bucketID uint32) {
	b.Elems[bucketID].Del(keyHash, &KVNode{Key: key, Value: nil, Next: nil})
}
