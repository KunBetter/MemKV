// MemKV
package MemKV

import (
	"github.com/spaolacci/murmur3"
)

const BUCKETSIZT = 0x0000FFFF
const REHASH = 1
const SATURATIONTHRESHOLD = 0.92

type MemKV struct {
	HashFunc  func([]byte) uint32
	Bucket    []*Bucket //2 bucket
	B1Size    uint32    //bucket1 size
	ReHashing bool
	Signal    chan int
}

func DB() *MemKV {
	kv := &MemKV{
		HashFunc:  murmur3.Sum32,
		Bucket:    make([]*Bucket, 2),
		B1Size:    BUCKETSIZT,
		ReHashing: false,
		Signal:    make(chan int),
	}
	kv.Bucket[0] = NewBucket(kv.B1Size + 1)
	kv.Bucket[1] = nil
	go kv.process()
	return kv
}

func (kv *MemKV) process() {
	for {
		select {
		case sig := <-kv.Signal:
			if sig == REHASH {
				kv.ReHashing = true
				kv.Bucket[1] = NewBucket(kv.B1Size*2 + 1)
				go kv.Rehash()
			}
		}
	}
}

func (kv *MemKV) Rehash() {

}

func (kv *MemKV) Hash32(d []byte, mod uint32) (keyHash, bucketID uint32) {
	keyHash = kv.HashFunc(d)
	hf := uint64(keyHash) * 2654435769
	bucketID = uint32(hf) & mod
	return
}

func (kv *MemKV) Put(key []byte, value interface{}) {
	if kv.ReHashing {
		keyHash, bucketID := kv.Hash32(key, kv.B1Size*2)
		kv.Bucket[1].Put(key, value, keyHash, bucketID)
	} else {
		keyHash, bucketID := kv.Hash32(key, kv.B1Size)
		kv.Bucket[0].Put(key, value, keyHash, bucketID)
		//rehash?
		if kv.Bucket[0].Saturation >= SATURATIONTHRESHOLD {
			kv.Signal <- REHASH
		}
	}
}

func (kv *MemKV) Get(key []byte) (value interface{}) {
	value = nil
	if kv.ReHashing {
		keyHash, bucketID := kv.Hash32(key, kv.B1Size*2)
		value = kv.Bucket[1].Get(key, keyHash, bucketID)
	}
	if value == nil {
		keyHash, bucketID := kv.Hash32(key, kv.B1Size)
		value = kv.Bucket[0].Get(key, keyHash, bucketID)
	}
	return value
}

func (kv *MemKV) Del(key []byte) {
	if kv.ReHashing {
		keyHash, bucketID := kv.Hash32(key, kv.B1Size*2)
		kv.Bucket[1].Del(key, keyHash, bucketID)
	}
	keyHash, bucketID := kv.Hash32(key, kv.B1Size)
	kv.Bucket[0].Del(key, keyHash, bucketID)
}
