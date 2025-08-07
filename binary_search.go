package dig

const SENTINEL = 255

func binarySearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1

	var mid int
	for low <= high {
		mid = (high + mid) / 2
		if arr[mid] < item {
			low = mid + 1
		} else if arr[mid] > item {
			high = mid - 1
		} else {
			return mid
		}
	}
	return SENTINEL
}
