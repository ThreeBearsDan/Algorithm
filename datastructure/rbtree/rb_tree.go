/**
 * Created by GoLand.
 * User: cruise
 * Date: 2021/6/5
 * Time: 8:20 下午
 */
package rbtree

import "strconv"

const (
	Red   = true
	Black = false
)

// the root of the tree
var Root *TreeNode

type TreeNode struct {
	key         string
	val         string
	left, right *TreeNode
	color       bool
}

func (node *TreeNode) isRed() bool {
	if node == nil {
		return false
	}

	return node.color
}

// key vs node.key
// if 1 then insert into right side
// if -1 then insert into left side
// if 0 then replace of its value
func (node *TreeNode) compare(key string) int {
	iNodeKey, _ := strconv.Atoi(node.key)
	iKey, _ := strconv.Atoi(key)

	if iKey > iNodeKey {
		return 1
	}

	if iKey < iNodeKey {
		return -1
	}

	return 0
}

func rotateLeft(n *TreeNode) *TreeNode {
	if !n.right.isRed() {
		return n
	}

	x := n.right
	n.right = x.left
	x.left = n
	x.color = n.color
	n.color = Red
	return x
}

func rotateRight(n *TreeNode) *TreeNode {
	if !n.left.isRed() {
		return n
	}

	x := n.left
	n.left = x.right
	x.right = n
	x.color = n.color
	n.color = Red
	return x
}

func flipColors(n *TreeNode) {
	if !n.right.isRed() {
		return
	}

	if !n.left.isRed() {
		return
	}

	n.left.color = Black
	n.right.color = Black
	n.color = Red
	Root.color = Black
}

func search(key string) (string, bool) {
	x := Root
	for x != nil {
		if x.compare(key) > 0 {
			x = x.right
		}

		if x.compare(key) < 0 {
			x = x.left
		}

		return x.val, true
	}

	return "null", false
}

func Put(h *TreeNode, key string, val string) *TreeNode {
	if h == nil {
		if Root == nil {
			return &TreeNode{
				key:   key,
				val:   val,
				color: Black,
			}
		}
		return &TreeNode{
			key:   key,
			val:   val,
			color: Red,
		}
	}

	if h.compare(key) > 0 {
		h.right = Put(h.right, key, val)
	}

	if h.compare(key) < 0 {
		h.left = Put(h.left, key, val)
	}

	if h.compare(key) == 0 {
		h.val = val
	}

	if h.right.isRed() && !h.left.isRed() {
		h = rotateLeft(h)
	}

	if h.left.isRed() && h.left.left.isRed() {
		h = rotateRight(h)
	}

	if h.right.isRed() && h.left.isRed() {
		flipColors(h)
	}

	return h
}
