
/*
Given a binary tree, return the inorder traversal of its nodes' values.
*/
  
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    a:=make([]int,0)
    a=append(a[:],getinorder(root,a)...)
    return a
}

func getinorder(root *TreeNode,a []int) []int {
    if root==nil {
        return a[:]
    }
    a=append(a[:],inorderTraversal(root.Left)...)
    a=append(a,root.Val)
    a=append(a[:],inorderTraversal(root.Right)...)
    return a[:]
}