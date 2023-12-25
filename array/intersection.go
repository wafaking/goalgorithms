package array

//两个数组的交集(leetcode-349)
//给定两个数组nums1和nums2，返回它们的交集。输出结果中的每个元素一定是唯一的。我们可以不考虑输出结果的顺序。
//示例1：输入：nums1=[1,2,2,1],nums2=[2,2]输出：[2]
//示例2：输入：nums1=[4,9,5],nums2=[9,4,9,8,4]输出：[9,4]解释：[4,9]也是可通过的

func intersection1(nums1 []int, nums2 []int) []int {
	var (
		m, existMap = make(map[int]struct{}, 0), make(map[int]struct{}, 0)
		ans         = make([]int, 0)
	)
	for i := range nums1 {
		if _, ok := m[nums1[i]]; !ok {
			m[nums1[i]] = struct{}{}
		}
	}
	for i := range nums2 {
		if _, ok := m[nums2[i]]; ok {
			if _, ok := existMap[nums2[i]]; !ok {
				ans = append(ans, nums2[i])
				existMap[nums2[i]] = struct{}{}
			}
		}
	}
	return ans
}
