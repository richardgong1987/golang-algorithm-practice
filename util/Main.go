package util

import "fmt"

func Main() {
	{
		instance := Get_MaxPriorityQueue_Instance()
		instance.enQueue(3)
		instance.enQueue(5)
		instance.enQueue(10)
		instance.enQueue(2)
		instance.enQueue(7)

		//for v := instance.deQueue(); v != 0; v = instance.deQueue() {
		//	fmt.Println(v)
		//}
	}
	{
		instance := Get_MinPriorityQueue_Instance()
		instance.enQueue(3)
		instance.enQueue(5)
		instance.enQueue(10)
		instance.enQueue(2)
		instance.enQueue(7)

		for v, err := instance.deQueue(); err == nil; v, err = instance.deQueue() {
			fmt.Println(v)
		}

	}
}
