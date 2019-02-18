package listops

// IntList is a list of ints
type IntList []int

// predFunc returns true or false whether or not an int satisfies the function
type predFunc func(int) bool

// unaryFunc transforms an int into another
type unaryFunc func(int) int

// binFunc combines to ints into another
type binFunc func(x, y int) int

// Length returns the number of items in a list
func (list *IntList) Length() int {
	length := 0
	for range *list {
		length++
	}
	return length
}

// Filter returns a copy of a list containing only items which
// satisy a predicate function
func (list *IntList) Filter(fn predFunc) IntList {
	filtered := make(IntList, 0)
	for _, value := range *list {
		if fn(value) {
			filtered = filtered.Append(IntList{value})
		}
	}
	return filtered
}

// Map returns a list where each value has been transformed by
// the supplied function
func (list *IntList) Map(fn unaryFunc) IntList {
	mapped := make(IntList, list.Length())
	for i, value := range *list {
		mapped[i] = fn(value)
	}
	return mapped
}

// Reverse returns a list where items are in reverse order
func (list *IntList) Reverse() IntList {
	reversed := make(IntList, list.Length())
	length := list.Length()
	for i, value := range *list {
		reversed[length-1-i] = value
	}
	return reversed
}

// Append returns a list with the original values plus the values supplied
// added at the end
func (list *IntList) Append(values IntList) IntList {
	length := list.Length()
	newList := make(IntList, length+values.Length())
	for i, value := range *list {
		newList[i] = value
	}
	for i, value := range values {
		newList[length+i] = value
	}

	return newList
}

// Concat appends each provided list to a lsit
func (list *IntList) Concat(lists []IntList) IntList {
	newList := make(IntList, 0)
	newList = newList.Append(*list)
	for _, l := range lists {
		newList = newList.Append(l)
	}
	return newList
}

// Foldl reduces a list into a single value by applying a function
// on an accumulator and each value from left to right
func (list *IntList) Foldl(fn binFunc, initial int) int {
	if list.Length() == 0 {
		return initial
	}

	l := *list
	head := l[0]
	tail := l[1:]
	return tail.Foldl(fn, fn(initial, head))
}

// Foldr reduces a list into a single value by applying a function
// on each value and an accumulator from right to left
func (list *IntList) Foldr(fn binFunc, initial int) int {
	if list.Length() == 0 {
		return initial
	}

	l := *list
	last := l[l.Length()-1]
	remainder := l[0 : l.Length()-1]
	return remainder.Foldr(fn, fn(last, initial))
}
