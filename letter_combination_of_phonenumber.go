/*
Given a digit string, return all possible letter combinations that the number could represent.

A mapping of digit to letters --> just like on the telephone buttons
*/
func letterCombinations(digits string) []string {
    length:=1
    if len(digits)==0 {
        ans:=make([]string,0)
        return ans
    }
    for _,v:=range digits {
        if string(v)=="7" || string(v)=="9" {
            length*=4
        } else {
            length*=3
        }
    }
    ans := make([]string,length)
    leng:=0
    for _,v := range digits {
        leng,ans=appender(string(v),ans,leng)
    }
    return ans
}

func appender(cur string, ans []string,leng int) (int,[]string){
    y:=make(map[string]string)
    y["2"]="abc"
    y["3"]="def"
    y["4"]="ghi"
    y["5"]="jkl"
    y["6"]="mno"
    y["7"]="pqrs"
    y["8"]="tuv"
    y["9"]="wxyz"
    x:=len(y[cur])
    temp:=make([]string,0)
    count:=0
    if leng==0 {
        for i:=0;i<x;i++ {
            ans[i]=string(y[cur][i])
            count++
        }
        return count,ans
    }
    for i:=0;i<len(y[cur]);i++ {
        for j:=0;j<leng;j++ {
            temp=append(temp,string(ans[j])+string(y[cur][i]))
            count++
        }
    }
    return count,temp
}