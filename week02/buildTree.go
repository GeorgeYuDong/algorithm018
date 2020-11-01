
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * 时间复杂度O(n),空间复杂度O(n)
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{preorder[0], nil, nil}
    stack := []*TreeNode{}
    stack = append(stack, root)
    var inorderIndex int
    for i := 1; i < len(preorder); i++ {
        preorderVal := preorder[i]
        node := stack[len(stack)-1]
        if node.Val != inorder[inorderIndex] {
            node.Left = &TreeNode{preorderVal, nil, nil}
            stack = append(stack, node.Left)
        } else {
            for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
                node = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                inorderIndex++
            }
            node.Right = &TreeNode{preorderVal, nil, nil}
            stack = append(stack, node.Right)
        }
    }
    return root
}

