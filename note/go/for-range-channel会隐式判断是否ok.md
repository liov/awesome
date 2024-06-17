```go
ch :=make(chan int)
for v:=range ch{
}
```
当chan 关闭时，range 会自动退出,否则阻塞