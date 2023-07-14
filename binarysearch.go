package binarysearch

import "errors"

// Search searches in the sorted array for the target
// and returns if found the index position
func Search(array []int, target int) (int, error) {

	startPos := 0
	endPos := len(array)

	for {
		if startPos == endPos {
			return -1, errors.New("No target found")
		}

		midPos := (startPos + endPos) / 2
		midValue := array[midPos]

		if target == midValue {
			return midPos, nil
		} else if target < midValue {
			endPos = midPos
			continue
		} else {
			startPos = midPos + 1
			continue
		}

	}
}
