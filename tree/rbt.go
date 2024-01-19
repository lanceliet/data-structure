package tree

import (
	"fmt"
)

const (
	RED   = 0
	BLACK = 1
)

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

type RbTree struct {
	null  *rbNode
	root  *rbNode
	count int
}

type rbNode struct {
	parent *rbNode
	left   *rbNode
	right  *rbNode
	color  int
	v      int
}

func NewRbTree() *RbTree {
	null := &rbNode{nil, nil, nil, BLACK, 0}
	return &RbTree{
		null:  null,
		root:  null,
		count: 0,
	}
}

func (t *RbTree) Find(cmp func(v int) int) *rbNode {
	node := t.root
	for node != t.null {
		switch cmp(node.v) {
		case -1:
			node = node.left
		case 1:
			node = node.right
		default:
			return node
		}
	}
	return t.null
}

func (t *RbTree) Insert(num int) {
	var p = t.null
	node := t.root
	for node != t.null {
		p = node
		if num < node.v {
			node = node.left
		} else if num > node.v {
			node = node.right
		} else {
			return
		}
	}
	i := &rbNode{
		parent: p,
		left:   t.null,
		right:  t.null,
		color:  RED,
		v:      num,
	}
	if p == t.null {
		t.root = i
	} else if num < p.v {
		p.left = i
	} else {
		p.right = i
	}
	t.insertFixup(i)
	t.count++
}

func (t *RbTree) insertFixup(i *rbNode) {
	for i.parent.color == RED {
		g := i.parent.parent
		if i.parent == g.left {
			u := g.right
			if u.color == RED {
				g.color = RED
				i.parent.color, u.color = BLACK, BLACK
				i = g
			} else {
				if i == i.parent.right {
					i = i.parent
					t.leftRotate(i)
				}
				i.parent.color = BLACK
				g.color = RED
				t.rightRotate(g)
			}
		} else {
			u := g.left
			if u.color == RED {
				g.color = RED
				u.color, i.parent.color = BLACK, BLACK
				i = g
			} else {
				if i == i.parent.left {
					i = i.parent
					t.rightRotate(i)
				}
				i.parent.color = BLACK
				g.color = RED
				t.leftRotate(g)
			}
		}
	}
	t.root.color = BLACK
}

func (t *RbTree) Delete(num int) {
	node := t.Find(func(v int) int {
		if num < v {
			return -1
		}
		if num > v {
			return 1
		}
		return 0
	})
	if node == t.null {
		return
	}
	var d *rbNode
	if node.left == t.null || node.right == t.null {
		d = node
	} else {
		d = t.leftmostNode(node.right)
	}
	if node != d {
		node.v = d.v
	}

	var s *rbNode
	if d.left != t.null {
		s = d.left
	} else {
		s = d.right
	}
	s.parent = d.parent
	if s.parent == t.null {
		t.root = s
	} else if s.parent.left == d {
		s.parent.left = s
	} else {
		s.parent.right = s
	}
	if d.color == BLACK {
		t.deleteFixup(s)
	}
	t.count--
}

func (t *RbTree) deleteFixup(d *rbNode) {
	for d.color == BLACK && d != t.root {
		if d == d.parent.left {
			b := d.parent.right
			if b.color == RED {
				b.color = BLACK
				d.parent.color = RED
				t.leftRotate(d.parent)
				b = d.parent.right
			}
			if b.left.color == BLACK && b.right.color == BLACK {
				b.color = RED
				d = d.parent
			} else {
				if b.right.color == BLACK {
					b.left.color = BLACK
					b.color = RED
					t.rightRotate(b)
					b = d.parent.right
				}
				b.right.color = BLACK
				b.color = d.parent.color
				d.parent.color = BLACK
				t.leftRotate(d.parent)
				d = t.root
			}
		} else {
			b := d.parent.left
			if b.color == RED {
				b.color = BLACK
				d.parent.color = RED
				t.rightRotate(d.parent)
				b = d.parent.left
			}
			if b.left.color == BLACK && b.right.color == BLACK {
				b.color = RED
				d = d.parent
			} else {
				if b.left.color == BLACK {
					b.right.color = BLACK
					b.color = RED
					t.leftRotate(b)
					b = d.parent.left
				}
				b.left.color = BLACK
				b.color = d.parent.color
				d.parent.color = BLACK
				t.rightRotate(d.parent)
				d = t.root
			}
		}
	}
	d.color = BLACK
}

func (t *RbTree) leftRotate(node *rbNode) {
	pivot := node.right
	pivot.parent = node.parent
	if pivot.parent == t.null {
		t.root = pivot
	} else if pivot.parent.left == node {
		pivot.parent.left = pivot
	} else {
		pivot.parent.right = pivot
	}
	node.right = pivot.left
	if node.right != t.null {
		node.right.parent = node
	}
	node.parent = pivot
	pivot.left = node
}

func (t *RbTree) rightRotate(node *rbNode) {
	pivot := node.left
	pivot.parent = node.parent
	if pivot.parent == t.null {
		t.root = pivot
	} else if pivot.parent.left == node {
		pivot.parent.left = pivot
	} else {
		pivot.parent.right = pivot
	}
	node.left = pivot.right
	if node.left != t.null {
		node.left.parent = node
	}
	node.parent = pivot
	pivot.right = node
}

func (t *RbTree) leftmostNode(node *rbNode) *rbNode {
	for node.left != t.null {
		node = node.left
	}
	return node
}

// String returns a string representation of container
func (t *RbTree) String() string {
	str := "RedBlackTree\n"
	if t.root != t.null {
		output(t.root, "", true, &str)
	}
	return str
}

func (node *rbNode) String() string {
	if node.color == BLACK {
		return fmt.Sprintf("\x1b[0;%dm%d\x1b[0m", textBlue, node.v)
	}
	return fmt.Sprintf("\x1b[0;%dm%d\x1b[0m", textRed, node.v)
}

func output(node *rbNode, prefix string, isTail bool, str *string) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.left, newPrefix, true, str)
	}
}
