package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount возвращает степень заполнения
// (количество установленных битов) значения x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 возвращает степень заполнения
// количество установленных битов) значения x,
// используя x&(x-1)
func PopCount2(x uint64) int {
	bitCount := 0

	for x != 0 {
		bitCount++
		x = x & (x - 1)
	}
	return bitCount
}
