/**
 * url: https://leetcode-cn.com/problems/super-egg-drop/
 * id: 887
 *
 * 你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
 * 每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
 * 你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
 * 每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
 * 你的目标是确切地知道 F 的值是多少。
 * 无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？
 *
 * 示例 1：
 * 输入：K = 1, N = 2
 * 输出：2
 *
 * 解释：
 * 鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
 * 否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
 * 如果它没碎，那么我们肯定知道 F = 2 。
 * 因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
 *
 * 示例 2：
 * 输入：K = 2, N = 6
 * 输出：3
 *
 * 示例 3：
 * 输入：K = 3, N = 14
 * 输出：4
 *
 * 提示：
 * 1 <= K <= 100
 * 1 <= N <= 10000
 */

package main

import (
	"fmt"
)

func main() {
	k := 1
	n := 2
	fmt.Println(superEggDrop(k, n))
}

func superEggDrop(K int, N int) int {
	//方法1
	dp := make([][]int, K+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, N+1)
	}
	for m := 1; m <= N; m++ {
		dp[0][m] = 0 // zero egg
		for k := 1; k <= K; k++ {
			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
			if dp[k][m] >= N {
				return m
			}
		}
	}
	return N

	//方法2
	moves := 0
	dp := make([]int, K+1)
	for dp[K] < N {
		moves++
		for i := K; i > 0; i-- {
			dp[i] += dp[i-1] + 1
			// dp[moves][k] = 1 + dp[moves-1][k-1] + dp[moves-1][k]
			// dp[0][i] = 0
		}

	}
	return moves
}

/**
 * 思路：
 * d[k][m] 表示 k 个鸡蛋， 在 m 次移动中可以确定的楼层数。
 * 递归关系如下：
 * 假定在 x 层扔了， 如果破了， 可以确定 x 层下面的 d[k-1][m-1] 层；
 * 如果没破， 可以确定楼上的 d[k][m-1] 层 另外 x 层被确认 。
 * 所以最终确认的层数为 1 + d[k-1][m-1] + d[k][m-1]
 *
 * 执行用时 : 12 ms, 在Super Egg Drop的Go提交中击败了33.33% 的用户
 * 内存消耗 : 11 MB, 在Super Egg Drop的Go提交中击败了10.26% 的用户
 */
