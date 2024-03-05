package main

import (
 "fmt"
 "sync"
)

//后面开启了1000个goroutine，所以这里channel的缓冲区设置成了1000
var ch = make(chan bool, 1000)

type Product struct {
 sync.Mutex
 stock int64 // 商品库存
}

func main() {
 p := Product{stock: 1000}
 for i := 1; i <= 1000; i++ {
  go p.handleRequest(i)
 }
 <-ch
}

func (p *Product) handleRequest(user int) {
 if p.tryAcquireLock() {
  if p.stock > 0 {
   // 执行秒杀逻辑
   p.stock--
   fmt.Printf("用户%d秒杀成功，剩余库存：%d\n", user, p.stock)
  } else {
   fmt.Printf("用户%d秒杀失败，库存不足\n", user)
  }
  //这里是不可以使用defer的，因为可能会加锁失败，unlock一个不存在的锁
  p.releaseLock()
 } else {
  fmt.Printf("用户%d未获取到锁，秒杀失败\n", user)
 }
}

func (p *Product) tryAcquireLock() bool {
//p.TryLock() 方法用于尝试获取锁，如果成功获取到锁，则相当于执行了 Lock() 操作，即加锁成功。 
 return p.TryLock()
}

func (p *Product) releaseLock() {
 p.Unlock()
 ch <- true
}
