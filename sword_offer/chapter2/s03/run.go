package s03

func run(a [][]int, k int) bool {
	if len(a) == 0 {
		return false
	}
	rows := len(a)
	cols := len(a[0])

	row, col := 0, cols-1

	for row < rows && col > 0 {
		if a[row][col] == k {
			return true
		} else if a[row][col] < k {
			row++
		} else {
			col--
		}
	}
	return false
}
