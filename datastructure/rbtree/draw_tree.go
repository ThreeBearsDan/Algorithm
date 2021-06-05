/**
 * Created by GoLand.
 * User: cruise
 * Date: 2021/6/5
 * Time: 8:19 下午
 */
package rbtree

import (
	"fmt"
	"math"
)

var (
	lProfile  = make([]int, MaxHeight)
	rProfile  = make([]int, MaxHeight)
	printNext = 0
)

const (
	MaxHeight = 1000
	Infinity  = 1 << 20

	// adjust gap between left and right nodes
	Gap = 2
)

type AsciiNode struct {
	left, right *AsciiNode

	// length of the edge from this node to its children
	edgeLen, height, labLen int

	// -1 = left, 0 = root, 1 = right
	parentDir int

	// max supported uint32 in dec, 10 digits max
	label string

	// color
	color bool
}

func buildAsciiTreeRecursive(t *TreeNode) *AsciiNode {
	if t == nil {
		return nil
	}

	node := new(AsciiNode)
	node.left = buildAsciiTreeRecursive(t.left)
	node.right = buildAsciiTreeRecursive(t.right)

	if node.left != nil {
		node.left.parentDir = -1
	}

	if node.right != nil {
		node.right.parentDir = 1
	}

	node.label = t.key
	node.color = t.color
	node.labLen = len(node.label)
	return node
}

// copy the tree into the ascii node structure
func buildAsciiTree(t *TreeNode) *AsciiNode {
	if t == nil {
		return nil
	}

	node := buildAsciiTreeRecursive(t)
	node.parentDir = 0
	return node
}

// The following function fills in the lprofile array for the given tree.
// It assumes that the center of the label of the root of this tree
// is located at a position (x,y).  It assumes that the edge_length
// fields have been computed for this tree.
func computeRprofile(node *AsciiNode, x, y int) {
	if node == nil {
		return
	}

	var notLeft int
	if node.parentDir != -1 {
		notLeft = 1
	}

	if rProfile[y] < x+int(math.Floor(float64(node.labLen-notLeft)/2)) {
		rProfile[y] = x + int(math.Floor(float64(node.labLen-notLeft)/2))
	}

	if node.right != nil {
		i := 1
		for i <= node.edgeLen && y+i < MaxHeight {
			if rProfile[y+1] < x+i {
				rProfile[y+1] = x + i
			}
			i += 1
		}
	}

	computeRprofile(node.left, x-node.edgeLen-1, y+node.edgeLen+1)
	computeRprofile(node.right, x+node.edgeLen+1, y+node.edgeLen+1)
}

func computeLprofile(node *AsciiNode, x, y int) {
	if node == nil {
		return
	}

	var isleft int
	if node.parentDir == -1 {
		isleft = 1
	}

	if lProfile[y] > x-int(math.Floor(float64(node.labLen-isleft)/2)) {
		lProfile[y] = x - int(math.Floor(float64(node.labLen-isleft)/2))
	}

	if node.left != nil {
		i := 1
		for i <= node.edgeLen && y+i < MaxHeight {
			if lProfile[y+i] > x-i {
				lProfile[y+i] = x - i
			}
			i += 1
		}
	}

	computeLprofile(node.left, x-node.edgeLen-1, y+node.edgeLen+1)
	computeLprofile(node.right, x+node.edgeLen+1, y+node.edgeLen+1)
}

// this function fills in the edge_length
// and height fields of the specified tree
func computeEdgeLengths(node *AsciiNode) {
	if node == nil {
		return
	}

	computeEdgeLengths(node.left)
	computeEdgeLengths(node.right)

	// first fill in the edge_length of node
	var hMin int
	if node.right == nil && node.left == nil {
		node.edgeLen = 0
	} else {
		if node.left != nil {
			i := 0
			for i < node.left.height && i < MaxHeight {
				rProfile[i] = -Infinity
				i++
			}
			computeRprofile(node.left, 0, 0)
			hMin = node.left.height
		} else {
			hMin = 0
		}

		if node.right != nil {
			i := 0
			for i < node.right.height && i < MaxHeight {
				lProfile[i] = Infinity
				i++
			}
			computeLprofile(node.right, 0, 0)
			hMin = int(math.Min(float64(node.right.height), float64(hMin)))
		} else {
			hMin = 0
		}

		delta := 4
		i := 0
		for i < hMin {
			if delta < Gap+1+rProfile[i]-lProfile[i] {
				delta = Gap + 1 + rProfile[i] - lProfile[i]
			}
			i += 1
		}

		// If the node has two children of height 1, then we allow the
		// two leaves to be within 1, instead of 2
		if ((node.left != nil && node.left.height == 1) ||
			(node.right != nil && node.right.height == 1)) &&
			delta > 4 {
			delta -= 1
		}

		node.edgeLen = int(math.Floor(float64((delta+1)/2 - 1)))
	}

	// now fill in the height of node
	h := 1
	if node.left != nil {
		if h < node.left.height+node.edgeLen+1 {
			h = node.left.height + node.edgeLen + 1
		}
	}

	if node.right != nil {
		if h < node.right.height+node.edgeLen+1 {
			h = node.right.height + node.edgeLen + 1
		}
	}
	node.height = h
}

// This function prints the given level of the given tree, assuming
// that the node has the given x coordinate.
func printLevel(node *AsciiNode, x, level int) {
	if node == nil {
		return
	}

	var isLeft int
	if node.parentDir == -1 {
		isLeft = 1
	}

	if level == 0 {
		spaces := x - printNext - int(math.Floor(float64(node.labLen-isLeft)/2))
		drawSpaces(spaces)

		printNext += spaces
		printWithColor(node.color, node.label)
		printNext += node.labLen
	} else if node.edgeLen >= level {
		if node.left != nil {
			spaces := x - printNext - level
			drawSpaces(spaces)
			printNext += spaces
			fmt.Print("/")
			printNext += 1
		}

		if node.right != nil {
			spaces := x - printNext + level
			drawSpaces(spaces)
			printNext += spaces
			fmt.Print("\\")
			printNext += 1
		}
	} else {
		printLevel(node.left, x-node.edgeLen-1, level-node.edgeLen-1)
		printLevel(node.right, x+node.edgeLen+1, level-node.edgeLen-1)
	}
}

func drawSpaces(count int) {
	for i := 0; i < count; i++ {
		fmt.Print(" ")
	}
}

func deserialize(s string) *TreeNode {
	if s == "{}" {
		return nil
	}

	return nil
}

// prints ascii tree for given Tree structure
func DrawTree(t *TreeNode) {
	if t == nil {
		return
	}

	proot := buildAsciiTree(t)
	computeEdgeLengths(proot)

	i := 0
	for i < proot.height && i < MaxHeight {
		lProfile[i] = Infinity
		i += 1
	}

	computeLprofile(proot, 0, 0)

	xMin := 0
	i = 0
	for i < proot.height && i < MaxHeight {
		if xMin > lProfile[i] {
			xMin = lProfile[i]
		}
		i += 1
	}

	i = 0
	for i < proot.height {
		printNext = 0
		printLevel(proot, -xMin, i)
		fmt.Println()
		i += 1
	}

	if proot.height >= MaxHeight {
		fmt.Printf("This tree is taller than %d, and may be drawn incorrectly.\n", MaxHeight)
	}
}

func printWithColor(color bool, s string) {
	if color {
		fmt.Printf("\033[1;31;m%s\033[0m", s)
		return
	}

	fmt.Printf(s)
}

