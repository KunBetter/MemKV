// example
package main

import (
	"fmt"
	"github.com/KunBetter/MemKV"
	"time"
)

func main() {
	kv := MemKV.DB()
	//put
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
	//get
	fmt.Println(kv.Get([]byte("a")))
	//update
	kv.Put([]byte("a"), "e")
	fmt.Println(kv.Get([]byte("a")))
	//delete
	kv.Del([]byte("a"))
	fmt.Println(kv.Get([]byte("a")))
}
