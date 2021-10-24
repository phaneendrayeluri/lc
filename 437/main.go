package main

func main() {}

func PathSum(nodes []int, targetSum int) int {
	return pathSum(toTree(nodes, 0), targetSum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func toTree(input []int, index int) *TreeNode {
	if index > len(input)-1 {
		return nil
	}
	if input[index] == 0 {
		return nil
	}
	return &TreeNode{
		Val:   input[index],
		Left:  toTree(input, (2*index)+1),
		Right: toTree(input, (2*index)+2),
	}
}

func pathSumAtNode(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	targetSum -= root.Val
	if targetSum == 0 {
		return 1 + pathSumAtNode(root.Left, targetSum) + pathSumAtNode(root.Right, targetSum)
	}
	return pathSumAtNode(root.Left, targetSum) + pathSumAtNode(root.Right, targetSum)
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSumAtNode(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}
