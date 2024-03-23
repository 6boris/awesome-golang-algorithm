package Solution

// 00000010100101000001111010011100
// 00111001011110000010100101000000
func Solution(num uint32) uint32 {
	var (
		a, b, aa, bb uint32
	)
	// 1
	a = 0xAAAAAAAA
	b = 0x55555555
	aa = (num & a) >> 1
	bb = (num & b) << 1
	num = aa | bb

	// 1100
	// 0011
	a = 0xCCCCCCCC
	b = 0x33333333
	aa = (num & a) >> 2
	bb = (num & b) << 2
	num = aa | bb

	// 111100001111000011110000
	a = 0xF0F0F0F0
	b = 0x0F0F0F0F
	aa = (num & a) >> 4
	bb = (num & b) << 4
	num = aa | bb

	// 8
	// 1111111100000000
	a = 0xFF00FF00
	b = 0x00FF00FF
	aa = (num & a) >> 8
	bb = (num & b) << 8
	num = aa | bb

	// 16
	a = 0xFFFF0000
	b = 0x0000FFFF
	aa = (num & a) >> 16
	bb = (num & b) << 16
	num = aa | bb

	return num
}
