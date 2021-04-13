<!-- TOC depthFrom:1 depthTo:2 -->

- [1. 逃逸分析](#1-逃逸分析)
    - [1.1. 分析命令](#11-分析命令)
    - [1.2. 逃逸场景（什么情况才分配到堆中）](#12-逃逸场景什么情况才分配到堆中)
    - [1.3. 参考链接](#13-参考链接)

<!-- /TOC -->

# 1. 逃逸分析

- 逃逸分析的好处是为了减少GC的压力，不逃逸的对象分配在栈上，当函数返回时就回收了资源，不需要GC标记清除。
- 逃逸分析完后可以确定哪些变量可以分配在栈上，栈的分配比堆快，性能好（逃逸的局部变量会分配在堆上，而没有发送逃逸的则有编辑器分配到栈上）
- 同步消除，如果定义的对象在方法上有同步锁，但在运行时，却只有一个线程在访问，此时逃逸分析的机器码，发去掉同步锁进行。

## 1.1. 分析命令
```go
go run -gcflags "-m -l" (-m打印逃逸分析信息，-l禁止内联编译)

/*
testProj go run -gcflags "-m -l" internal/test1/main.go
# command-line-arguments
internal/test1/main.go:4:2: moved to heap: a
internal/test1/main.go:5:11: main make([]*int, 1) does not escape
*/

go tool compile -S main.go | grep runtime.newobject（汇编代码中搜runtime.newobject指令，该指令用于生成堆对象）

/*
testProj go tool compile -S internal/test1/main.go | grep newobject
        0x0028 00040 (internal/test1/main.go:4) CALL    runtime.newobject(SB)
*/
```

* 申请到`栈内存`好处：函数返回直接释放，不会引起垃圾回收，对性能没有影响。
* 申请到`堆上面的内存`才会引起垃圾回收，如果这个过程（特指垃圾回收不断被触发）过于高频就会导致 gc 压力过大，程序性能出问题。

## 1.2. 逃逸场景（什么情况才分配到堆中）

    1. 指针逃逸
    2. 栈空间不足逃逸（空间开辟过大）
    3. 动态类型逃逸（不确定长度大小）
    4. 闭包引用对象逃逸

### 1.2.1.以下情况一定发生指针逃逸
- 在某个函数中new或字面量创建出的变量，将其指针作为函数返回值，则该变量一定发生逃逸（构造函数返回的指针变量一定逃逸）；
- `被已经逃逸的变量引用的指针`，一定发生逃逸。
- `被指针类型的slice、map和chan引用的指针`一定发生逃逸

备注：stack overflow上有人提问为什么使用指针的chan比使用值的chan慢30%，答案就在这里：使用指针的chan发生逃逸，gc拖慢了速度。[问题地址](https://stackoverflow.com/questions/41178729/why-passing-pointers-to-channel-is-slower)

### 1.2.2.一些必然不会逃逸的情况

- 指针被未发生逃逸的变量引用；
- 仅仅在函数内对变量做取址操作，而未将指针传出；

### 1.2.3.有一些情况可能发生逃逸，也可能不会发生逃逸

- 将指针作为入参传给别的函数；这里还是要看指针在被传入的函数中的处理过程，如果发生了上边的三种情况，则会逃逸；否则不会逃逸；

## 1.3. 参考链接

- [golang 逃逸分析详解](https://zhuanlan.zhihu.com/p/91559562)
- [Golang内存分配逃逸分析](https://driverzhang.github.io/post/golang%E5%86%85%E5%AD%98%E5%88%86%E9%85%8D%E9%80%83%E9%80%B8%E5%88%86%E6%9E%90/)
- [Go的sync](https://driverzhang.github.io/post/go%E7%9A%84sync.pool%E4%B8%B4%E6%97%B6%E5%AF%B9%E8%B1%A1%E6%B1%A0/)