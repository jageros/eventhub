package eventhub

import "log"

var (
	maxSeq    = 0
	listeners = make(map[int][]*listener) // map[事件id][]handleFunc
)

type listener struct {
	seq     int
	handler func(args ...interface{})
}

/****************************************************************
 * func：监听事件函数
 * eventID： 事件id
 * handler： 事件函数
 * return seq ： 该事件当前handle的序列号， 用于取消监听时参数
 ***************************************************************/
func Subscribe(eventID int, handler func(args ...interface{})) (seq int) {
	maxSeq++
	seq = maxSeq
	ln := &listener{
		seq:     seq,
		handler: handler,
	}
	ls, ok := listeners[eventID]
	if !ok {
		ls = []*listener{}
	}
	listeners[eventID] = append(ls, ln)
	return seq
}

/*******************************************
 * func： 发布事件
 * eventID 事件id
 * args 给handle传的参数
 * return nil
 *******************************************/
func Publish(eventID int, args ...interface{}) {
	ls, ok := listeners[eventID]
	if !ok {
		return
	}
	for _, l := range ls {
		err := catchPanic(func() {
			l.handler(args...)
		})
		if err != nil {
			log.Printf("eventId=%d, seq=%d, handle err=%v\n", eventID, l.seq, err)
		}
	}
}

/*************************************
 * func： 取消监听
 * eventID 事件ID
 * seq 事件序列号 （ps：因为一个事件可能存在多个handle函数）
 * return nil
 *************************************/
func Unsubscribe(eventID int, seq int) {
	ls, ok := listeners[eventID]
	if !ok {
		return
	}
	index := -1
	for i, l := range ls {
		if l.seq == seq {
			index = i
			break
		}
	}
	if index >= 0 {
		listeners[eventID] = append(ls[:index], ls[index+1:]...)
	}
}

/*******************************************
 * func： 对执行的handle捕获异常并恢复
 * f  需要执行的函数
 * return err 执行f返回的错误
 *******************************************/
func catchPanic(f func()) (err interface{}) {
	defer func() {
		err = recover()
		if err != nil {
			log.Printf("CatchPanic panic: 【%s】\n", err)
		}
	}()

	f()
	return
}
