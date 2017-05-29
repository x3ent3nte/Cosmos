package mysort

type Sortable interface {
	Len() int
	Less(x, y int) bool
	Swap(x, y int)
	Reslice(start, end int) Sortable
}

func Insertion(elems Sortable) {
	for i := 1; i < elems.Len(); i++ {
		current := i 
		for current >= 1 && elems.Less(current, current - 1) {
			elems.Swap(current - 1, current)
			current--
		}
	}
}

func QuickSort(elems Sortable) {
	if elems.Len() <= 1 {
		return
	}
	pivot := 0
	left := 0
	right := elems.Len() - 1

	for left <= right {
		for elems.Less(left, pivot) {
			left++
		}
		for elems.Less(pivot, right) {
			right--
		}
		if left <= right {
			elems.Swap(left, right)
			left++
			right--
		}
	}

	if right > 0 {
		QuickSort(elems.Reslice(0, right + 1))
	}
	if left < elems.Len() - 1 {
		QuickSort(elems.Reslice(left, elems.Len()))
	}
}



