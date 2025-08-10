package dig

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []uint8{1, 2, 4, 5, 6, 8, 12, 15, 16, 17, 20, 22, 34}
	var item uint8 = 5

	var expected uint8 = 3
	received := binarySearch(arr, item)
	if received != expected {
		t.Errorf("expected %d, received %d", expected, received)
	}
}

func TestBinarySearchOnNonExistentItem(t *testing.T) {
	arr := []uint8{1, 2, 4, 5, 6, 8, 12, 15, 16, 17, 20, 22, 34}
	var item uint8 = 3

	received := binarySearch(arr, item)
	if received != SENTINEL {
		t.Errorf("expected %d, received %d", SENTINEL, received)
	}
}
