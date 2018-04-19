package quicksort

import "sync"
const (
	minStep = 10240
)

func SimpleSort(arr []int, ops ...bool){
	b := true
	parallel := false
	if len(ops) > 0{
		b = ops[0]
		if len(ops) > 1{
			parallel = ops[1]
		}
	}
	if b{
		if parallel{
			parallelRecursiveQuickSort(arr, 0, len(arr)-1)
		}else{
			recursiveQuickSort(arr, 0, len(arr)-1)
		}

	}else{
		nonRecursiveQuickSort(arr, 0, len(arr)-1)
	}
}



func recursiveQuickSort(arr []int, left, right int){
	if left < right{
		p := partition(arr, left, right)
		recursiveQuickSort(arr, left, p-1)
		recursiveQuickSort(arr, p+1, right)
	}
}

func parallelRecursiveQuickSort(arr []int, left, right int, lwg ...*sync.WaitGroup) {
	if left >= right {
		return
	}
	partPoint := partition(arr, left, right)
	if partPoint < left {
		return
	}
	wg := &sync.WaitGroup{}

	if partPoint-left >= minStep {
		wg.Add(1)
		go parallelRecursiveQuickSort(arr, left, partPoint-1, wg)
	}
	if right-partPoint >= minStep {
		wg.Add(1)
		go parallelRecursiveQuickSort(arr, partPoint+1, right, wg)
	}

	if partPoint-left < minStep{
		parallelRecursiveQuickSort(arr, left, partPoint-1)
	}
	if right-partPoint < minStep {
		parallelRecursiveQuickSort(arr, partPoint+1, right)
	}

	wg.Wait()
	if len(lwg) > 0 {
		lwg[0].Done()
	}
}

func nonRecursiveQuickSort(arr []int, left, right int){
	s := &Stack{}
	if left<right{
		p := partition(arr, left, right)
		if p-1 > left{
			s.Push(left)
			s.Push(p-1)
		}
		if p+1 < right{
			s.Push(p+1)
			s.Push(right)
		}
		for !s.IsEmpty(){
			r := s.Pop()
			l := s.Pop()
			p = partition(arr, l, r)
			if p-1>l{
				s.Push(l)
				s.Push(p-1)
			}
			if p+1 < r{
				s.Push(p+1)
				s.Push(r)
			}
		}
	}
}