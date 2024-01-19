package tree

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"testing"
	"time"
)

func (t *RbTree) IsBlackNodeBalanced() bool {
	if t.root.color != 1 {
		return false
	}
	count := t.root.blackNodes(t.null)
	if count == -1 {
		return false
	}
	return true
}

func (node *rbNode) blackNodes(nilNode *rbNode) int {
	if node == nilNode {
		return 1
	}
	left := node.left.blackNodes(nilNode)
	if left == -1 {
		return -1
	}
	right := node.right.blackNodes(nilNode)
	if right == -1 {
		return -1
	}
	if left != right {
		return -1
	}
	if node.color == 1 {
		left++
	}
	return left
}

func (t *RbTree) IsValidBST() bool {
	visit := make([]*rbNode, 0)
	var pre = math.MinInt
	node := t.root.left
	for len(visit) > 0 {
		// visit node
		if node != t.null {
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

func (t *RbTree) check(size int) {
	if !t.IsBlackNodeBalanced() {
		fmt.Println(t.String())
		panic("tree isn't balanced!")
	}
	if !t.IsValidBST() {
		fmt.Println(t.String())
		panic("tree isn't a bst!")
	}
	if size != -1 && size != t.count {
		fmt.Println(t.String())
		panic(fmt.Sprintf("tree's size is not right, expect %d, got %d", size, t.count))
	}
}

func TestRandomInsertAndDeleteRbTree(t *testing.T) {
	tree := NewRbTree()
	var arr []int
	dict := map[int]bool{}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10002; i++ {
		lrand := rand.Intn(30)
		if lrand > 25 {
			if len(arr) == 0 {
				continue
			}
			idx := rand.Intn(len(arr))
			v := arr[idx]
			arr = slices.Delete(arr, idx, idx+1)
			delete(dict, v)
			tree.Delete(v)
			tree.check(len(arr))
		} else {
			v := rand.Intn(100000)
			if !dict[v] {
				arr = append(arr, v)
				dict[v] = true
			}
			tree.Insert(v)
			tree.check(len(arr))
		}
	}

	// random Delete
	for i := 0; i < 10000; i++ {
		v := rand.Intn(1000000)
		tree.Delete(v)
		tree.check(-1)
	}
}
