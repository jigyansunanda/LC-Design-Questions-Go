package binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
	head  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bstIterator := BSTIterator{}

	// Push all left side nodes to stack: O(h)
	for root != nil {
		bstIterator.stack = append(bstIterator.stack, root)
		root = root.Left
	}

	// Set head of bst iterator
	if len(bstIterator.stack) > 0 {
		smallestNodeIndex := len(bstIterator.stack) - 1
		bstIterator.head = bstIterator.stack[smallestNodeIndex]
	}

	return bstIterator
}

/*
This following implementations has the time complexity as follows:
Next() : O(h)
HasNext(): O(1)
*/

func (bstIterator *BSTIterator) Next() int {
	currentHead := bstIterator.head

	// First pop the current head from stack
	lastNodeIndex := len(bstIterator.stack) - 1
	bstIterator.stack = bstIterator.stack[:lastNodeIndex]

	// Check whether right node of current head exists or not
	// If exists: then traverse across it's left branch and push left nodes
	// If does not exist: do nothing
	if currentHead.Right != nil {
		rightTopNode := currentHead.Right
		for rightTopNode != nil {
			bstIterator.stack = append(bstIterator.stack, rightTopNode)
			rightTopNode = rightTopNode.Left
		}
	}

	if len(bstIterator.stack) > 0 {
		lastNodeIndex := len(bstIterator.stack) - 1
		bstIterator.head = bstIterator.stack[lastNodeIndex]
	} else {
		bstIterator.head = nil
	}

	return currentHead.Val
}

func (bstIterator *BSTIterator) HasNext() bool {
	return bstIterator.head != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
