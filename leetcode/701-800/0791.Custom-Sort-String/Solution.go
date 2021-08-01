package Solution

import "sort"

func Solution(order string, str string) string {
    mapper := make(map[byte]int)
    for i:= range order {
        mapper[order[i]] = i - 26
    }
    res := []byte(str)
    sort.Slice(res, func(i, j int) bool {
        return mapper[res[i]] < mapper[res[j]]
    })
    return string(res)
}
