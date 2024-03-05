package dynamics

import "goalgorithms/common"

//最长重复子数组(leetcode-718)
//给两个整数数组nums1和nums2，返回两个数组中公共的、长度最长的子数组的长度。
//示例1：输入：nums1=[1,2,3,2,1],nums2=[3,2,1,4,7]输出：3,解释：长度最长的公共子数组是[3,2,1]。
//示例2：输入：nums1=[0,0,0,0,0],nums2=[0,0,0,0,0]输出：5

func findLength(nums1 []int, nums2 []int) int {
	// dp[i][j]表示text1的前i个字符与text2的前j个字符的最长公共子串的长度
	// 此处用text1表示纵坐标，text2表示横坐标
	var (
		dp        = make([][]int, len(nums1)+1)
		maxLength int // 记录最大长度(最大长度子串可能在前面)
	)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	// 第一行为空字符与nums2的公共子串，全为0,第一列同理
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if nums1[i-1] == nums2[j-1] { // 匹配
				dp[i][j] = dp[i-1][j-1] + 1 // 取上一个匹配数量加1
				maxLength = common.MaxInTwo(maxLength, dp[i][j])
			}
		}
	}

	return maxLength
}
