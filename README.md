MemKV
==========
实现一个基于内存的KV存储微系统。  
利用链表法解决key值哈希冲突问题。  
系统有两部分组成：  
1 哈希bucket。  
2 链表。

斐波那契（Fibonacci）散列法
-----
1 对于16位整数而言，这个乘数是40503  
2 对于32位整数而言，这个乘数是2654435769  
3 对于64位整数而言，这个乘数是11400714819323198485  
对32位整数，示例公式:  
	index = (value * 2654435769) >> 28
	
Requires
-----
go get github.com/spaolacci/murmur3

Usage
-----
```go
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
```