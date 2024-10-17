package sprint

func TransposeMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}

	transposed := make([][]int, len(matrix[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(matrix))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}
