package common

import (
	"cmp"
	"fmt"
)

// BinarySearch returns the index of the first occurence of target in the search space.
// It is expected that the search space is an ordered array.
func BinarySearch[T cmp.Ordered](target T, searchSpace []T) (int, error) {
	index, err := recursiveBS(target, searchSpace, 0, len(searchSpace))
	if err != nil {
		return 0, err
	}

	// scan backwards through a sequence of equal values to ensure we return the first matching index
	for {
		if index == 0 || searchSpace[index-1] != target {
			break
		}
		index--
	}

	return index, nil
}

// recursiveBS returns the index of any occurence of target in the search space.
// It is expected that the search space is an ordered array.
// start is an inclusive index, end is an exclusive index.
// To search an entire space, this function should be called as recursiveBS(<some target>, space, 0, len(space)).
func recursiveBS[T cmp.Ordered](target T, searchSpace []T, start int, end int) (int, error) {
	if start > end {
		return 0, fmt.Errorf("start index %v greater than end index %v", start, end)
	}

	if start == end {
		return 0, fmt.Errorf("inclusive index start equals exclusive index end (both are %v)", start)
	}

	midpoint := (end-start)/2 + start
	elem := searchSpace[midpoint]

	comparison := cmp.Compare(elem, target)

	if start+1 == end && comparison != 0 {
		// target not present in search space, escape
		return 0, fmt.Errorf("target %v not present in search space", target)
	}

	if comparison == 0 {
		return midpoint, nil
	}

	if comparison > 0 { // elem > target
		return recursiveBS(target, searchSpace, start, midpoint)
	} else { // elem < target
		return recursiveBS(target, searchSpace, midpoint, end)
	}
}
