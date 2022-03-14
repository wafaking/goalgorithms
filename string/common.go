package str

func reverse(str string) string {
	bt := []byte(str)
	for start, end := 0, len(str)-1; start < end; start, end = start+1, end-1 {
		bt[start], bt[end] = bt[end], bt[start]
	}
	return string(bt)
}
