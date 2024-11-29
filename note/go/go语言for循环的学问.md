```go
for i:=0;i<10;i++{
	
}
go1.22后,变量i是否每次都重新创建
```

```go
s:=make([]int,10)
for i:=0;i<len(s);i++{
	
}
每次循环是否都会调用len(s),效率是否比下面这种低
l:=len(s)
for i:=0;i<l;i++{
	
}
是否比使用range低
for i:=range s{}
和下面是否一致
for i:=range len(s){
	
}

```

```go


```