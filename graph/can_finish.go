package graph

//课程表(leetcode-207)
//你这个学期必须选修numCourses门课程，记为0到numCourses-1。
//在选修某些课程之前需要一些先修课程。先修课程按数组prerequisites给出，其中prerequisites[i]=[ai,bi]，表示如果要学习课程ai则必须先学习课程bi。
//例如，先修课程对[0,1]表示：想要学习课程0，你需要先完成课程1。
//请你判断是否可能完成所有课程的学习？如果可以，返回true；否则，返回false。
//示例1：输入：numCourses=2,prerequisites=[[1,0]]输出：true
//解释：总共有2门课程。学习课程1之前，你需要完成课程0。这是可能的。
//示例2：输入：numCourses=2,prerequisites=[[1,0],[0,1]]输出：false
//解释：总共有2门课程。学习课程1之前，你需要先完成课程0；并且学习课程0之前，你还应先完成课程1。这是不可能的。
//注：
//	prerequisites[i].length==2
//	prerequisites[i]中的所有课程对互不相同

// 广度优先
func canFinish1(numCourses int, prerequisites [][]int) bool {
	//[A,B]即学习A必须要先学习B，拓扑图为B-->A则B的出度为1，A的入度为1,出度为0
	var (
		// edges[i]表示i的出度课程列表
		// 即入度为A的列表[B]
		edges = make([][]int, numCourses)
		// 存储节点的入度数
		inDegree = make([]int, numCourses)
		queue    = make([]int, 0)
		ans      int // 用于记录已经学习的课程
	)
	for _, item := range prerequisites {
		edges[item[1]] = append(edges[item[1]], item[0])
		inDegree[item[0]]++
	}

	// 找出入度为0的节点先处理
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		// 因添加进queue中的都是入度为0的节点，因此直接计入已学完课程即可
		ans++
		// 将入度为cur的节点减去1
		for _, v := range edges[cur] {
			inDegree[v]--
			//即如果节点的入度为0，表明可以学习此课程了
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return ans >= numCourses
}

// 深度优先
func canFinish2(numCourses int, prerequisites [][]int) bool {
	var (
		// edges[i]表示i的出度列表
		edges = make([][]int, numCourses)
		// visited[i]用于记录i的搜索状态,0:未搜索;1:搜索中;2:已完成
		visited = make([]int, numCourses)
		// 表示是否是无环图，默认是无环
		valid = true
		dfs   func(i int)
	)
	dfs = func(i int) {
		// 将该节点标记为搜索中
		visited[i] = 1
		// 处理i的出度节点列表
		for _, v := range edges[i] {
			if visited[v] == 0 {
				dfs(v)
				if !valid { // 如果v是有环图要结束
					return
				}
			} else if visited[v] == 1 { // 说明这个也在搜索中，有环
				valid = false
				return
			}
		}
		// 当前节点搜索完成
		visited[i] = 2
	}

	// 填充i的出度列表
	for _, item := range prerequisites {
		edges[item[1]] = append(edges[item[1]], item[0])
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 { // 处理未搜索过的节点
			dfs(i)
		}
	}

	return valid
}
