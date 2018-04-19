package quicksort

import "sync"

const (
	insertionSortThreshold = 16
	ms = 10240
)

var g bool
func init(){
	g = true
}
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n-1)
}

func insertionSort(data Interface, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func quickSort(data Interface, lo int, hi int,lwg ...*sync.WaitGroup) {
	if hi-lo < insertionSortThreshold {
		if hi-lo > 0 {
			insertionSort(data, lo, hi)
		}
		return
	}

	midpoint := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
	// insertion sort lo, mid, hi elements
	if data.Less(midpoint, lo) {
		data.Swap(midpoint, lo)
	}
	if data.Less(hi, midpoint) {
		data.Swap(hi, midpoint)
		if data.Less(midpoint, lo) {
			data.Swap(midpoint, lo)
		}
	}
	// p, q, r (terms from the paper) are now sorted, put q at a[lo+1]
	data.Swap(lo+1, midpoint)

	// Pointers a and b initially point to the first element of the array while c
	// and d initially point to the last element of the array.
	a := lo + 2
	b := lo + 2
	c := hi - 1
	d := hi - 1

	for b <= c {
		for data.Less(b, lo+1) && b <= c {
			if data.Less(b, lo) {
				data.Swap(a, b)
				a++
			}
			b++
		}

		for data.Less(lo+1, c) && b <= c {
			if data.Less(hi, c) {
				data.Swap(c, d)
				d--
			}
			c--
		}

		if b <= c {
			if data.Less(hi, b) {
				if data.Less(c, lo) {
					data.Swap(b, a)
					data.Swap(a, c)
					a++
				} else {
					data.Swap(b, c)
				}
				data.Swap(c, d)
				b++
				c--
				d--
			} else {
				if data.Less(c, lo) {
					data.Swap(b, a)
					data.Swap(a, c)
					a++
				} else {
					data.Swap(b, c)
				}
				b++
				c--
			}
		}
	}

	a--
	b--
	c++
	d++
	data.Swap(lo+1, a)
	data.Swap(a, b)

	a--
	data.Swap(lo, a)
	data.Swap(hi, d)
	wg := &sync.WaitGroup{}
	if g && (a-lo >= ms || b-a >= ms || d-b>=ms || hi-d >=ms) {
		wg.Add(4)
		go quickSort(data, lo, a-1, wg)
		go quickSort(data, a+1, b-1, wg)
		go quickSort(data, b+1, d-1, wg)
		go quickSort(data, d+1, hi, wg)
	}else{
		quickSort(data, lo, a-1)
		quickSort(data, a+1, b-1)
		quickSort(data, b+1, d-1)
		quickSort(data, d+1, hi)
	}
	wg.Wait()
	if len(lwg) > 0{
		lwg[0].Done()
	}
}

// IsSorted reports whether data is sorted.
func AnyIsSorted(data Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

// Convenience types for common cases

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p IntSlice) Sort() { Sort(p) }

// Float64Slice attaches the methods of Interface to []float64, sorting in increasing order
// (not-a-number values are treated as less than other values).
type Float64Slice []float64

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// isNaN is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN(f float64) bool {
	return f != f
}

// Sort is a convenience method.
func (p Float64Slice) Sort() { Sort(p) }

// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p StringSlice) Sort() { Sort(p) }

// Convenience wrappers for common cases

// Ints sorts a slice of ints in increasing order.
func Ints(a []int) { Sort(IntSlice(a)) }

// Float64s sorts a slice of float64s in increasing order
// (not-a-number values are treated as less than other values).
func Float64s(a []float64) { Sort(Float64Slice(a)) }

// Strings sorts a slice of strings in increasing order.
func Strings(a []string) { Sort(StringSlice(a)) }

// IntsAreSorted tests whether a slice of ints is sorted in increasing order.
func IntsAreSorted(a []int) bool { return IsSorted(IntSlice(a)) }

// Float64sAreSorted tests whether a slice of float64s is sorted in increasing order
// (not-a-number values are treated as less than other values).
func Float64sAreSorted(a []float64) bool { return AnyIsSorted(Float64Slice(a)) }

// StringsAreSorted tests whether a slice of strings is sorted in increasing order.
func StringsAreSorted(a []string) bool { return AnyIsSorted(StringSlice(a)) }
