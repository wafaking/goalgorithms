package dynamics

import (
	"sort"
)

//划分为k个相等的子集(leetcode-698)
//给定一个正整数数组nums和一个正整数k，找出是否有可能把这个数组分成k个非空子集，其总和都相等。
//示例1：输入：nums=[4,3,2,3,5,2,1],k=4输出：True说明：有可能将其分成4个子集(5),(1,4),(2,3),(2,3)等于总和。
//示例2:输入:nums=[1,2,3,4],k=3输出:false
//注：0<nums[i]<10000,每个元素的频率在[1,4]范围内

// 深度遍历
func canPartitionKSubsets1(nums []int, k int) bool {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if k == 0 || sum%k != 0 {
		return false
	}

	sum /= k
	sort.Ints(nums)
	// 最后一个元素大于均值，则无法等分；
	if nums[len(nums)-1] > sum {
		return false
	}

	// 创建k个桶，每个桶的总和都为sum(sum为均分后的值)
	var (
		buckets = make([]int, k)
		dfs     func(idx int) bool
	)
	for i := range buckets {
		buckets[i] = sum
	}

	dfs = func(idx int) bool {
		if idx < 0 {
			// 遍历-1位置，说明元素已经全部放到桶里了，因为每个桶内的总和是一样的，如果有一个元素无法放入，则不会遍历到-1位置
			return true
		}

		// 遍历每个桶，看idx位置的数值适合放哪个桶里
		for i := 0; i < k; i++ {
			// 即如果一个桶的值与当前桶值相等，那当前元素肯定也不能放入当前桶中(减少递归次数，可以省去)
			if i > 0 && buckets[i] == buckets[i-1] {
				continue
			}

			// 如果nums[idx]刚好满足要求 || 可以放入num[idx]且放入num[idx]后还能放入最小元素
			if buckets[i] == nums[idx] || (idx > 0 && buckets[i]-nums[idx] >= nums[0]) {
				// 当前元素放入桶i中
				buckets[i] -= nums[idx]
				if dfs(idx - 1) { // 继续处理下一个元素
					return true
				}
				// 回溯:当前元素放入桶i中不合适，取出
				buckets[i] += nums[idx]
			}

		}
		return false
	}

	return dfs(len(nums) - 1)
}

// TODO 状态压缩+记忆化搜索
func canPartitionKSubsets2(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if k == 0 || sum%k > 0 {
		return false
	}
	per := sum / k
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] > per {
		return false
	}

	dp := make([]bool, 1<<n)
	for i := range dp {
		dp[i] = true
	}

	var dfs func(int, int) bool
	dfs = func(s, p int) bool {
		if s == 0 {
			return true
		}
		if !dp[s] {
			return dp[s]
		}
		dp[s] = false
		for i, num := range nums {
			if num+p > per {
				break
			}
			if s>>i&1 > 0 && dfs(s^1<<i, (p+nums[i])%per) {
				return true
			}
		}
		return false
	}
	return dfs(1<<n-1, 0)
}

// TODO 状态压缩 + 动态规划
func canPartitionKSubsets3(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k > 0 {
		return false
	}
	per := sum / k
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] > per {
		return false
	}

	dp := make([]bool, 1<<n)
	dp[0] = true
	curSum := make([]int, 1<<n)
	for i, v := range dp {
		if !v {
			continue
		}
		for j, num := range nums {
			if curSum[i]+num > per {
				break
			}
			if i>>j&1 == 0 {
				next := i | 1<<j
				if !dp[next] {
					curSum[next] = (curSum[i] + nums[j]) % per
					dp[next] = true

				}
			}
		}
	}
	return dp[1<<n-1]
}

/*
	class Solution {
    boolean[] visited;
    int avg;
    int n;
    int[] values;
    public boolean canPartitionKSubsets(int[] nums, int k) {
        n = nums.length;
        // 如果总和不能平均分为k份，返回false
        int total = Arrays.stream(nums).sum();
        if (total % k != 0) return false;

        avg = total / k;
        Arrays.sort(nums);

        // 最大的数大于avg，无解
        if (nums[n - 1] > avg) return false;

        values = nums;

        // visited数组，存放每个状态(state)是否访问过
        // state=100101，表示nums中下标为0、2、5的数字被使用过了
        visited = new boolean[1 << n];

        // 一开始，所有数字都没有被使用过
        // (1 << n) - 1 相当于0~n-1位都是1，比如n=2，就是 011
        return dfs((1 << n) - 1, 0);
    }

    boolean dfs(int state, int total) {
        // state=0说明每一个数都被使用了
        if (state == 0) {
            return true;
        }
        // 剩余数字的状态为state第二次被访问到，则这个状态无解，有解的话已经返回true了，不会再访问到这个状态
        if (visited[state]) return false;

        // state被访问过
        visited[state] = true;
        // 枚举剩余的数字
        for (int i = 0;i < n;++i) {
            // state二进制表示，第i位为0，说明该位置的数字被使用了
            if ((state & (1 << i)) == 0) continue;

            // 引入这个数字如果大于avg，因为从小到大排序，所以后面的数字也就不用看了，无解
            if (values[i] + total > avg) break;

            // state & (~(1 << i))   通过位运算，将state的第i位设置为0，
            // values[i] + total     代表将values[i]放入当前集合
            // 通过 %avg,能够使得values[i] + total = avg时，变为0，相当于引入一个新的空集合
            if (dfs(state & (~(1 << i)), (values[i] + total) % avg)) {
                return true;
            }
        }
        return false;
    }
}
*/
