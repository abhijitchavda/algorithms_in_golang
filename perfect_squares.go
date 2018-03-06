/*Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.
For example, given n = 12, return 3 because 12 = 4 + 4 + 4; given n = 13, return 2 because 13 = 4 + 9.
*/

//using Lagrange's Four Square theorem

func numSquares(n int) int {
    if issqr(n) {
        return 1
    }
    sqr:=int(math.Sqrt(float64(n)))
    for i:=1;i<=sqr;i++ {
        if (issqr(n-i*i)){
            return 2
        }
    }
    for ;n%4==0; {
        n=n/4
    }
    if (n-7)%8==0 {
        return 4
    }
    return 3
}

func issqr(n int) bool {
    sqr:=int(math.Sqrt(float64(n)))
    return (sqr*sqr)==n
}