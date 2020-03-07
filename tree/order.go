/**
二叉树的遍历
按前序，中序，后序
按层遍历
*/
package main

import "fmt"

func main() {
	t := &tree{
		data: 1,
		left: &tree{
			data: 2,
			left: &tree{
				data: 4,
			},
			right: &tree{
				data: 5,
			},
		},
		right: &tree{
			data: 3,
			left: &tree{
				data: 6,
			},
			right: &tree{
				data: 7,
			},
		},
	}

	fmt.Print("preOrder: ")
	preOrder(t)
	fmt.Println()

	fmt.Print("inOrder: ")
	inOrder(t)
	fmt.Println()

	fmt.Print("postOrder: ")
	postOrder(t)
	fmt.Println()

	fmt.Print("layerOrder: ")
	layerOrder(t)
	fmt.Println()
}

type tree struct {
	data  int
	left  *tree
	right *tree
}

func preOrder(t *tree) {
	if t == nil {
		return
	}

	fmt.Printf("%d ", t.data)
	preOrder(t.left)
	preOrder(t.right)
}

func inOrder(t *tree) {
	if t == nil {
		return
	}

	inOrder(t.left)
	fmt.Printf("%d ", t.data)
	inOrder(t.right)
}

func postOrder(t *tree) {
	if t == nil {
		return
	}

	postOrder(t.left)
	postOrder(t.right)
	fmt.Printf("%d ", t.data)
}

func layerOrder(t *tree) {
	printLayerNodes([]*tree{t})
}

func printLayerNodes(ts []*tree) {
	if len(ts) == 0 {
		return
	}

	var nextLayerTrees []*tree

	for _, t := range ts {
		if t != nil {
			fmt.Printf("%d ", t.data)
			nextLayerTrees = append(nextLayerTrees, t.left)
			nextLayerTrees = append(nextLayerTrees, t.right)
		}
	}

	printLayerNodes(nextLayerTrees)
}
