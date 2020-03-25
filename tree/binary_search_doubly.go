/**
二叉搜索树的某个节点的前驱和后继节点
思考方式：（拿前驱举例，后继也一样，只不过是前驱做法的垂直镜像）
给定节点
1. 如果有左子树，在其中找到最大的那个节点
2. 如果没有左子树，需要看给点节点是否是其父节点的右子树，
如果是则父节点就是前驱，不然就继续往上找，直到父子关系是右子树
*/
package main

import "fmt"

func main() {
	var t = &treeNode{
		data: 100,
	}
	t.insert(50)
	t.insert(30)
	t.insert(80)
	t.insert(60)
	t.insert(70)
	t.insert(90)

	layerPrint([]*treeNode{t}, 0)

	rv := t.find(60)
	fmt.Printf("found 60s: %v\n", rv)
	if len(rv) > 0 {
		fmt.Printf("1st of 60s prev: %v\n", rv[0].findPrevNode())
		fmt.Printf("1st of 60s next: %v\n", rv[0].findNextNode())
	}
}

func layerPrint(links []*treeNode, depth int) {
	var nexts []*treeNode
	var i int

	for {
		if i == 0 {
			fmt.Printf("[Depth:%d] ", depth)
		}

		currentNode := links[i]

		fmt.Printf("[%p->%s:%p][%v] ",
			currentNode.parent, currentNode.side, currentNode, currentNode)

		if currentNode.left != nil {
			nexts = append(nexts, currentNode.left)
		}
		if currentNode.right != nil {
			nexts = append(nexts, currentNode.right)
		}

		i += 1
		if i > len(links)-1 {
			break
		}
	}
	fmt.Println()

	if len(nexts) > 0 {
		layerPrint(nexts, depth+1)
	}
}

type treeNode struct {
	data   int
	side   string    // 给定节点是其父节点哪个子树
	parent *treeNode // 给定节点的父节点
	left   *treeNode
	right  *treeNode
}

func newTreeNode(x int, side string,
	parentNode *treeNode) *treeNode {
	return &treeNode{
		data:   x,
		side:   side,
		parent: parentNode,
	}
}

func (tn *treeNode) String() string {
	return fmt.Sprintf("V:%d,L:%p,R:%p",
		tn.data, tn.left, tn.right)
}

func (tn *treeNode) insert(x int) {
	var currentNode = tn
	if currentNode == nil {
		return
	}

	for {
		if currentNode.data <= x {
			if currentNode.right != nil {
				currentNode = currentNode.right
			} else {
				currentNode.right = newTreeNode(x, "R", currentNode)
				break
			}
		} else {
			if currentNode.left != nil {
				currentNode = currentNode.left
			} else {
				currentNode.left = newTreeNode(x, "L", currentNode)
				break
			}
		}
	}
}

func (tn *treeNode) find(x int) []*treeNode {
	var currentNode = tn

	var rs []*treeNode

	for {
		if currentNode == nil {
			break
		}

		if currentNode.data > x {
			currentNode = currentNode.left
		} else {
			if currentNode.data == x {
				rs = append(rs, currentNode)
			}
			currentNode = currentNode.right
		}
	}

	return rs
}

func (tn *treeNode) findMinNode() *treeNode {
	var currentNode = tn
	if currentNode == nil {
		return nil
	}

	for {
		if currentNode.left == nil {
			break
		}
		currentNode = currentNode.left
	}

	return currentNode
}

func (tn *treeNode) findMaxNode() *treeNode {
	var currentNode = tn
	if currentNode == nil {
		return nil
	}

	for {
		if currentNode.right == nil {
			break
		}
		currentNode = currentNode.right
	}
	return currentNode
}

func (tn *treeNode) findPrevNode() *treeNode {
	var currentNode = tn
	if currentNode.left != nil {
		return currentNode.left.findMaxNode()
	} else {
		for {
			if currentNode.parent == nil {
				return nil
			}

			if currentNode.side == "R" {
				return currentNode.parent
			}

			currentNode = currentNode.parent
		}
	}
}

func (tn *treeNode) findNextNode() *treeNode {
	var currentNode = tn
	if currentNode.right != nil {
		return currentNode.right.findMinNode()
	} else {
		for {
			if currentNode.parent == nil {
				return nil
			}

			if currentNode.side == "L" {
				return currentNode.parent
			}

			currentNode = currentNode.parent
		}
	}
}
