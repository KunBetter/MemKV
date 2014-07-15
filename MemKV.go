// MemKV
package MemKV

import (
	"github.com/spaolacci/murmur3"
)

const MOD = 0x0000FFFF

type MemKV struct {
	HashFunc func([]byte) uint32
	Bucket   *Bucket
}

func MEMKV() *MemKV {
	kv := &MemKV{
		HashFunc: murmur3.Sum32,
		Bucket:   NewBucket(MOD + 1),
	}
	return kv
}

func (kv *MemKV) Put(key []byte, value interface{}) {
	hash, bucketID := kv.Hash32(key)
	kv.Bucket.Put(key, value, hash, bucketID)
}

func (kv *MemKV) Get(key []byte) interface{} {
	hash, bucketID := kv.Hash32(key)
	return kv.Bucket.Get(key, hash, bucketID)
}
