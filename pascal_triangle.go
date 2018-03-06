func generate(numRows int) [][]int {
    a:=make([][]int,numRows)
    x:=0
    for i:=1 ; i<=numRows ;i++ {
        a[x]=make([]int,i)
        x++
    }
    for i:=0 ; numRows>0 ;i++ {
        for j:=0 ; j<=i ; j++ {
            if j==0 {
                a[i][j]=1
                continue
            }
            if j == i {
                a[i][j]=1
                break
            }
            a[i][j]=a[i-1][j]+a[i-1][j-1]
        }
        numRows--
    }
    return a
}