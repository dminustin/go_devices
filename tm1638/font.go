// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tm1638

import "strconv"

/*
The bits are displayed by mapping bellow
 -- 0 --
|       |
5       1
 -- 6 --
4       2
|       |
 -- 3 --  .7
*/

// definition for standard hexadecimal numbers
var fontNumber = []byte{
	b("00111111"), // 0
	b("00000110"), // 1
	b("01011011"), // 2
	b("01001111"), // 3
	b("01100110"), // 4
	b("01101101"), // 5
	b("01111101"), // 6
	b("00000111"), // 7
	b("01111111"), // 8
	b("01101111"), // 9
	b("01110111"), // A
	b("01111100"), // B
	b("00111001"), // C
	b("01011110"), // D
	b("01111001"), // E
	b("01110001"), // F
}

// definition for error
var fontErrorData = []byte{
	b("01111001"), // E
	b("01010000"), // r
	b("01010000"), // r
	b("01011100"), // o
	b("01010000"), // r
	0,
	0,
	0,
}

// definition for the displayable ASCII chars
var fontDefault = []byte{
	b("00000000"), // (32)  <space>
	b("10000110"), // (33)	!
	b("00100010"), // (34)	"
	b("01111110"), // (35)	#
	b("01101101"), // (36)	$
	b("00000000"), // (37)	%
	b("00000000"), // (38)	&
	b("00000010"), // (39)	'
	b("00110000"), // (40)	(
	b("00000110"), // (41)	)
	b("01100011"), // (42)	*
	b("00000000"), // (43)	+
	b("00000100"), // (44)	,
	b("01000000"), // (45)	-
	b("10000000"), // (46)	.
	b("01010010"), // (47)	/
	b("00111111"), // (48)	0
	b("00000110"), // (49)	1
	b("01011011"), // (50)	2
	b("01001111"), // (51)	3
	b("01100110"), // (52)	4
	b("01101101"), // (53)	5
	b("01111101"), // (54)	6
	b("00100111"), // (55)	7
	b("01111111"), // (56)	8
	b("01101111"), // (57)	9
	b("00000000"), // (58)	:
	b("00000000"), // (59)	;
	b("00000000"), // (60)	<
	b("01001000"), // (61)	=
	b("00000000"), // (62)	>
	b("01010011"), // (63)	?
	b("01011111"), // (64)	@
	b("01110111"), // (65)	A
	b("01111111"), // (66)	B
	b("00111001"), // (67)	C
	b("00111111"), // (68)	D
	b("01111001"), // (69)	E
	b("01110001"), // (70)	F
	b("00111101"), // (71)	G
	b("01110110"), // (72)	H
	b("00000110"), // (73)	I
	b("00011111"), // (74)	J
	b("01101001"), // (75)	K
	b("00111000"), // (76)	L
	b("00010101"), // (77)	M
	b("00110111"), // (78)	N
	b("00111111"), // (79)	O
	b("01110011"), // (80)	P
	b("01100111"), // (81)	Q
	b("00110001"), // (82)	R
	b("01101101"), // (83)	S
	b("01111000"), // (84)	T
	b("00111110"), // (85)	U
	b("00101010"), // (86)	V
	b("00011101"), // (87)	W
	b("01110110"), // (88)	X
	b("01101110"), // (89)	Y
	b("01011011"), // (90)	Z
	b("00111001"), // (91)	[
	b("01100100"), // (92)	\ (this can't be the last char on a line, even in comment or it'll concat)
	b("00001111"), // (93)	]
	b("00000000"), // (94)	^
	b("00001000"), // (95)	_
	b("00100000"), // (96)	`
	b("01011111"), // (97)	a
	b("01111100"), // (98)	b
	b("01011000"), // (99)	c
	b("01011110"), // (100)	d
	b("01111011"), // (101)	e
	b("00110001"), // (102)	f
	b("01101111"), // (103)	g
	b("01110100"), // (104)	h
	b("00000100"), // (105)	i
	b("00001110"), // (106)	j
	b("01110101"), // (107)	k
	b("00110000"), // (108)	l
	b("01010101"), // (109)	m
	b("01010100"), // (110)	n
	b("01011100"), // (111)	o
	b("01110011"), // (112)	p
	b("01100111"), // (113)	q
	b("01010000"), // (114)	r
	b("01101101"), // (115)	s
	b("01111000"), // (116)	t
	b("00011100"), // (117)	u
	b("00101010"), // (118)	v
	b("00011101"), // (119)	w
	b("01110110"), // (120)	x
	b("01101110"), // (121)	y
	b("01000111"), // (122)	z
	b("01000110"), // (123)	{
	b("00000110"), // (124)	|
	b("01110000"), // (125)	}
	b("00000001"), // (126)	~
}

func b(s string) byte {
	u, err := strconv.ParseUint(s, 2, 8)
	if err != nil {
		panic(err)
	}
	return byte(u)
}
