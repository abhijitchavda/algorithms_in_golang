//Count the number of prime numbers less than a non-negative number, n.
// using Sieve of Eratosthenes
func countPrimes(n int) int {
    count:=0
    if n<=1 {
        return 0
    }
    isp:=make([]bool,n)
    if n>2 {
    isp[2]=true
    }
    for i:=3; i<n ;i++ {
        if i%2==0{
            isp[i]=false
        } else{
            isp[i]=true
        }
    }
    for i:=3;i*i<n;i++ {
        if isp[i] {
            for j:=i*i;j<n;j+=i {
                isp[j]=false
            }
        }
    }
    
    for i:=2;i<n;i++ {
        if isp[i] {
            count++
        }
    }
    return count
}