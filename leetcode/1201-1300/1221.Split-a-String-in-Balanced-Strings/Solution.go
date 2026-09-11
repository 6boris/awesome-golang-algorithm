package Solution

func Solution(s string) int {
    var ret int
    r, l := 0, 0
    for i := range s {
        if s[i] == 'R' {
            r++
        } else {
            l++
        }
        if l == r {
            ret++
            l, r = 0, 0
        }
    }
    return ret
}
