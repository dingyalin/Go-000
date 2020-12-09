
# Goroutine

- 把并发交给调用者
- 知道goroutine什么时候能够退出
- 控制goroutine什么时候能够退出



## 实践


```
```

# Memory model

- https://golang.org/ref/mem

## Happen-Before

## MESI

cache line

内存屏障


原子性

可见性

双写 mysql double wirte buffer

nginx memory barires


# sync

原子性
可见性

最晚加锁 最早释放 锁间轻量

atomic.Value  Copy-On-Write

尾递归优化

Barging
Handsoff
Spinning

Barging + Spinning + 饥饿

pause intel pause

fast pause
slow pause


Fork-join MapReduce

errgroup  https://pkg.go.dev/golang.org/x/sync/errgroup

 局部变量+闭包

kratos

sync.Pool 高频的临时对象 (Request-Driven) ring buffer 双向链表

# chan
Overload

发送者close chan   [发送 nil??]


master workers

fan out

kafka --> fan out  至少一次


go interface nil

# context
上下文元数据 取消信号 超时控制

DDD


