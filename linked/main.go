package linked

import "fmt"

func Main() {
	list1 := IntArrayToNodeList([]int{1, 2, 4})
	list2 := IntArrayToNodeList([]int{1, 3, 4})
	lists := MergeTwoLists(list1, list2)
	fmt.Println(lists)

}
