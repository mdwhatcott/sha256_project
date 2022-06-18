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

func messageSchedule(input string) (result []uint32) {
	result = make([]uint32, 64)
	for i := 0; i < 16; i++ {
		n := i * 4
		b := input[n : n+4]
		result[i] = uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	}
	for i := 16; i < 64; i++ {
		result[i] = result[i-16] + littleSigma0(result[i-15]) + result[i-7] + littleSigma1(result[i-2])
	}
	return result
}

func bigSigma0(x uint32) uint32 {
	return rightRotate32(x, 2) ^ rightRotate32(x, 13) ^ rightRotate32(x, 22)
}

func bigSigma1(x uint32) uint32 {
	return rightRotate32(x, 6) ^ rightRotate32(x, 11) ^ rightRotate32(x, 25)
}

func bitwiseComplement(x uint32) uint32 {
	return uint32(-int(x) - 1) // https://wiki.python.org/moin/BitwiseOperators (see the ~ operator)
}

func choice(x, y, z uint32) uint32 {
	return (x & y) ^ (bitwiseComplement(x) & z)
}

func majority(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func round(state []uint32, roundConstant, scheduleWord uint32) (result []uint32) {
	ch := choice(state[4], state[5], state[6])
	temp1 := state[7] + bigSigma1(state[4]) + ch + roundConstant + scheduleWord
	temp2 := bigSigma0(state[0]) + majority(state[0], state[1], state[2])
	return append(result,
		temp1+temp2,
		state[0],
		state[1],
		state[2],
		state[3]+temp1,
		state[4],
		state[5],
		state[6],
	)
}

var RoundConstants = []uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

func compress(input []uint32, block string) (result []uint32) {
	schedule := messageSchedule(block)
	state := input[:]
	for i := 0; i < 64; i++ {
		state = round(state, RoundConstants[i], schedule[i])
	}
	return append(result,
		input[0]+state[0],
		input[1]+state[1],
		input[2]+state[2],
		input[3]+state[3],
		input[4]+state[4],
		input[5]+state[5],
		input[6]+state[6],
		input[7]+state[7],
	)
}
