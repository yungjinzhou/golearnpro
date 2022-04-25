package testing

import (
	"fmt"
	"sync"
	"testing"
)

// sysnc.Pool
//适用于通过复用，降低复杂对象的创建和GC代价
//协程安全，会有锁的开销
//生命周期受GC影响，不适合于做连接池等，需要自己管理生命周期的资源的池化



// 对象的获取流程
//●尝试从私有对象获取
//●私有对象不存在，尝试从当前Processor的共享池获取
//●如果当前Processor共享池也是空的，那么就尝试去其他Processor的共享池获取
//●如果所有子池都是空的，最后就用用户指定的New函数产生一个新的对象返回

// 对象的放回
//●如果私有对象不存在则保存为私有对象
//●如果私有对象存在，放入当前Processor子池的共享池中





func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
	New：func() interface{} {
	fmt.Println("Create a new object.")
	return 100
	},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC() // GC 会清除sync.poll中的缓存对象，对象的缓存有效期为下一次GC之前
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}






//
//func TestSyncPolInMultiGroutine(t *testing.T) {
//	pool := &sync.Pool{
//		New：func() interface{} {
//			fmt.Println("Create a new object.")
//			return 10
//	},
//	}
//
//
//	pool.Put(100)
//	pool.Put(100)
//	pool.Put(100)
//
//
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(id int) {
//			t.Log(pool.Get())
//			wg.Done()
//	}(i)
//	}
//	wg.Wait()
//}