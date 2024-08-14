如果是简单结构体
```go
type Foo struct {}
func New() *Foo{
return &Foo{}
}
```
如果是带配置的用于数据操作的结构体
分为一段式生成和二段式
一段式，直接生成,这时Option直接定义在该结构体上
这种适用于结构体本身就是配置属性的,大部分字段都是基本类型,只是带了操作方法
```go
type Server struct {
	Port int
}
type Option func(*Server) // type Option interface{Apply(*Server)}
func New(opts ...Option) *Server {
	s:=&Server{}
	for _, opt := range opts {
		opt(s)
   }
   return s
}
```
二段式生成,即有两个定义，Config 和 由Config 生成的 结构体
这种适用于结构体本身不是单纯的配置,比如含有可变的数据结构,chan,cache,map等
此时Option定义在Config上
```go
type Config struct {
	Cap int
}
func (c *Config) New() *Cache {
	return &Cache{
		data: make([]int,0, c.Cap),
    }
}

type Option func(*Config)
func NewConfig(opts ...Option) *Config {
    c:=&Config{}
    for _, opt := range opts {
        opt(s)
    }
	return c
}

type Cache struct {
	data []int
}

func New(opts ...Option) *Cache {
    c:=NewConfig(opts...)
	return c.New()
}
```