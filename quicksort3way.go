package quicksort


func quickSort3Way(arr []int, left, right int){
	if left < right{
		p,q := triPartition(arr, left, right)
		quickSort3Way(arr, left, p-1)
		quickSort3Way(arr, p+1, q-1)
		quickSort3Way(arr, q+1, right)
	}
}
