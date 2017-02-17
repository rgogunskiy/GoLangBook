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
// используя цикл
func PopCount2(x uint64) int {
	var bitCount int = 0

	var i uint
	for ; i <= 8; i++ {
		bitCount += int(pc[byte(x>>(i*8))])
	}

	return bitCount
}
