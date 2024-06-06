```go
var ch chan int
// 这种操作会阻塞
<-ch
// 这种操作会阻塞
ch<-1
```
所以使用select,在chan进入case前,将chan置为nil可以避免选中该case,变相实现条件分支
```go
var ch chan int
select{
case <-ch:
case ch<-1:
default:
}
}

```
