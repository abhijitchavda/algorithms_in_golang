
/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/

 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

type queue struct {
    arr [] *TreeNode
}
func (s *queue) pop() *TreeNode {
    x:=s.arr[0]
    y:=len(s.arr)
    s.arr=s.arr[1:y]
    return x
}
func (s *queue) push (x *TreeNode) {
    s.arr=append(s.arr,x)
}
func (s *queue) isempty() bool {
    if len(s.arr)==0 {
        return true
    } else {
        return false
    }
}
func (s *queue) length() int {
    return len(s.arr)
}

func levelOrder(root *TreeNode) [][]int {
    a:=[][]int{}
    if root==nil {
        return a
    }
    var s queue
    s.push(root)
    x:=1
    length:=s.length()
    z:=make([]int,0)
    for ;!s.isempty(); {
        if length==0 {
            a=append(a,z[len(z)-x:])
            length=s.length()
            x=length
        }
        v:=s.pop()
        z=append(z,v.Val)
        length--
        if v.Left!=nil {
        s.push(v.Left)    
        }
        if v.Right!=nil{
        s.push(v.Right)
        }
    }
    a=append(a,z[len(z)-x:])
    return a
}
