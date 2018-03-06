/*
Given a non-empty array of integers, return the k most frequent elements.

For example,
Given [1,1,1,2,2,3] and k = 2, return [1,2].

*/


type Pair struct {
  Key int
  Value int
}

type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortMapByValue(m map[int]int) PairList {
   p := make(PairList, len(m))
   i := 0
   for k, v := range m {
      p[i] = Pair{k, v}
       i++
   }
   sort.Sort(p)
   return p
}

func topKFrequent(nums []int, k int) []int {
    a:=make(map[int]int)
    b:=make([]int,k)
    for _,v:=range nums {
        a[v]++
    }
    var p PairList
    p=sortMapByValue(a)
    x:=len(p)
    j:=0
    for i:=x-1;k>0;k-- {
        b[j]=p[i].Key
        i--
        j++
    }
    return b
}