package dig

const SENTINEL = 255

func binarySearch(arr []uint8, item uint8) uint8 {
	var low uint8 = 0
	high := uint8(len(arr) - 1)

	for low <= high {
		if mid := (high + low) / 2; arr[mid] < item {
			low = mid + 1
		} else if arr[mid] > item {
			high = mid - 1
		} else {
			return uint8(mid)
		}
	}
	return SENTINEL
}
