package sort

type ArraySortInterface interface {
	Length() int
	Compare(i, j int) bool
	Swap(i, j int)
}

type StringSort []string

//Length : Length of Array of String
func (si StringSort) Length() int {
	return len(si)
}

//Compare : Check if the first index in array of string should be swapped with second index
func (si StringSort) Compare(i, j int) bool {
	if si[i] <= si[j] {
		return true
	}
	return false
}

func (si StringSort) Swap(i, j int) {
	si[i], si[j] = si[j], si[i]
}
