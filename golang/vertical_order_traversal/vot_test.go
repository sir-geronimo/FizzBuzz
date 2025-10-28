package vot

import (
	"reflect"
	"testing"
)

func TestVerticalOrder(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name: "Full tree example",
			root: &TreeNode{3,
				&TreeNode{9,
					&TreeNode{4, nil, nil},
					&TreeNode{0, nil, nil},
				},
				&TreeNode{8,
					&TreeNode{1, nil, nil},
					&TreeNode{7, nil, nil},
				},
			},
			expected: [][]int{
				{4},
				{9},
				{3, 0, 1},
				{8},
				{7},
			},
		},
		{
			name: "Single node",
			root: &TreeNode{1, nil, nil},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "Left-skewed tree",
			root: &TreeNode{3,
				&TreeNode{2,
					&TreeNode{1, nil, nil},
					nil,
				},
				nil,
			},
			expected: [][]int{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "Right-skewed tree",
			root: &TreeNode{1,
				nil,
				&TreeNode{2,
					nil,
					&TreeNode{3, nil, nil},
				},
			},
			expected: [][]int{
				{1},
				{2},
				{3},
			},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name: "Complex mixed structure",
			root: &TreeNode{10,
				&TreeNode{6,
					&TreeNode{4, nil, nil},
					&TreeNode{8, nil, nil},
				},
				&TreeNode{14,
					&TreeNode{12, nil, nil},
					&TreeNode{16, nil, nil},
				},
			},
			expected: [][]int{
				{4},
				{6},
				{10, 8, 12},
				{14},
				{16},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := verticalOrder(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
