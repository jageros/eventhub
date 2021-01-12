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

func test() {
	fmt.Printf("End!\n")
}