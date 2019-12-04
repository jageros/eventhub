# eventhub
Subscribe/Publish event in program.

## 这是什么？
+ 这是一个开源的用于golang程序内部发布和监听事件的package，其原理是通过定义一个全局变量存储监听的事件id和对应的handle函数，当发布事件时，通过事件id查找对应的handle，如果找到则执行该handle函数。

## 使用说明
### 安装
``go get github.com/jageros/eventhub``

### 使用
#### 监听事件
+ ``seq := eventhub.Subscribe(eventID, handle)``
+ eventID为事件id， 返回的seq为序列号，两者都是int型， handle原型为func(args ...interface{}){}，args参数类型为interface{}型，通过事件发布函数传入

#### 发布事件
+ ``eventhub.Publish(eventID, arg1, arg2 ··· )``
+ eventID为事件id， argx为参数，将传给事件监听的handle

#### 取消监听
+ ``Unsubscribe(eventID int, seq int)``
+ eventID为事件id， seq为序列号，序列从监听事件函数返回值获得

### 使用示例
``github.com/jageros/eventhub/example/main.go``
```go
package main

import (
	"fmt"
	"github.com/jageros/eventhub"
)

func main() {
	// 监听事件
	eventhub.Subscribe(2, func(args ...interface{}) {
		fmt.Printf("Subscribe1 eventId=2 args=%v\n", args)
	})
	eventhub.Subscribe(1, func(args ...interface{}) {
		fmt.Printf("Subscribe2 eventId=1 args=%v\n", args)
	})
	eventhub.Subscribe(3, func(args ...interface{}) {
		fmt.Printf("Subscribe3 eventId=3 args=%v\n", args)
		if arg, ok := args[0].(func()); ok {
			arg()
		}
	})

	// 监听并取消监听
	seq := eventhub.Subscribe(1, func(args ...interface{}) {
		fmt.Printf("Subscribe4 eventId=1 args=%+v\n", args)
	})
	eventhub.Unsubscribe(1, seq)

	// 发布事件
	eventhub.Publish(1, 10, 100)
	eventhub.Publish(2, 20, 200)
	eventhub.Publish(3, test)
}

// 此函数用作参数
func test() {
	fmt.Printf("End!\n")
}
```

