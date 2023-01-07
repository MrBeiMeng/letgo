package code_lists

/*图片平滑器 | https://leetcode.cn/problems/image-smoother*/

func imageSmoother(img [][]int) [][]int {

	result := make([][]int, 0)
	for i := range img {
		tmpArr := make([]int, 0)
		for j := 0; j < len(img[i]); j++ {
			tmpArr = append(tmpArr, img[i][j])
		}
		result = append(result, tmpArr)
	}

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			result[i][j] = imageSmootherGetCell(i, j, img)
		}
	}

	return result
}

func imageSmootherGetCell(y, x int, img [][]int) int {

	var total int
	var sum int

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			tmpI := y + i
			tmpJ := x + j

			if tmpI < 0 || tmpJ < 0 || tmpI >= len(img) || tmpJ >= len(img[y]) {
				continue
			}

			total++
			sum += img[tmpI][tmpJ]
		}
	}

	return sum / total
}
