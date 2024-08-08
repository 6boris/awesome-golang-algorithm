package Solution

func Solution(rows int, cols int, rStart int, cStart int) [][]int {
	ans := make([][]int, 0)
	steps := 1
	first := true
	end := rows * cols
	for end > 0 {
		if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
			ans = append(ans, []int{rStart, cStart})
			end--
		}
		//right
		if !first {
			for s := steps - 1; s > 0 && end > 0; s-- {
				rStart--
				if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
					ans = append(ans, []int{rStart, cStart})
					end--
				}
			}
		}
		for s := steps; s > 0 && end > 0; s-- {
			cStart++
			if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
				ans = append(ans, []int{rStart, cStart})
				end--
			}
		}
		for s := steps; s > 0 && end > 0; s-- {
			rStart++
			if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
				ans = append(ans, []int{rStart, cStart})
				end--
			}
		}

		for s := steps; s > 0 && end > 0; s-- {
			cStart--
			if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
				ans = append(ans, []int{rStart, cStart})
				end--
			}
		}
		cStart--
		steps += 2
		first = false
	}
	return ans
}
