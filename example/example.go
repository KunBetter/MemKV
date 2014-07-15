// example
package main

import (
	"fmt"
	"github.com/KunBetter/MemKV"
	"time"
)

func main() {
	kv := MemKV.MEMKV()
	kv.Put([]byte("a"), "b")
	kv.Put([]byte("c"), "d")
	s := time.Now()
	for i := 0; i < 100000000; i++ {
		kv.Get([]byte("c"))
	}
	e := time.Now()
	t := e.UnixNano() - s.UnixNano()
	var es float64 = float64(t)
	fmt.Println(kv.Get([]byte("c")), t, es/1e9)
}
