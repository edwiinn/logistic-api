package sort

func Sort(arr ArraySortInterface) {
	quickSort(arr, 0, arr.Length()-1)
}

//Strings : Sort Array of String
func Strings(arr []string) {
	Sort(StringSort(arr))
}

//https://www.geeksforgeeks.org/quick-sort/
func quickSort(arr ArraySortInterface, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr ArraySortInterface, low, high int) int {
	pivot := high
	i := low - 1
	for j := low; j < high; j++ {
		if arr.Compare(j, pivot) {
			i++
			arr.Swap(i, j)
		}
	}
	arr.Swap(i+1, high)
	return i + 1
}
