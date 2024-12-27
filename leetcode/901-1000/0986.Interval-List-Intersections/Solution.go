package Solution

func Solution(firstList [][]int, secondList [][]int) [][]int {
	la, lb := len(firstList), len(secondList)
	if la == 0 || lb == 0 {
		return nil
	}
	a, b := 0, 0
	ans := make([][]int, 0)
	for a < la && b < lb {
		af, as := firstList[a][0], firstList[a][1]
		bf, bs := secondList[b][0], secondList[b][1]
		if af > bs {
			b++
			continue
		}
		if bf > as {
			a++
			continue
		}
		ans = append(ans, []int{max(af, bf), min(as, bs)})
		if bs > as {
			a++
		} else {
			b++
		}
	}
	return ans
}
