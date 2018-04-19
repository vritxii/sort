package quicksort

func dualPivotQuickSort(arr []int, left, right, div int){
	l := right - left
	if l < 27{
		for i:=left + 1;i<=right;i++{
			for j:=i;j>left && arr[j] < arr[j-1]; j--{
				swap(arr, j, j-1)
			}
		}
		return
	}
	third := int(l/div)
	m1 := left + third
	m2 := right - third
	if m1<=left{
		m1 = left + 1
	}
	if m2>=right{
		m2 = right - 1
	}
	if arr[m1] < arr[m2]{
		swap(arr,m1, left)
		swap(arr, m2, right)
	}else {
		swap(arr, m1, right)
		swap(arr, m2, left)
	}

	pivot1 := arr[left]
	pivot2 := arr[right]
	less := left + 1
	great := right - 1
	for k:=less; k<=great;k++{
		if arr[k] < pivot1 {
			swap(arr, k, less)
			less ++
		}else if arr[k] > pivot2 {
			for k < great && arr[great] > pivot2 {
				great--
			}
			swap(arr, k, great)
			great--
			if arr[k] < pivot1 {
				swap(arr, k, less)
				less ++
			}
		}
	}
	dist := great - less
	if dist < 13{
		div ++
	}
	swap(arr, less - 1, left);
	swap(arr, great + 1, right);
	// subarrays
	dualPivotQuickSort(arr, left, less - 2, div);
	dualPivotQuickSort(arr, great + 2, right, div);
	// equal elements
	if dist > l - 13 && pivot1 != pivot2 {
		for k := less; k <= great; k++ {
		if arr[k] == pivot1 {
			swap(arr, k, less);
			less ++
		} else if arr[k] == pivot2 {
			swap(arr, k, great);
			great --
			if arr[k] == pivot1 {
				swap(arr, k, less);
				less++
			}
		}
		}
	}
	// subarray
	if pivot1 < pivot2 {
		dualPivotQuickSort(arr, less, great, div);
	}
}