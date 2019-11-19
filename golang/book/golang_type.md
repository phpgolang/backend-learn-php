# golang中type常见用法

golang中type是一个非常重要的关键字，常见用法是定义结构，接口等。但是type中还有很多其他用法，如下：  

- 定义结构
  
> 衍生类型的方法和原类型的方法互相不相关

```go
type Person struct{
    name string
    age int
}

type Mutex struct{}
type OtherMutex Mutex //定义新的类型

func (m *Mutex) Lock(){
    println("Lock")
}
func (m *Mutex) UnLock(){
    println("UnLock")
}

func main(){
    m := Mutex{}
    m.Lock()   //OtherMutex 不具有Lock 和UnLock
}

```

- 定义接口

> 只要一个类型包含一个接口的所有方法，就可以说这个类型实现了该接口  
> 一个类型可以同时实现多个接口  
> 所有类型都实现了空接口
  
```go
// 定义cbs为空接口
var cbs interface{}

type Persioner interface{
    Speak(s string)
}

type Teacher struct {
    Position string
}
func (t Teacher) Speak(s string){
    t.Position = s
    println("我是老师： "+s)
}

type Student struct {
    Position string
}
func (t Student) Speak(s string){
    t.Positon = s
    println("我是学生： "+s)
}
```

- 类型定义

```go
type Myint int  //定义一个新的类型，

//定义一个类型方法
func (m Myint) showValue() {
    fmt.Println("show int", m)
}

func main() {
    var m Myint = 9 //变量声明
    m.showValue()
}
```

>新定义的类型，可以定义方法，如上例的 showValue()

- 定义函数类型

```go
type cb func(s string)  // 定义一个函数类型

//对函数类型再定义方法
func (f cb) ServerCb() error{
    f("cbb")
    fmt.Println("server cb")
    return nil
}

func main() {
    s :=cb(func( s string){
        fmt.Println("sss", s)
    })
    s.ServerCb()
}
```

## 参考

> [golang中type常用用法](https://www.cnblogs.com/smartrui/p/11425822.html)
