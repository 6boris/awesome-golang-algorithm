package Solution

func Solution(coordinates string) bool {
	x, y := coordinates[0]-'a', coordinates[1]-'0'

	black := (x&1 ==0 && y&1 != 0) || (x&1 != 0 && y &1 == 0)
	return !black
}
