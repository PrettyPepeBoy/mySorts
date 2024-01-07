package quicksort

func Sort(arr []int, low int, high int) {
	if low < high {
		index := Partition(arr, low, high)

		Sort(arr, low, index-1)
		Sort(arr, index+1, high)
	}
}

func Partition(arr []int, low int, high int) int {
	i := low - 1
	pivot := arr[high]
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[high], arr[i+1] = arr[i+1], arr[high]
	return i + 1
}

func MySort(arr []int, start int, end int) {
	if start < end {
		index := MyPartition(arr, start, end)

		MySort(arr, start, index-1)
		MySort(arr, index, end)
	}
}

func MyPartition(mas []int, start int, end int) int {
	elem := (mas[start] + mas[end]) / 2
	i := start
	j := end
	for i <= j {
		for mas[i] <= elem {
			i++
		}
		for mas[j] > elem {
			j--
		}
		if i >= j {
			break
		}

		mas[i], mas[j] = mas[j], mas[i]
	}
	return i
}
