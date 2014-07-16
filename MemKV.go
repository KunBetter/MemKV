// MemKV
package MemKV

import (
	"github.com/spaolacci/murmur3"
	"runtime"
)

const MOD = 0x000FFFFF

type MemKV struct {
	HashFunc func([]byte) uint32
	Bucket   *Bucket
}

func DB() *MemKV {
	runtime.GOMAXPROCS(runtime.NumCPU())
	kv := &MemKV{
		HashFunc: murmur3.Sum32,
		Bucket:   NewBucket(MOD + 1),
	}
	return kv
}

func (kv *MemKV) Hash32(d []byte) (keyHash, bucketID uint32) {
	keyHash = kv.HashFunc(d)
	hf := uint64(keyHash) * 2654435769
	bucketID = uint32(hf) & MOD
	return
}

func (kv *MemKV) Put(key []byte, value interface{}) {
	keyHash, bucketID := kv.Hash32(key)
	kv.Bucket.Put(key, value, keyHash, bucketID)
}

func (kv *MemKV) Get(key []byte) interface{} {
	keyHash, bucketID := kv.Hash32(key)
	return kv.Bucket.Get(key, keyHash, bucketID)
}

func (kv *MemKV) Del(key []byte) {
	keyHash, bucketID := kv.Hash32(key)
	kv.Bucket.Del(key, keyHash, bucketID)
}
