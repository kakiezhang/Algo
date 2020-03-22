/**
二叉搜索树
*/
package main

import "fmt"

func main() {
	t := &treeNode{
		data: 100,
		left: &treeNode{
			data: 50,
			left: &treeNode{
				data: 30,
			},
			right: &treeNode{
				data: 80,
				left: &treeNode{
					data: 60,
					right: &treeNode{
						data: 70,
					},
				},
				right: &treeNode{
					data: 90,
				},
			},
		},
	}

	testlayerPrint(t)

	testFind(t, 80)

	t.insert(80)
	t.insert(80)
	testlayerPrint(t)

	testFind(t, 80)
}

func testlayerPrint(t *treeNode) {
	layerPrint([]*treeLink{
		&treeLink{
			current: t,
		},
	}, 0)
}

func testFind(t *treeNode, x int) {
	node := t.find(x)
	fmt.Printf("x[%d] found: %v\n", x, node)
}

func layerPrint(links []*treeLink, depth int) {
	var nexts []*treeLink
	var i int

	for {
		if i == 0 {
			fmt.Printf("[Depth:%d] ", depth)
		}

		currentNode := links[i].current
		parentNode := links[i].parent

		fmt.Printf("[%p->%s:%p][%v] ",
			parentNode, links[i].side, currentNode, currentNode)

		if currentNode.left != nil {
			nexts = append(nexts, &treeLink{
				side:    "L",
				parent:  currentNode,
				current: currentNode.left,
			})
		}
		if currentNode.right != nil {
			nexts = append(nexts, &treeLink{
				side:    "R",
				parent:  currentNode,
				current: currentNode.right,
			})
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

type treeLink struct {
	side    string
	parent  *treeNode
	current *treeNode
}

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

func newTreeNode(x int) *treeNode {
	return &treeNode{
		data: x,
	}
}

func (tn *treeNode) String() string {
	return fmt.Sprintf("V:%d,L:%p,R:%p",
		tn.data, tn.left, tn.right)
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
				currentNode.right = newTreeNode(x)
				break
			}
		} else {
			if currentNode.left != nil {
				currentNode = currentNode.left
			} else {
				currentNode.left = newTreeNode(x)
				break
			}
		}
	}
}

func (tn *treeNode) findMinNode() (*treeNode, *treeNode) {
	var currentNode = tn
	if currentNode == nil {
		return nil, nil
	}

	var parentNode *treeNode

	for {
		if currentNode.left == nil {
			break
		}
		parentNode = currentNode
		currentNode = currentNode.left
	}

	return currentNode, parentNode
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

// func (tn *treeNode) findPrevNode() *treeNode {
// }

// func (tn *treeNode) findNextNode() *treeNode {
// }

func (tn *treeNode) del(x int) {
	var currentNode = tn
	var parentNode *treeNode
	var side int // 代表 currentNode 在 parentNode 的哪个子树中，0 是左，1 是右

	for {
		if currentNode == nil {
			break
		}

		if currentNode.data > x {
			parentNode = currentNode
			currentNode = currentNode.right
		} else if currentNode.data < x {
			parentNode = currentNode
			currentNode = currentNode.left
		} else {

			if currentNode.right == nil && currentNode.left == nil {
				// 没有左右子树
				if side == 0 {
					parentNode.left = nil
				} else {
					parentNode.right = nil
				}
			} else if currentNode.right != nil {
				// 右子树存在
				if side == 0 {
					parentNode.left = currentNode.right
				} else {
					parentNode.right = currentNode.right
				}
			} else if currentNode.left != nil {
				// 左子树存在
				if side == 0 {
					parentNode.left = currentNode.left
				} else {
					parentNode.right = currentNode.left
				}
			} else {
				// 左右子树都存在
				currentRight := currentNode.right
				currentLeft := currentNode.left

				// 最小的节点没有左子树
				minNode, minParent := currentNode.right.findMinNode()
				minRight := minNode.right

				minNode.right = currentRight
				minNode.left = currentLeft

				currentNode.right = minRight
				currentNode.left = nil

				if minParent != nil {
					minParent.left = currentNode
				} else {
					minNode.right = currentNode
				}

				if side == 0 {
					parentNode.left = minNode
				} else {
					parentNode.right = minNode
				}
			}
		}
	}
	return
}
