
```go

type EqualKey[T compare] interface {
	EqualKey() T
}
```

```go

func HasCoincideByKey[S ~[]E, E EqualKey[T],T comparable](s1, s2 S) bool {
	return ture
}
```

```go

func HasCoincideByKey[S ~[]EqualKey[T],T comparable](s1, s2 S) bool {
	return ture
}
```
上面是两种完全不同的结果,第一种函参数是具体类型切片,第二种函数参数是接口切片