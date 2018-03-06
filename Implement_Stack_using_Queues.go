/*

Implement the following operations of a stack using queues.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
empty() -- Return whether the stack is empty.
*/
type MyStack struct {
    a []int
    top int
}


/** Initialized data structure here. */
func Constructor() MyStack {
    var this MyStack
    return this
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.top=x
    this.a=append(this.a,x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    b:=make([]int,0)
    for ; 0<len(this.a)-1 ; {
        b=append(b,this.a[0])
        this.a=this.a[1:]
    }
    x:=this.a[0]
    this.a=this.a[:0]
    if len(b)!=0 {
    for ;len(b)>1; {
        this.Push(b[0])
        b=b[1:]
    }
    
    this.Push(b[0])
    b=b[:0]
    }
    return x
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.top
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    if len(this.a)==0 {
        return true
    }
    return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */