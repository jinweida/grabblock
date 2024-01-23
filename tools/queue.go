package tools

import (
	"time"
)

//先进先出队列
//type Queue struct {
//	data *list.List
//	lock *sync.RWMutex
//}

//func NewQueue() *Queue {
//	return &Queue{data: list.New(), lock: new(sync.RWMutex)}
//}
//
////入队列
//func (q *Queue) Push(v interface{}) {
//	q.lock.Lock()
//	q.data.PushFront(v)
//	defer q.lock.Unlock()
//}
//
////出队列
//func (q *Queue) Pop() (interface{}, bool) {
//	q.lock.Lock()
//	if q.data.Len() > 0 {
//		elm := q.data.Back()
//		v := elm.Value
//		q.data.Remove(elm)
//		defer q.lock.Unlock()
//		return v, true
//	}
//	defer q.lock.Unlock()
//	return nil, false
//}
//
////队列长度
//func (q *Queue) Qsize() int {
//	defer q.lock.RUnlock()
//	q.lock.RLock()
//	return q.data.Len()
//}
//
////判断是否为空
//func (q *Queue) IsEmpty() bool {
//	defer q.lock.RUnlock()
//	q.lock.RLock()
//	return !(q.data.Len() > 0)
//}
type DataContainer struct {
	Queue chan interface{}
}

func NewDataContainer(max_queue_len int) (dc *DataContainer) {
	dc = &DataContainer{}
	dc.Queue = make(chan interface{}, max_queue_len)
	return dc
}

//非阻塞push
func (dc *DataContainer) Push(data interface{}, waittime time.Duration) bool {
	click := time.After(waittime)
	select {
	case dc.Queue <- data:
		return true
	case <-click:
		return false
	}
}

//非阻塞pop
func (dc *DataContainer) Pop(waittime time.Duration) (data interface{}) {
	click := time.After(waittime)
	select {
	case data = <-dc.Queue:
		return data
	case <-click:
		return nil
	}
}
