
//Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
 }
 
func kthSmallest(root *TreeNode, k int) int {
    ans,_:=getmin(root,k,0)
    return ans
}
func getmin(node *TreeNode,count int,ans int) (int,int) {
    if ans != 0 {
        return ans,count
    }
    if node==nil {
        return ans,count
    }
    ans,count=getmin(node.Left,count,ans)
    count--
    if count==0{
        ans=node.Val
        return ans,count
    }
    ans,count=getmin(node.Right,count,ans)
    return ans,count
}