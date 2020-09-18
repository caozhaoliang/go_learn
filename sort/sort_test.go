package sort

import (
	"fmt"
	"testing"
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