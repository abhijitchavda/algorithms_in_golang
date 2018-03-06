/*
Rotate an array of n elements to the right by k steps.

For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
*/
func rotate(nums []int, k int)  {
    n:=len(nums)
    if k>n {
        k=k-n
    }
    a:=make([]int,n-k)
    b:=make([]int,k)
    j:=0
    for i:=0 ; i<n ; i++ {
        if(i<n-k) {
            a[i]=nums[i]
        } else {
            b[j]=nums[i]
            j++
        }
    }
    for i,v:=range b {
        nums[i]=v
    }
    j=0
    for i:=len(b) ; i<n ; i++ {
        nums[i]=a[j]
        j++
    }
}