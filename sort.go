/**
 * 十大排序算法
 */

package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 1, 4, 8, 9, 7, 3, 2, 6}
	fmt.Println("初始数组: ", nums)
	nums = []int{5, 1, 4, 8, 9, 7, 3, 2, 6}
	BubbleSort(nums)
	nums = []int{5, 1, 4, 8, 9, 7, 3, 2, 6}
	SelectionSort(nums)
	nums = []int{5, 1, 4, 8, 9, 7, 3, 2, 6}
	InsertionSort(nums)
	nums = []int{5, 1, 4, 8, 9, 7, 3, 2, 6}
	ShellSort(nums)
	nums = []int{5, 1, 4, 8, 9, 7, 3, 2, 6}
	MergeSort(nums)
	nums = []int{5, 1, 4, 8, 9, 7, 3, 2, 6}
	QuickSort(nums)
}

// 冒泡排序（冒泡）
func BubbleSort(nums []int) {
	size := len(nums)
	for i := size - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	fmt.Println("冒泡排序: ", nums)
}

// 选择排序（每次选最小的）
func SelectionSort(nums []int) {
	size := len(nums)
	for i := 0; i < size; i++ {
		minIndex := i
		for j := i; j < size; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		temp := nums[i]
		nums[i] = nums[minIndex]
		nums[minIndex] = temp
	}
	fmt.Println("选择排序: ", nums)
}

// 插入排序（往有序列中插入）
func InsertionSort(nums []int) {
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			} else {
				break
			}
		}
	}
	fmt.Println("插入排序: ", nums)
}

// 希尔排序（间隔序列的插入排序）
func ShellSort(nums []int) {
	size := len(nums)
	for gap := size / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < size; i++ {
			j, cur := i, nums[i]
			for j-gap >= 0 && cur < nums[j-gap] {
				nums[j] = nums[j-gap]
				j = j - gap
			}
			nums[j] = cur
		}
	}
	fmt.Println("希尔排序: ", nums)
}

// 归并排序（给左右子序列排序，再合并）
func MergeSort(nums []int) {
	size := len(nums)
	middle := size / 2
	left := nums[0:middle]
	right := nums[middle:size]
	nums = Merge(left, right)
	fmt.Println("归并排序: ", nums)
}
func Merge(left []int, right []int) []int {
	// 左有序
	lSize := len(left)
	if lSize > 1 {
		lMiddle := lSize / 2
		left = Merge(left[0:lMiddle], left[lMiddle:lSize])
	}

	// 右有序
	rSize := len(right)
	if rSize > 1 {
		rMiddle := rSize / 2
		right = Merge(right[0:rMiddle], right[rMiddle:rSize])
	}

	// 合并左右有序
	var nums []int
	i, j := 0, 0
	for len(nums) < lSize+rSize {
		if i < lSize && j < rSize {
			if left[i] < right[j] {
				nums = append(nums, left[i])
				i++
			} else {
				nums = append(nums, right[j])
				j++
			}
		}
		if i >= lSize && j < rSize {
			nums = append(nums, right[j])
			j++
		}
		if j >= rSize && i < lSize {
			nums = append(nums, left[i])
			i++
		}
	}
	//fmt.Println(left,right,nums)
	return nums
}

// 快速排序（每个数的左边都小于右边）
func QuickSort(nums []int) {
	size := len(nums)
	nums = Quick(nums, 0, size-1)
	fmt.Println("快速排序: ", nums)
}
func Quick(nums []int, left int, right int) []int {
	if left >= right {
		return nums
	}

	f, i, j := true, left, right
	for i < j {
		if f {
			if nums[i] > nums[left] {
				f = !f
				continue
			}
			i++
		} else {
			if nums[j] <= nums[left] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
				f = !f
				continue
			}
			j--
		}
	}
	index := i
	if nums[index] > nums[left] {
		index = i - 1
	}
	temp := nums[index]
	nums[index] = nums[left]
	nums[left] = temp

	nums = Quick(nums, left, index-1)
	nums = Quick(nums, index+1, right)

	return nums
}

// 堆排序
func HeapSort([]nums int) {

}

