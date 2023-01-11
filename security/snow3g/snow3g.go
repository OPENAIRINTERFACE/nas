// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package snow3g

var sr = [...]byte{
	0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76,
	0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0,
	0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15,
	0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75,
	0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84,
	0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf,
	0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8,
	0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2,
	0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73,
	0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb,
	0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79,
	0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08,
	0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a,
	0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e,
	0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf,
	0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16,
}

var sq = [...]byte{
	0x25, 0x24, 0x73, 0x67, 0xd7, 0xae, 0x5c, 0x30, 0xa4, 0xee, 0x6e, 0xcb, 0x7d, 0xb5, 0x82, 0xdb,
	0xe4, 0x8e, 0x48, 0x49, 0x4f, 0x5d, 0x6a, 0x78, 0x70, 0x88, 0xe8, 0x5f, 0x5e, 0x84, 0x65, 0xe2,
	0xd8, 0xe9, 0xcc, 0xed, 0x40, 0x2f, 0x11, 0x28, 0x57, 0xd2, 0xac, 0xe3, 0x4a, 0x15, 0x1b, 0xb9,
	0xb2, 0x80, 0x85, 0xa6, 0x2e, 0x02, 0x47, 0x29, 0x07, 0x4b, 0x0e, 0xc1, 0x51, 0xaa, 0x89, 0xd4,
	0xca, 0x01, 0x46, 0xb3, 0xef, 0xdd, 0x44, 0x7b, 0xc2, 0x7f, 0xbe, 0xc3, 0x9f, 0x20, 0x4c, 0x64,
	0x83, 0xa2, 0x68, 0x42, 0x13, 0xb4, 0x41, 0xcd, 0xba, 0xc6, 0xbb, 0x6d, 0x4d, 0x71, 0x21, 0xf4,
	0x8d, 0xb0, 0xe5, 0x93, 0xfe, 0x8f, 0xe6, 0xcf, 0x43, 0x45, 0x31, 0x22, 0x37, 0x36, 0x96, 0xfa,
	0xbc, 0x0f, 0x08, 0x52, 0x1d, 0x55, 0x1a, 0xc5, 0x4e, 0x23, 0x69, 0x7a, 0x92, 0xff, 0x5b, 0x5a,
	0xeb, 0x9a, 0x1c, 0xa9, 0xd1, 0x7e, 0x0d, 0xfc, 0x50, 0x8a, 0xb6, 0x62, 0xf5, 0x0a, 0xf8, 0xdc,
	0x03, 0x3c, 0x0c, 0x39, 0xf1, 0xb8, 0xf3, 0x3d, 0xf2, 0xd5, 0x97, 0x66, 0x81, 0x32, 0xa0, 0x00,
	0x06, 0xce, 0xf6, 0xea, 0xb7, 0x17, 0xf7, 0x8c, 0x79, 0xd6, 0xa7, 0xbf, 0x8b, 0x3f, 0x1f, 0x53,
	0x63, 0x75, 0x35, 0x2c, 0x60, 0xfd, 0x27, 0xd3, 0x94, 0xa5, 0x7c, 0xa1, 0x05, 0x58, 0x2d, 0xbd,
	0xd9, 0xc7, 0xaf, 0x6b, 0x54, 0x0b, 0xe0, 0x38, 0x04, 0xc8, 0x9d, 0xe7, 0x14, 0xb1, 0x87, 0x9c,
	0xdf, 0x6f, 0xf9, 0xda, 0x2a, 0xc4, 0x59, 0x16, 0x74, 0x91, 0xab, 0x26, 0x61, 0x76, 0x34, 0x2b,
	0xad, 0x99, 0xfb, 0x72, 0xec, 0x33, 0x12, 0xde, 0x98, 0x3b, 0xc0, 0x9b, 0x3e, 0x18, 0x10, 0x3a,
	0x56, 0xe1, 0x77, 0xc9, 0x1e, 0x9e, 0x95, 0xa3, 0x90, 0x19, 0xa8, 0x6c, 0x09, 0xd0, 0xf0, 0x86,
}

type Lfsr struct {
	s [16]uint32
}

type Fsm struct {
	r [3]uint32
}

var lfsr Lfsr
var fsm Fsm

func mulx(V, c byte) byte {
	if V&0x80 != 0 {
		return (V << 1) ^ c
	} else {
		return V << 1
	}
}

func mulxPow(V, i, c byte) byte {
	if i == 0 {
		return V
	} else {
		return mulx(mulxPow(V, i-1, c), c)
	}
}

func s1(w uint32) uint32 {
	w0 := (w >> 24) & 0xff
	w1 := (w >> 16) & 0xff
	w2 := (w >> 8) & 0xff
	w3 := w & 0xff

	r0 := uint32(
		mulx(sr[w0], 0x1b) ^ sr[w1] ^ sr[w2] ^ mulx(sr[w3], 0x1b) ^ sr[w3],
	)
	r1 := uint32(
		mulx(sr[w0], 0x1b) ^ sr[w0] ^ mulx(sr[w1], 0x1b) ^ sr[w2] ^ sr[w3],
	)
	r2 := uint32(
		sr[w0] ^ mulx(sr[w1], 0x1b) ^ sr[w1] ^ mulx(sr[w2], 0x1b) ^ sr[w3],
	)
	r3 := uint32(
		sr[w0] ^ sr[w1] ^ mulx(sr[w2], 0x1b) ^ sr[w2] ^ mulx(sr[w3], 0x1b),
	)
	return (r0 << 24) | (r1 << 16) | (r2 << 8) | r3
}

