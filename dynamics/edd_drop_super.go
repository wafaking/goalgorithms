package dynamics

//鸡蛋掉落(leetcode-887)
//给你k枚相同的鸡蛋和一栋从第1层到第n层共有n层楼的建筑。
//已知存在楼层f(0<=f<=n),任何从高于f的楼层落下的鸡蛋都会碎，从f楼层或比它低的楼层落下的鸡蛋都不会破。
//每次你可以取一枚没有碎的鸡蛋把它从楼层x扔下（满足1<=x<=n）:
//	1.如果鸡蛋碎了，就不能再次使用它;
//	2.如果鸡蛋没有摔碎，则可以重复使用;
//请你计算并返回要确定f确切的值的最小操作次数是多少？

//示例1：输入：k=1,n=2,输出：2; 解释：
//	鸡蛋从1楼掉落。如果它碎了，肯定能得出f=0。
//	否则，鸡蛋从2楼掉落。如果它碎了，肯定能得出f=1。
//	如果它没碎，那么肯定能得出f=2。
//	因此，在最坏的情况下我们需要移动2次以确定f是多少。
//示例2：输入：k=2,n=6,输出：3
//示例3：输入：k=3,n=14,输出：4

func superEggDrop(k int, n int) int {
	return 0
}
