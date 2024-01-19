package tree

type AvlTree struct {
	left  *AvlTree
	right *AvlTree
	v     int
	h     int
}

func (t *AvlTree) insert(v int) *AvlTree {
	if t == nil {
		return &AvlTree{
			v: v,
			h: 1,
		}
	}
	if v < t.v {
		t.left = t.left.insert(v)
	} else if v > t.v {
		t.right = t.right.insert(v)
	} else {
		return t
	}
	return t.doBalance()
}

func (t *AvlTree) delete(v int) *AvlTree {
	if t == nil {
		return nil
	}
	if t.v == v {
		if t.left == nil {
			return t.right
		}
		if t.right == nil {
			return t.left
		}
		s := t.right
		for s.left != nil {
			s = s.left
		}
		t.right = t.right.delete(s.v)
		t.v = s.v
		return t.doBalance()
	}
	if v < t.v {
		t.left = t.left.delete(v)
	} else {
		t.right = t.right.delete(v)
	}
	return t.doBalance()
}

func (t *AvlTree) doBalance() *AvlTree {
	switch t.balance() {
	case -2:
		if t.left.balance() > 0 {
			t.left = t.left.leftRotate()
		}
		return t.rightRotate()
	case 2:
		if t.right.balance() < 0 {
			t.right = t.right.rightRotate()
		}
		return t.leftRotate()
	default:
		t.updateHeight()
		return t
	}
}

func (t *AvlTree) balance() int {
	factor := 0
	if t.left != nil {
		factor -= t.left.h
	}
	if t.right != nil {
		factor += t.right.h
	}
	return factor
}

func (t *AvlTree) leftRotate() *AvlTree {
	pivot := t.right
	t.right = pivot.left
	pivot.left = t
	t.updateHeight()
	pivot.updateHeight()
	return pivot
}

func (t *AvlTree) rightRotate() *AvlTree {
	pivot := t.left
	t.left = pivot.right
	pivot.right = t
	t.updateHeight()
	pivot.updateHeight()
	return pivot
}

func (t *AvlTree) updateHeight() {
	t.h = 0
	if t.left != nil {
		t.h = t.left.h
	}
	if t.right != nil && t.right.h > t.h {
		t.h = t.right.h
	}
	t.h++
}
