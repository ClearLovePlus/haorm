package component

import "fmt"

//冒泡排序
func BubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			var temp int
			if array[i] > array[j] {
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	fmt.Print(array)
	return array
}
