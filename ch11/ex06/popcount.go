package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

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

func PopCount3(x uint64) int {
	var ret uint64
	for i := 0; i < 64; i++ {
		b := x >> 1
		ret += x - 2*b
		x = b
		if x == 0 {
			break
		}
	}
	return int(ret)
}

func PopCount4(x uint64) int {
	var ret uint64
	for {
		if x != 0 {
			ret++
		} else {
			break
		}
		x = x & (x - 1)
	}
	return int(ret)
}
