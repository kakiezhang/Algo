/**
二叉搜索树
*/
package main

func main() {
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

func (tn *treeNode) find(x int) []*treeNode {
	var currentNode = tn

	var rs []*treeNode

	for {
		if currentNode == nil {
			break
		}

		if currentNode.data == x {
			rs = append(rs, currentNode)
			break
		} else if currentNode.data > x {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
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
		if currentNode.data >= x {
			if currentNode.right != nil {
				currentNode = currentNode.right
			} else {
				currentNode.right = newTreeNode(x)
			}
		} else {
			if currentNode.left != nil {
				currentNode = currentNode.left
			} else {
				currentNode.left = newTreeNode(x)
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
