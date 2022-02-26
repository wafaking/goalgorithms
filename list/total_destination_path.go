package list

//对于一个m*n的网格，从左上角的方格到右下角的方格，共有多少条路径（只允许向右和向下）

// 方法类似青蛙跳台阶问题，即f(m,n)=f(m-1,n)+f(m,n-1)
func TotalDestinationPath(m, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	return TotalDestinationPath(m-1, n) + TotalDestinationPath(m, n-1)
}
