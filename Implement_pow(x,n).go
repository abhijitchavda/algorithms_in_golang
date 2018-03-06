/*

Implementation of Pow(x,n)

Decresed number of iterations by 50% 

*/

func myPow(x float64, n int) float64 {
    if x==1 {
        return 1
    }
    if x==0 {
        return 0
    }
    if x==-1 && n%2==0 {
        return 1
    }
    if x==-1 && n%2!=0 {
        return -1
    }
    var ans float64
    ans=1
    flag:=0
    f:=0
    if x<0&&n%2==0{
        x=-1*x
    }
    if x<0&&n%2!=0{
        x=-1*x
        f=1
    }
    if n<0 {
        flag=1
        n*=-1
    }
    for i:=1;i<=n;i++ {
        ans=ans*x
        if n%i==0{
            n=n/i
            x=ans
            i=1
        }
        if ans<=1e-315{
            return 0.0;
        }
    }
    if flag==1{
        if f==1 {
            return -1*(1/ans)
        }
        return 1/ans
    } else {
        if f==1 {
            return -1*ans
        }
    return ans    
    }
    
}