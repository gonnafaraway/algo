package math

func GeneratePascalTriangle(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, len(result)+1)
		row[0], row[len(row)-1] = 1, 1
		if i > 0 {
			midRows := len(row) - 2
			for j := 1; j <= midRows; j++ {
				row[j] = result[j][j-1] + result[j][j]
				if i > 2 && j+1 < len(result) {
					row[j] = result[len(result)-1][j-1] + result[len(result)-1][j]
				} else {
					row[j] = result[j][j-1] + result[j][j]
				}
			}
		}
		result = append(result, row)
	}
	return result
}
