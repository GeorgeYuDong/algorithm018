func nthUglyNumber(n int) int {
    a,b,c:= 0,0,0
    dp:= make([]int,n,n)
    for i:=0;i<n;i++{
        dp[i]=1
    }
    min := func(a,b int)int{
        if a>b{
            return b
        }
        return a
    }
    for i:=1;i<n;i++{
        t1,t2,t3:=dp[a]*2,dp[b]*3,dp[c]*5
        dp[i]=min(min(t1,t2),t3)
        if dp[i] ==t1 {
            a++
        }
        if dp[i] ==t2 {
            b++
        }
        if dp[i] ==t3 {
            c++
        }
    }
    return dp[n-1]
}

