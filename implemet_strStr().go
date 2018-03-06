//Implement strStr().
//Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack

func strStr(haystack string, needle string) int {
    i,j,index := 0,0,-1
    flag :=0
    if len(needle)==0 {
        return 0
    }
    if len(needle)>len(haystack) {
        return index
    }    
    for i=0 ; i<len(haystack) ; i++ {
        if flag==len(needle) {
           return index 
        }
        if needle[j]==haystack[i] {
            if(index==-1){
                index=i
                fmt.Println(index)
            }
            flag=flag+1
            j++
            continue
        } else {
            if index!=-1 {
            i=index}
            index=-1
            flag=0;
            j=0;
            continue
        }
    }
    if flag==len(needle) {
        return index
    } else {
        return -1
    }
}