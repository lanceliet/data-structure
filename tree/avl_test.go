package tree

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"testing"
	"time"
)

func TestRandomInsertAndDeleteAvl(t *testing.T) {
	var avl *AvlTree
	var arr []int
	dict := make(map[int]bool)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10000; i++ {
		lRand := rand.Intn(30)
		if lRand > 25 {
			if len(arr) == 0 {
				continue
			}
			idx := rand.Intn(len(arr))
			v := arr[idx]
			arr = slices.Delete(arr, idx, idx+1)
			delete(dict, v)
			avl = avl.Delete(v)
			avl.check(len(arr))

		} else {
			v := rand.Intn(100000)
			if !dict[v] {
				arr = append(arr, v)
				dict[v] = true
			}
			avl = avl.Insert(v)
			avl.check(len(arr))

		}
	}

	// random Delete
	for i := 0; i < 10000; i++ {
		v := rand.Intn(100000)
		avl = avl.Delete(v)
		avl.check(-1)
	}
}

func (t *AvlTree) check(size int) {
	if !t.IsAvlBalanced() {
		panic("tree isn't balanced!")
	}
	if !t.IsValidBST() {
		panic("tree isn't a bst!")
	}
	if size != -1 {
		s := t.Size()
		if size != s {
			panic(fmt.Sprintf("tree's size is not right, expect %d, got %d", size, t.Size()))
		}
	}
}

func (t *AvlTree) Size() int {
	if t == nil {
		return 0
	}
	return t.left.Size() + t.right.Size() + 1
}

func (t *AvlTree) IsValidBST() bool {
	if t == nil {
		return true
	}
	visit := make([]*AvlTree, 0)
	pre := math.MinInt
	node := t.left
	for len(visit) > 0 || node != nil {

		// visit node
		if node != nil {
			visit = append(visit, node)
			node = node.left
			continue
		}
		node = visit[len(visit)-1]
		visit = visit[:len(visit)-1]
		if pre >= node.v {
			return false
		}
		pre = node.v
		node = node.right
	}
	return true
}

func (t *AvlTree) IsAvlBalanced() bool {
	var getDepth func(parent *AvlTree) int
	getDepth = func(parent *AvlTree) int {
		if parent == nil {
			return 0
		}
		left := getDepth(parent.left)
		if left == -1 {
			return left
		}
		right := getDepth(parent.right)
		if right == -1 {
			return right
		}
		if abs(left-right) > 1 {
			return -1
		}
		return max(left, right) + 1
	}
	if getDepth(t) == -1 {
		return false
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
