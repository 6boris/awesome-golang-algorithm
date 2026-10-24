package Solution

func Solution(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	days := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	acc := [12]int{}

	for i := 1; i < 12; i++ {
		acc[i] = acc[i-1] + days[i-1]
	}

	parse := func(date string) int {
		_date := int(date[0]-'0')*10 + int(date[1]-'0')
		_day := int(date[3]-'0')*10 + int(date[4]-'0')
		return acc[_date-1] + _day
	}

	as := parse(arriveAlice)
	ae := parse(leaveAlice)
	bs := parse(arriveBob)
	be := parse(leaveBob)
	return max(0, min(ae, be)-max(as, bs)+1)
}
