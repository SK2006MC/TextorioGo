package core

type Inventory struct {
	capacity  int
	usedspace int
	items     Items
}

func (i *Inventory) IsSpaceLeft(c int) bool {
	if i.usedspace+c <= i.capacity {
		return false
	}
	return true
}

func (i *Inventory) Has(in Item, c int) bool {
	if i.items[in.name] < c {
		return false
	}
	return true
}

func (i *Inventory) Get(in Item, c int) {
	if i.Has(in, c) {
		i.items[in.name] -= c
		i.usedspace -= c
	}
}

func (i *Inventory) Put(item Item, c int) {
	if i.IsSpaceLeft(c) {
		i.items[item.name] += c
		i.usedspace += c
	}
}
