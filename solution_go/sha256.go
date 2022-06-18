package main

func rightRotate32(x uint32, n int) uint32 {
	right := x >> n
	left := x << (32 - n)
	return left + right
}

func littleSigma0(x uint32) uint32 {
	return rightRotate32(x, 7) ^ rightRotate32(x, 18) ^ (x >> 3)
}

func littleSigma1(x uint32) uint32 {
	return rightRotate32(x, 17) ^ rightRotate32(x, 19) ^ (x >> 10)
}
