package tree

type AvlTree struct {
	left  *AvlTree
	right *AvlTree
	v     int
	h     int
}

func (t *AvlTree) Insert(num int) *AvlTree {
	if t == nil {
		return &AvlTree{
			v: num,
			h: 1,
		}
	}
	if num == t.v {
		return t
	}
	if num < t.v {
		t.left = t.left.Insert(num)
	} else {
		t.right = t.right.Insert(num)
	}
	return t.doBalance()
}

func (t *AvlTree) Delete(num int) *AvlTree {
	if t == nil {
		return nil
	}
	if num == t.v {
		if t.left == nil {
			return t.right
		}
		if t.right == nil {
			return t.left
		}
		d := t.right.leftmostNode()
		t.right = t.right.Delete(d.v)
		t.v = d.v
	} else if num < t.v {
		t.left = t.left.Delete(num)
	} else {
		t.right = t.right.Delete(num)
	}
	return t.doBalance()
}

func (t *AvlTree) leftmostNode() *AvlTree {
	node := t
	for node.left != nil {
		node = node.left
	}
	return node
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
