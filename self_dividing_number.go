/*A self-dividing number is a number that is divisible by every digit it contains.

Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.
*/

func selfDividingNumbers(left int, right int) []int {
    ans:=make([]int,right-left)
    f:=0
    for i:=left ; i<=right ;i++ {
        if i>0 && i<=9 {
            ans[f]=i
            f++
            continue
        }
        n:=i
        for ;n>0; {
            x:=n%10
            if x==0 {
                break
            }
            if i%x==0 {
                n=int(n/10)
                continue
            } else {
                break
            }
        }
        if n==0 {
            ans[f]=i
            f++
        } 
    }
    return ans[:f]
}