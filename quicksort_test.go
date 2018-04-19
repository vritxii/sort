package quicksort

import (
	"fmt"
	"math/rand"
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestSimpleQuickSortNonRecursive(t *testing.T) {
	N := 1024000
	arr := make([]int, N, N)
	r := 10*N
	for i:=0;i<N;i++{
		arr[i] = rand.Intn(r)
	}
	n := time.Now()
	SimpleSort(arr, false, false)
	fmt.Print("SimpleQuickSortNonRecursive: ")
	fmt.Println(time.Now().Sub(n))
	assert.Equal(t, true,IsSorted(arr))

}

func TestSimpleQuickSortRecursive(t *testing.T) {
	N := 1024000
	arr := make([]int, N, N)
	r := 10*N
	for i:=0;i<N;i++{
		arr[i] = rand.Intn(r)
	}
	n := time.Now()
	SimpleSort(arr, true, false)
	fmt.Print("SimpleQuickSortRecursive: ")
	fmt.Println(time.Now().Sub(n))
	assert.Equal(t, true,IsSorted(arr))
}

func TestSimpleQuickSortParallelRecursive(t *testing.T) {
	N := 1024000
	arr := make([]int, N, N)
	r := 10*N
	for i:=0;i<N;i++{
		arr[i] = rand.Intn(r)
	}
	n := time.Now()
	SimpleSort(arr, true, true)
	fmt.Print("SimpleQuickSortParallelRecursive: ")
	fmt.Println(time.Now().Sub(n))
	assert.Equal(t, true,IsSorted(arr))
}

func TestTriWayQuickSort(t *testing.T) {
	Debug = false
	N := 1024000
	arr := make([]int, N, N)
	r := 10*N
	for i:=0;i<N;i++{
		arr[i] = rand.Intn(r)
	}
	//fmt.Println(arr)
	n:=time.Now()
	quickSort3Way(arr, 0, len(arr)-1)
	fmt.Print("3 Way-QuickSort: ")
	fmt.Println(time.Now().Sub(n))
	//fmt.Println(arr)
	assert.Equal(t, true, IsSorted(arr))
}

func TestDualPivotQuickSort(t *testing.T){
	N := 1024000
	//arr1 := make([]int, N, N)
	arr2 := make([]int, N, N)
	r := 10*N
	for i:=0;i<N;i++{
		x := rand.Intn(r)
		//arr1[i] = x
		arr2[i] = x
	}
	//fmt.Println(arr2)
	n := time.Now()
	dualPivotQuickSort(arr2, 0, len(arr2)-1, 3)
	fmt.Print("Dual Pivot QuickSort: ")
	fmt.Println(time.Now().Sub(n))
	//fmt.Println(arr2)
	assert.Equal(t, true, IsSorted(arr2))
	//assert.Equal(t, true, IsSameElements(arr1,arr2))
}

func TestSortIntSlice(t *testing.T) {
	g = true
	N := 1024000
	arr:= make([]int, N, N)
	r := 10*N
	for i:=0;i<N;i++{
		arr[i] = rand.Intn(r)
	}
	a := IntSlice(arr[0:])
	//fmt.Println(a)
	n := time.Now()
	Sort(a)
	fmt.Print("General 4-Way QuickSort: ")
	fmt.Println(time.Now().Sub(n))
	//fmt.Println(a)
	assert.Equal(t, true, IsSorted(a))
}