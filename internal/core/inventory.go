package core

// Inventory represents an inventory of items, with a limited capacity.
type Inventory struct {
	capacity  int
	usedspace int
	items     Items
}


// IsSpaceLeft checks if there is enough space in the inventory for a given count of items.
// It returns true if there is space, and false otherwise.
// IsSpaceLeft is intended to check if there is enough space in the inventory for a given count of items.
// It should return true if there is space, and false otherwise.

func (i *Inventory) IsSpaceLeft(c int) bool {
	return i.usedspace+c <= i.capacity
}

// Has checks if the inventory contains at least a given count of a specific item.
func (i *Inventory) Has(in Item, c int) bool {
	if i.items[in.name] < c {
		return false
	}
	return true
}

// Get removes a given count of a specific item from the inventory.
// It does nothing if the inventory does not have enough of the item.
func (i *Inventory) Get(in Item, c int) {
	if i.Has(in, c) {
		i.items[in.name] -= c
		i.usedspace -= c
	}
}

// Put adds a given count of a specific item to the inventory.
// It does nothing if there is not enough space in the inventory.
func (i *Inventory) Put(item Item, c int) {
	if i.IsSpaceLeft(c) {
		i.items[item.name] += c
		i.usedspace += c
	}
}
