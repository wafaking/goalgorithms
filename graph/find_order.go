package graph

//课程表II(leetcode-210)
//现在你总共有numCourses门课需要选，记为0到numCourses-1。给你一个数组prerequisites，其中prerequisites[i]=[ai,bi]，表示在选修课程ai前必须先选修bi。
//例如，想要学习课程0，你需要先完成课程1，我们用一个匹配来表示：[0,1]。
//返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回任意一种就可以了。如果不可能完成所有课程，返回一个空数组。
//示例1：输入：numCourses=2,prerequisites=[[1,0]]输出：[0,1]
//解释：总共有2门课程。要学习课程1，你需要先完成课程0。因此，正确的课程顺序为[0,1]。
//示例2：输入：numCourses=4,prerequisites=[[1,0],[2,0],[3,1],[3,2]]输出：[0,2,1,3]
//解释：总共有4门课程。要学习课程3，你应该先完成课程1和课程2。并且课程1和课程2都应该排在课程0之后。
//因此，一个正确的课程顺序是[0,1,2,3]。另一个正确的排序是[0,2,1,3]。
//示例3：输入：numCourses=1,prerequisites=[]输出：[0]

// TODO
func findOrder(numCourses int, prerequisites [][]int) []int {
	var ans = make([]int, 0)
	return ans
}
