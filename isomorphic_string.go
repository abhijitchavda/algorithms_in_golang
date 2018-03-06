/*
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

For example,
Given "egg", "add", return true.

Given "foo", "bar", return false.
*/

func isIsomorphic(s string, t string) bool {
    a:=make(map[string]int)
    for i,v:=range s {
        if a[string(v)]==0{
            a[string(v)]=i+1
        } else {
            a[string(v)]*=10
            a[string(v)]+=i+1
        }
    }
    x:=make([]int,len(a))
    j:=0
    for i,v:=range a{
        x[j]=v
        j++
        delete(a,i)
    }
    for i,v:=range t {
        if a[string(v)]==0{
            a[string(v)]=i+1
        } else {
            a[string(v)]*=10
            a[string(v)]+=i+1
        }
    }
    y:=make([]int,len(a))
    k:=0
    for _,v:=range a {
        y[k]=v
        k++
    }
    if j!=k {
        return false
    }
    sort.Ints(x)
    sort.Ints(y)
    
    for i:=0;i<j;i++ {
        if x[i]!=y[i] {
            return false
        }
    }
    return true
    
}