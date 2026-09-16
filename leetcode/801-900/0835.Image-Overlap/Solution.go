package Solution

func Solution(img1 [][]int, img2 [][]int) int {

	ans := 0
	for y := 0; y < len(img1); y++ {
		for x := 0; x < len(img1); x++ {
			ans1 := shift(x, y, img1, img2)
			if ans1 > ans {
				ans = ans1
			}

			ans1 = shift(x, y, img2, img1)
			if ans1 > ans {
				ans = ans1
			}

		}
	}
	return ans
}

func shift(x, y int, img1, img2 [][]int) int {
	lShift, rShift := 0, 0
	length := len(img1)
	img2Row := 0

	for img1Row := y; img1Row < length; img1Row++ {
		img2Col := 0
		for img1Col := x; img1Col < length; img1Col++ {
			// 向左移动x各单位，会导致第x列移动到第0列所以
			if img1[img1Row][img1Col] == 1 && img2[img2Row][img2Col] == 1 {
				lShift++
			}

			// 向右移动x个单位，那么第0列会被移动，img2Col是从0开始的
			// img1[img1Row][img2Col]那个
			// img1 的某一行向右移动一个单位后变成 0 1 1。那么要比较这一行只需要比较 1 1 部分即可。在不改变原数组的情况下
			// 要访问到1 1 我们需要从下标0开始。
			if img1[img1Row][img2Col] == 1 && img2[img2Row][img1Col] == 1 {
				rShift++
			}
			img2Col++
		}
		img2Row++
	}

	if lShift > rShift {
		return lShift
	}
	return rShift
}
