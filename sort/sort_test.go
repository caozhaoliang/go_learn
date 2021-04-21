package sort

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestInsertSort(t *testing.T) {
	data := []int{1, 0, 35, 9, 7, 8, 5, 4, 3, 33}
	fmt.Println("sort before", data)
	InsertSort(data)
	fmt.Println("sort after", data)
}

func TestBubbleSort(t *testing.T) {
	//BubbleSort
	data := []int{1, 0, 35, 9, 7, 8, 5, 4, 3, 33}
	t.Log(data)
	BubbleSort(data)
	t.Log(data)
}
func TestBubbleSort2(t *testing.T) {
	data := []int{1, 0, 35, 9, 7, 8, 5, 4, 3, 33}
	t.Log(data)
	quickSort(data,0,9)
	t.Log(data)
}

func TestMergeSort(t *testing.T) {
	data := []int{1, 0, 35, 9, 7, 8, 5, 4, 3, 33}
	t.Log(data)
	data = mergeSort(data)
	t.Log(data)
}
func TestSize(t *testing.T){
	a := "曹"
	b :='a'
	c := 1
	fmt.Println(unsafe.Sizeof(a),unsafe.Sizeof(&a),unsafe.Sizeof(b),unsafe.Sizeof(&b))
	fmt.Println(unsafe.Sizeof(c))
}