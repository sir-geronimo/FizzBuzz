package vot

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeWithCoords struct {
	row, col, value int
}

// verticalOrder returns node values from leftmost to rightmost column
func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// Step 1: DFS traversal to collect all nodes with coordinates
	nodes := make([]nodeWithCoords, 0)
	var traverse func(node *TreeNode, row, col int)

	// DFS function to record nodes with their (row, col) coordinates
	traverse = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}

		// Record the node with its coordinates
		nodes = append(nodes, nodeWithCoords{row, col, node.Val})

		// Traverse left and right children with updated coordinates
		traverse(node.Left, row+1, col-1)
		traverse(node.Right, row+1, col+1)
	}

	traverse(root, 0, 0)

	// Step 2: Sort nodes by (col asc, row asc, value asc)
	slices.SortFunc(nodes, func(a, b nodeWithCoords) int {
		// Sort by column first
		if a.col != b.col {
			return a.col - b.col // left-most column first
		}

		// Same column, sort by row
		if a.row != b.row {
			return a.row - b.row // top rows first
		}

		// Same column and row, sort by value
		return a.value - b.value // smaller values first
	})

	// Step 3: Group by column and collect stack
	output := make([][]int, 0) // Final output structure
	prevCol := nodes[0].col    // Initialise with the first node's column
	var values []int           // To accumulate values for the current column

	// Iterate through sorted nodes to group by column
	for _, n := range nodes {
		// New column detected
		if n.col != prevCol {
			// Append collected values to output
			output = append(output, values)
			// Reset values accumulator for new column
			values = []int{}
			// Update previous column
			prevCol = n.col
		}

		// Append current node's value to accumulated values
		values = append(values, n.value)
	}

	// Append the last column's values to output
	if len(values) > 0 {
		output = append(output, values)
	}

	return output
}
