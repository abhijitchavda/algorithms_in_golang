/* Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.
*/

func longestPalindrome(s string) int {
    var a [58]int
    var count int
    flag:=0
    for _,v:=range s {
        a[int(v)-65]++
    }
    for _,v:=range a {
        if v==1{
            flag=1
            continue
        }
        if v%2==0 {
            count+=((v/2)*2)
            continue
        } else {
            v=v-1
            flag=1
            count+=((v/2)*2)
        }
    }
    if flag==1 {
        return count+1
    } else {
        return count
    }
}