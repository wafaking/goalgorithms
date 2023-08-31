package dynamics

//最后一块石头的重量II(leetcode-1049)
//有一堆石头，用整数数组stones表示。其中stones[i]表示第i块石头的重量。
//每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为x和y，且x<=y。那么粉碎的可能结果如下：
//如果x==y，那么两块石头都会被完全粉碎；
//如果x!=y，那么重量为x的石头将会完全粉碎，而重量为y的石头新重量为y-x。
//最后，最多只会剩下一块石头。返回此石头最小的可能重量。如果没有石头剩下，就返回0。

//示例1：输入：stones=[2,7,4,1,8,1],输出：1
//解释：
//	组合2和4，得到2，所以数组转化为[2,7,1,8,1]，
//	组合7和8，得到1，所以数组转化为[2,1,1,1]，
//	组合2和1，得到1，所以数组转化为[1,1,1]，
//	组合1和1，得到0，所以数组转化为[1]，这就是最优值。
//示例2： 输入：stones=[31,26,33,21,40],输出：5

//TODO ---
func lastStoneWeightII(stones []int) int {
	//1,1,2,4,7,8
	//var sum int
	//for _, v := range stones {
	//	sum += v
	//}
	//sort.Ints(stones)
	//var (
	//	mid         = sum / 2
	//	left, right = 0, len(stones) - 1
	//	temp        int
	//	diff        int
	//)
	//
	//for left < right {
	//	temp = stones[left] + stones[right]
	//	diff = abs(mid - temp)
	//	if diff==0{
	//		return 0
	//	}else if
	//}
	return 0
}
