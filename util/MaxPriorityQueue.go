package util

import "errors"

type MaxPriorityQueue struct {
	array []int
	size  int
}

func Get_MaxPriorityQueue_Instance() *MaxPriorityQueue {
	return &MaxPriorityQueue{
		array: make([]int, 32),
		size:  0,
	}
}
func (pq *MaxPriorityQueue) Len() int {
	return len(pq.array)
}

/**
 * 入队
 * @param key 入队元素
 */
func (h *MaxPriorityQueue) enQueue(key int) {
	//队列长度超出范围，扩容
	if h.size > len(h.array) {
		h.resize()
	}
	h.array[h.size] = key
	h.size += 1
	h.downAdjust()
}

func (h *MaxPriorityQueue) deQueue() (int, error) {
	if h.size <= 0 {
		return 0, errors.New("the queue is empty !")
	}
	//获取堆顶元素
	head := h.array[0]
	h.size -= 1
	h.array[0] = h.array[h.size]

	h.upAdjust()

	return head, nil
}

/**
 * 上浮调整
 */
func (h *MaxPriorityQueue) upAdjust() {
	childIndex := h.size - 1
	parentIndex := (childIndex - 1) / 2

	// temp保存插入的叶子节点值，用于最后的赋值
	tmp := h.array[childIndex]
	for childIndex > 0 && tmp > h.array[parentIndex] {
		//无需真正交换，单向赋值即可
		h.array[childIndex] = h.array[parentIndex]
		childIndex = parentIndex
		parentIndex = (parentIndex - 1) / 2
	}
	h.array[childIndex] = tmp
}
func (h *MaxPriorityQueue) downAdjust() {
	// temp保存父节点值，用于最后的赋值
	parentIndex := 0
	tmp := h.array[parentIndex]
	childIndex := 1
	for childIndex < h.size {
		// 如果有右孩子，且右孩子大于左孩子的值，则定位到右孩子
		if childIndex+1 < h.size && h.array[childIndex+1] > h.array[childIndex] {
			childIndex += 1
		}

		// 如果父节点大于任何一个孩子的值，直接跳出
		if tmp >= h.array[childIndex] {
			break
		}

		//无需真正交换，单向赋值即可
		h.array[parentIndex] = h.array[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	h.array[parentIndex] = tmp
}

func (h *MaxPriorityQueue) resize() {
	//队列容量翻倍
	newSize := h.size * 2
	newArray := make([]int, newSize)
	h.array = append(newArray, h.array...)
}
