package levenshtein

// Distance - This function returns the Levenshtein-Distance
// between the two argument strings
func Distance(s1 string, s2 string) int {
	matrix := initMatrix(s1, s2)

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {

			topValue := matrix[i-1][j]
			leftValue := matrix[i][j-1]
			topLeftValue := matrix[i-1][j-1]

			minValue := min(topValue, leftValue, topLeftValue)

			if s1[i-1] == s2[j-1] {
				matrix[i][j] = minValue
			} else {
				matrix[i][j] = minValue + 1
			}
		}
	}

	return matrix[len(s1)][len(s2)]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}

func initMatrix(s1 string, s2 string) [][]int {
	s1Len := len(s1) + 1
	s2Len := len(s2) + 1
	matrix := make([][]int, s1Len)

	for i := 0; i < s1Len; i++ {

		matrix[i] = make([]int, s2Len)

		for j := 0; j < s2Len; j++ {
			if i == 0 {
				matrix[i][j] = j
			} else if j == 0 {
				matrix[i][j] = i
			}
		}
	}

	return matrix
}
