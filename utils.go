package quicksort

import (
	"fmt"
	"reflect"
)
var Debug bool
func swap(arr []int, a,b int){
	if Debug{
		fmt.Printf("a,b:%d,%d\n", a,b)
	}
	t := arr[a]
	arr[a] = arr[b]
	arr[b] = t
}

func partition(arr []int, left, right int)int{
	x := arr[right]
	tail := left - 1
	for i:=left;i<right;i++{
		if arr[i] < x{
			tail ++
			swap(arr, tail, i)
		}
	}
	tail ++
	swap(arr, tail, right)
	return tail
}

func triPartition(arr []int, left, right int)(head, tail int){
	if arr[left] > arr[right]{
		swap(arr, left, right)
	}
	head = left
	tail = right
	x := arr[left]
	y := arr[right]
	for i:=left+1;i<tail;i++{
		if arr[i] < x{
			head ++
			swap(arr, head, i)
		}
		if arr[i] > y{
			tail --
			swap(arr, i, tail)
			i--
		}
	}

	if Debug{
		fmt.Println(arr)
		fmt.Printf("p,q:%d,%d\n", head,tail)
	}
	swap(arr, left, head)
	swap(arr,tail,right)

	return
}

func IsSorted(arr []int)(b bool){
	b = true
	N := len(arr)
	for i:=0;i<N-1;i++{
		if arr[i] > arr[i+1]{
			b = false
			return
		}
	}
	return
}

// SliceIsSorted tests whether a slice is sorted.
//
// The function panics if the provided interface is not a slice.
func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool {
	rv := reflect.ValueOf(slice)
	n := rv.Len()
	for i := n - 1; i > 0; i-- {
		if less(i, i-1) {
			return false
		}
	}
	return true
}


func IsSameElements(arr1, arr2 []int)bool{
	quickSort3Way(arr1,0, len(arr1)-1)
	N:=len(arr1)
	if len(arr1) != len(arr2){
		return false
	}
	for i:=0; i<N;i++{
		if arr1[i] != arr2[i]{
			fmt.Println(arr1[i],arr2[i])
			return false
		}
	}
	return true
}