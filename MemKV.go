// MemKV
package MemKV

import (
	"github.com/spaolacci/murmur3"
)

const BUCKETSIZT = 0x0000FFFF

type MemKV struct {
	HashFunc func([]byte) uint32
	Bucket   []*Bucket //2 bucket
	B1Size   uint32    //bucket1 size
	CurBIx   int       //current bucket
}

func DB() *MemKV {
	kv := &MemKV{
		HashFunc: murmur3.Sum32,
		Bucket:   make([]*Bucket, 2),
		B1Size:   BUCKETSIZT,
		CurBIx:   0,
	}
	kv.Bucket[0] = NewBucket(kv.B1Size + 1)
	kv.Bucket[1] = nil
	return kv
}

func (kv *MemKV) Hash32(d []byte, mod uint32) (keyHash, bucketID uint32) {
	keyHash = kv.HashFunc(d)
	hf := uint64(keyHash) * 2654435769
	bucketID = uint32(hf) & mod
	return
}

func (kv *MemKV) Put(key []byte, value interface{}) {
	keyHash, bucketID := kv.Hash32(key, kv.B1Size)
	kv.Bucket[kv.CurBIx].Put(key, value, keyHash, bucketID)
}

func (kv *MemKV) Get(key []byte) interface{} {
	keyHash, bucketID := kv.Hash32(key, kv.B1Size)
	return kv.Bucket[kv.CurBIx].Get(key, keyHash, bucketID)
}

func (kv *MemKV) Del(key []byte) {
	keyHash, bucketID := kv.Hash32(key, kv.B1Size)
	kv.Bucket[kv.CurBIx].Del(key, keyHash, bucketID)
}
