// Utils
package MemKV

import ()

func (kv *MemKV) Hash32(d []byte) (hash, bucketID uint32) {
	hash = kv.HashFunc(d)
	hf := uint64(hash) * 2654435769
	bucketID = uint32(hf) & MOD
	return
}
