package util

import "fmt"

func Main() {
	//{
	//instance := Get_MaxPriorityQueue_Instance()
	//instance.enQueue(3)
	//instance.enQueue(5)
	//instance.enQueue(10)
	//instance.enQueue(2)
	//instance.enQueue(7)

	//for v, err := instance.deQueue(); err == nil; v, err = instance.deQueue() {
	//	fmt.Println(v)
	//}
	//}
	//{
	//	instance := Get_MinPriorityQueue_Instance()
	//	instance.enQueue(3)
	//	instance.enQueue(5)
	//	instance.enQueue(10)
	//	instance.enQueue(2)
	//	instance.enQueue(7)
	//
	//	for v, err := instance.deQueue(); err == nil; v, err = instance.deQueue() {
	//		fmt.Println(v)
	//	}
	//
	//}

	//h := &IntHeap{}
	//heap.Init(h)
	//heap.Push(h, 3)
	//heap.Push(h, 1)
	//heap.Push(h, 5)
	//heap.Push(h, 2)
	//heap.Push(h, 8)
	//
	//fmt.Printf("minimum: %d\n", (*h)[0])
	//for h.Len() > 0 {
	//	fmt.Printf("%d ", heap.Pop(h))
	//}
	var a byte = 'a'
	var z byte = 'z'
	for ; a < z; a += 1 {
		fmt.Println(string(a))
	}
}
