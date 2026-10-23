package Solution

import "sort"

type Seg struct {
	left, right int
}

func Solution(s string) []string {
	seg := make([]Seg, 26)

	for i := 0; i < 26; i++ {
		seg[i] = Seg{-1, -1}
	}

	// Preprocess the left and right endpoints.
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		if seg[idx].left == -1 {
			seg[idx].left = i
			seg[idx].right = i
		} else {
			seg[idx].right = i
		}
	}

	for i := 0; i < 26; i++ {
		if seg[i].left != -1 {
			j := seg[i].left
			for j <= seg[i].right {
				idx := int(s[j] - 'a')

				if !(seg[i].left <= seg[idx].left &&
					seg[idx].right <= seg[i].right) {

					if seg[idx].left < seg[i].left {
						seg[i].left = seg[idx].left
					}

					if seg[idx].right > seg[i].right {
						seg[i].right = seg[idx].right
					}

					j = seg[i].left
				}

				j++
			}
		}
	}

	// Greedily select intervals.
	sort.Slice(seg, func(i, j int) bool {
		if seg[i].right == seg[j].right {
			return seg[i].left > seg[j].left
		}
		return seg[i].right < seg[j].right
	})

	ans := []string{}
	end := -1

	for _, segment := range seg {
		left, right := segment.left, segment.right

		if left == -1 {
			continue
		}

		if end == -1 || left > end {
			end = right
			ans = append(ans, s[left:right+1])
		}
	}

	return ans
}
