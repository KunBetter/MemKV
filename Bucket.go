// Bucket
package MemKV

type Bucket struct {
	Elems      []*BST
	Size       uint32
	Blank      uint32
	Saturation float32 //饱和度
}

func NewBucket(size uint32) *Bucket {
	b := &Bucket{
		Size:  size,
		Blank: size,
		Elems: make([]*BST, size),
	}
	var i uint32 = 0
	for i = 0; i < size; i++ {
		b.Elems[i] = BSTer()
	}
	return b
}

func (b *Bucket) BucketSaturation() {
	b.Saturation = float32(b.Size-b.Blank) / float32(b.Size)
}

func (b *Bucket) Put(key []byte, value interface{}, keyHash, bucketID uint32) {
	ECount := b.Elems[bucketID].NCount
	b.Elems[bucketID].Add(keyHash, &KVNode{Key: key, Value: value, Next: nil})
	if ECount == 0 && b.Elems[bucketID].NCount > 0 {
		b.Blank--
		b.BucketSaturation()
	}
}

func (b *Bucket) Get(key []byte, keyHash, bucketID uint32) interface{} {
	return b.Elems[bucketID].Get(keyHash, &KVNode{Key: key, Value: nil, Next: nil})
}

func (b *Bucket) Del(key []byte, keyHash, bucketID uint32) {
	ECount := b.Elems[bucketID].NCount
	b.Elems[bucketID].Del(keyHash, &KVNode{Key: key, Value: nil, Next: nil})
	if ECount > 0 && b.Elems[bucketID].NCount == 0 {
		b.Blank++
		b.BucketSaturation()
	}
}
