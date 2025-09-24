package core

import "testing"

func TestIsSpaceLeft(t *testing.T) {
	inv := &Inventory{
		capacity:  10,
		usedspace: 5,
	}

	// There is space for 5 more items.
	// The function should return true, but it returns false.
	if !inv.IsSpaceLeft(5) {
		t.Errorf("IsSpaceLeft failed: expected true, got false")
	}

	// There is no space for 6 more items.
	// The function should return false, but it returns true.
	if inv.IsSpaceLeft(6) {
		t.Errorf("IsSpaceLeft failed: expected false, got true")
	}
}