func s2(w uint32) uint32 {
	w0 := (w >> 24) & 0xff
	w1 := (w >> 16) & 0xff
	w2 := (w >> 8) & 0xff
	w3 := w & 0xff

	r0 := uint32(
		mulx(sq[w0], 0x69) ^ sq[w1] ^ sq[w2] ^ mulx(sq[w3], 0x69) ^ sq[w3],
	)
	r1 := uint32(
		mulx(sq[w0], 0x69) ^ sq[w0] ^ mulx(sq[w1], 0x69) ^ sq[w2] ^ sq[w3],
	)
	r2 := uint32(
		sq[w0] ^ mulx(sq[w1], 0x69) ^ sq[w1] ^ mulx(sq[w2], 0x69) ^ sq[w3],
	)
	r3 := uint32(
		sq[w0] ^ sq[w1] ^ mulx(sq[w2], 0x69) ^ sq[w2] ^ mulx(sq[w3], 0x69),
	)
	return (r0 << 24) | (r1 << 16) | (r2 << 8) | r3
}

func mulAlpha(c byte) uint32 {
	r0 := uint32(mulxPow(c, 23, 0xa9))
	r1 := uint32(mulxPow(c, 245, 0xa9))
	r2 := uint32(mulxPow(c, 48, 0xa9))
	r3 := uint32(mulxPow(c, 239, 0xa9))
	return (r0 << 24) | (r1 << 16) | (r2 << 8) | r3
}

func divAlpha(c byte) uint32 {
	r0 := uint32(mulxPow(c, 16, 0xa9))
	r1 := uint32(mulxPow(c, 39, 0xa9))
	r2 := uint32(mulxPow(c, 6, 0xa9))
	r3 := uint32(mulxPow(c, 64, 0xa9))
	return (r0 << 24) | (r1 << 16) | (r2 << 8) | r3
}

func lfsrInitialisationMode(F uint32) {
	v := (lfsr.s[0] << 8) ^ mulAlpha(
		byte(lfsr.s[0]>>24)&0xff,
	) ^ lfsr.s[2] ^ (lfsr.s[11] >> 8) ^
		divAlpha(
			byte(lfsr.s[11]&0xff),
		) ^ F
	for i := 0; i < 15; i++ {
		lfsr.s[i] = lfsr.s[i+1]
	}
	lfsr.s[15] = v
}

func lfsrKeystreamMode() {
	v := (lfsr.s[0] << 8) ^ mulAlpha(
		byte(lfsr.s[0]>>24)&0xff,
	) ^ lfsr.s[2] ^ (lfsr.s[11] >> 8) ^
		divAlpha(
			byte(lfsr.s[11]&0xff),
		)
	for i := 0; i < 15; i++ {
		lfsr.s[i] = lfsr.s[i+1]
	}
	lfsr.s[15] = v
}

func clockFsm(s15, s5 uint32) uint32 {
	F := (s15 + fsm.r[0]) ^ fsm.r[1]
	r := fsm.r[1] + (fsm.r[2] ^ s5)
	fsm.r[2] = s2(fsm.r[1])
	fsm.r[1] = s1(fsm.r[0])
	fsm.r[0] = r
	return F
}

func InitSnow3g(k, iv [4]uint32) {
	lfsr.s[0] = k[0] ^ 0xffffffff
	lfsr.s[1] = k[1] ^ 0xffffffff
	lfsr.s[2] = k[2] ^ 0xffffffff
	lfsr.s[3] = k[3] ^ 0xffffffff
	lfsr.s[4] = k[0]
	lfsr.s[5] = k[1]
	lfsr.s[6] = k[2]
	lfsr.s[7] = k[3]
	lfsr.s[8] = k[0] ^ 0xffffffff
	lfsr.s[9] = k[1] ^ 0xffffffff ^ iv[3]
	lfsr.s[10] = k[2] ^ 0xffffffff ^ iv[2]
	lfsr.s[11] = k[3] ^ 0xffffffff
	lfsr.s[12] = k[0] ^ iv[1]
	lfsr.s[13] = k[1]
	lfsr.s[14] = k[2]
	lfsr.s[15] = k[3] ^ iv[0]

	for i := 0; i < 3; i++ {
		fsm.r[i] = 0
	}

	for i := 0; i < 32; i++ {
		F := clockFsm(lfsr.s[15], lfsr.s[5])
		lfsrInitialisationMode(F)
	}
}

func GenerateKeystream(n int, ks []uint32) {
	clockFsm(lfsr.s[15], lfsr.s[5])
	lfsrKeystreamMode()

	for i := 0; i < n; i++ {
		F := clockFsm(lfsr.s[15], lfsr.s[5])
		ks[i] = F ^ lfsr.s[0]
		lfsrKeystreamMode()
	}
}
