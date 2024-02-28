package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	path []*TreeNode
}

func getAllLeftNodes(n *TreeNode) []*TreeNode {
	var path []*TreeNode
	for ; n != nil; n = n.Left {
		path = append(path, n)
	}
	return path
}

func pathValues(p []*TreeNode) []int {
	var r []int
	for _, n := range p {
		r = append(r, n.Val)
	}
	return r
}

func computeNextPath(p []*TreeNode) []*TreeNode {
	var child *TreeNode
	var back bool
	for {
		n := len(p)
		if n == 0 {
			return nil
		}
		last := p[n-1]
		if last.Right == child {
			p = p[:n-1]
			child = last
			back = true
			continue
		}
		if back {
			return p
		}
		return append(p, getAllLeftNodes(last.Right)...)
	}
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		path: getAllLeftNodes(root),
	}
}

func (i *BSTIterator) Next() int {
	last := i.path[len(i.path)-1]
	newP := computeNextPath(i.path)
	i.path = newP
	return last.Val
}

func (i *BSTIterator) HasNext() bool {
	return len(i.path) > 0
}
