//Given an array of strings, group anagrams together.
func groupAnagrams(strs []string) [][]string {
    a:=make(map[string]int)
    k:=1
    ans:=make([][]string,0)
    for _,v:=range strs {
        if a[SortString(v)]==0 {
            a[SortString(v)]=k
            x:=make([]string,0)
            ans=append(ans,x)
            k++
        }
        ans[a[SortString(v)]-1]=append(ans[a[SortString(v)]-1],v)
    }
    return ans
}
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}