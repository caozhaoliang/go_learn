package sort

import (
	"fmt"
)

// 对data进行正向排序
// 1、从第一个元素开始，该元素可以认为已经被排序；
// 2、取出下一个元素，在已经排序的元素序列中从后向前扫描；
// 3、如果该元素（已排序）大于新元素，将该元素移到下一位置；
// 4、重复步骤 3，直到找到已排序的元素小于或者等于新元素的位置；
// 5、将新元素插入到该位置后；
// 6、重复步骤 2~5。
// 难度系数 O(n^2)
//
func InsertSort(data []int) {
	for i := 1; i < len(data); i++ {
		tmp := data[i]
		// 从后往前扫描
		for j := i; j > 0 && tmp < data[j-1]; j-- {
			// 后和前交换值
			data[j], data[j-1] = data[j-1], tmp
		}
	}
}

func BubbleSort(data []int) {
	// 从头到尾 每个元素与之后的元素比较
	hasSort := false
	for i:= 0; i < len(data)-1;i++ {
		hasSort =false
		for j := 0; j<len(data)-i-1;j++ {
			if data[j]>data[j+1] {
				data[j],data[j+1]=data[j+1],data[j]
				hasSort=true
			}
		}
		fmt.Println(data)
		if !hasSort {
			break
		}
	}
}

func quickSort(arr []int,first,last int) int {
	left := first
	right := last
	pivot := arr[first]
	for left != right {
		for first < last && pivot < arr[right] {
			right--
		}
		for first<last && pivot >= arr[left] {
			left++
		}
		if left <right {
			arr[left],arr[right] = arr[right],arr[left]
		}
	}
	arr[first],arr[left] = arr[left],arr[first]
	return left
}
// 归并排序
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	p := len(nums) / 2
	left := mergeSort(nums[:p])
	right := mergeSort(nums[p:])
	return merge(left,right)
}
func merge(left,right []int) []int {
	i,j := 0,0
	m,n := len(left),len(right)
	result := make([]int,0,n+m)
	for {
		if i>= m || j>= n {
			break
		}
		if left[i] <= right[j] {
			result = append(result,left[i])
			i++
		}else{
			result = append(result,right[j])
			j++
		}
	}
	if i != m {
		for ;i<m;i++ {
			result = append(result,left[i])
		}
	}
	if j != n {
		for ; j<n;j++ {
			result = append(result,right[j])
		}
	}
	return result
}