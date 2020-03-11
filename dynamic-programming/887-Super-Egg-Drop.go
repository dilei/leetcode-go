package dynamic_programming

/*
需要将问题转化一下，变成已知鸡蛋个数，和操作次数，求最多能测多少层楼的临界点.
还是使用动态规划 Dynamic Programming 来做，用一个二维 DP 数组，
其中 dp[i][j] 表示当有i次操作，且有j个鸡蛋时能测出的最高的楼层数。
再来考虑状态转移方程如何写，由于 dp[i][j] 表示的是在第i次移动且使用第j个鸡蛋测试第 dp[i-1][j-1]+1 层，因为上一个状态是第i-1次移动，且用第j-1个鸡蛋。
此时还是有两种情况：

鸡蛋碎掉：说明至少可以测到的不会碎的层数就是 dp[i-1][j-1]。
鸡蛋没碎：那这个鸡蛋可以继续利用，此时我们还可以再向上查找 dp[i-1][j] 层。
那么加上当前层，总共可以通过i次操作和j个鸡蛋查找的层数范围是 [0, dp[i-1][j-1] + dp[i-1][j] + 1]，这样就可以得到状态转移方程如下：
*/
func superEggDrop(K int, N int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
	}

	m := 0
	for dp[m][K] < N {
		m++
		for j := 1; j <= K; j++ {
			dp[m][j] = dp[m-1][j-1] + dp[m-1][j] + 1
		}
	}
	return m
}
